package errs

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func TestKindExitCodes(t *testing.T) {
	cases := map[Kind]int{
		KindInternal:     1,
		KindValidation:   2,
		KindBadRequest:   2,
		KindAuthRequired: 3,
		KindAuthInvalid:  3,
		KindNotFound:     4,
		KindRateLimited:  5,
		KindNetwork:      6,
		KindCancelled:    130,
	}
	for k, want := range cases {
		if got := k.ExitCode(); got != want {
			t.Errorf("%s: exit code = %d, want %d", k, got, want)
		}
	}
}

func TestClassifyNilIsNil(t *testing.T) {
	if got := Classify(nil); got != nil {
		t.Fatalf("Classify(nil) = %v, want nil", got)
	}
}

func TestClassifyIsIdempotent(t *testing.T) {
	in := AuthInvalid(nil, "nope")
	out := Classify(in)
	if out != in {
		t.Fatal("Classify on an already-classified error should return the same pointer")
	}
}

func TestClassifyDetectsCancellation(t *testing.T) {
	e := Classify(context.Canceled)
	if e == nil || e.Kind != KindCancelled {
		t.Fatalf("expected KindCancelled, got %v", e)
	}
	if e.ExitCode() != 130 {
		t.Fatalf("expected exit 130, got %d", e.ExitCode())
	}
}

func TestClassifyDetectsSDKUnauthorized(t *testing.T) {
	e := Classify(sdk.ErrorUnauthorized)
	if e == nil || e.Kind != KindAuthInvalid {
		t.Fatalf("expected KindAuthInvalid, got %v", e)
	}
	if e.HTTPStatus != 401 {
		t.Fatalf("expected HTTPStatus 401, got %d", e.HTTPStatus)
	}
}

func TestClassifyMapsReqErrorByStatus(t *testing.T) {
	cases := []struct {
		status int
		want   Kind
	}{
		{401, KindAuthInvalid},
		{402, KindAuthInvalid},
		{403, KindAuthInvalid},
		{404, KindNotFound},
		{429, KindRateLimited},
		{400, KindBadRequest},
		{500, KindInternal},
	}
	for _, c := range cases {
		err := sdk.ReqError{StatusCode: c.status, Reason: sdk.MetaError{Error: true, Errors: []string{"boom"}}}
		got := Classify(err)
		if got == nil {
			t.Fatalf("status=%d: Classify returned nil", c.status)
		}
		if got.Kind != c.want {
			t.Errorf("status=%d: kind = %s, want %s", c.status, got.Kind, c.want)
		}
		if got.HTTPStatus != c.status {
			t.Errorf("status=%d: HTTPStatus = %d, want %d", c.status, got.HTTPStatus, c.status)
		}
	}
}

func TestClassifyDetectsFlagError(t *testing.T) {
	err := ui.Error("missing arg")
	got := Classify(err)
	if got == nil || got.Kind != KindValidation {
		t.Fatalf("expected KindValidation, got %v", got)
	}
}

func TestErrorUnwrapPreservesChain(t *testing.T) {
	base := errors.New("root cause")
	wrapped := Wrap(KindInternal, base, "outer")
	if !errors.Is(wrapped, base) {
		t.Fatal("errors.Is should see through Wrap")
	}
}

func TestValidationHelperFormats(t *testing.T) {
	e := Validation("need %s", "input")
	if e.Kind != KindValidation {
		t.Errorf("kind = %s, want %s", e.Kind, KindValidation)
	}
	if e.Error() != "need input" {
		t.Errorf("message = %q, want %q", e.Error(), "need input")
	}
}

// The v4 endpoints answer an unentitled token with 402 and a route-specific
// message. That message is more precise than anything we could substitute, so
// it must reach the user verbatim rather than being replaced by a generic hint.
func TestClassifyPreservesEntitlementMessages(t *testing.T) {
	cases := []struct {
		name string
		msg  string
	}{
		{"advisory", "This endpoint requires a valid trial or paid subscription"},
		{"backup", "V4 backups require the Exploit & Vulnerability Intelligence subscription"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := sdk.ReqError{StatusCode: 402, Reason: sdk.MetaError{Error: true, Errors: []string{c.msg}}}
			got := Classify(err)
			if got == nil {
				t.Fatal("Classify returned nil")
			}
			if got.Kind != KindAuthInvalid {
				t.Errorf("kind = %s, want %s", got.Kind, KindAuthInvalid)
			}
			if got.HTTPStatus != 402 {
				t.Errorf("HTTPStatus = %d, want 402", got.HTTPStatus)
			}
			if !strings.Contains(got.Error(), c.msg) {
				t.Errorf("message %q does not carry the server's wording %q", got.Error(), c.msg)
			}
			// Exit 3 is the point of the 402 case: without it a 402 falls to
			// the default 4xx branch and reports exit 2, telling a script to
			// fix its query when the problem is its entitlement.
			if code := got.Kind.ExitCode(); code != 3 {
				t.Errorf("exit code = %d, want 3 (auth, not bad request)", code)
			}
		})
	}
}

// Every status with a case in fromReqError needs a fallback message for a body
// that does not decode, or the user sees "errors: []".
func TestClassifyHasAFallbackMessageForEveryMappedStatus(t *testing.T) {
	for _, status := range []int{401, 402, 403, 404, 429} {
		got := Classify(sdk.ReqError{StatusCode: status})
		if got == nil {
			t.Fatalf("status=%d: Classify returned nil", status)
		}
		if got.Error() == "" || strings.Contains(got.Error(), "[]") {
			t.Errorf("status=%d: message = %q, want a human fallback", status, got.Error())
		}
	}
}

// A v4 string-shaped error body must survive decoding all the way to the
// surfaced message; before MetaError.UnmarshalJSON it arrived as an empty slice.
func TestClassifySurfacesV4StringShapedErrors(t *testing.T) {
	var metaError sdk.MetaError
	body := `{"error":"feed not found: \"nosuchfeed\""}`
	if err := json.Unmarshal([]byte(body), &metaError); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	got := Classify(sdk.ReqError{StatusCode: 404, Reason: metaError})
	if got == nil {
		t.Fatal("Classify returned nil")
	}
	if got.Kind != KindNotFound {
		t.Errorf("kind = %s, want %s", got.Kind, KindNotFound)
	}
	if !strings.Contains(got.Error(), "nosuchfeed") {
		t.Errorf("message %q lost the feed name", got.Error())
	}
}
