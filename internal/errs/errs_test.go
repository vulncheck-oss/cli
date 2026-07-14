package errs

import (
	"context"
	"errors"
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
