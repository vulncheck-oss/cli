package errs

import (
	"context"
	"errors"
	"net"
	"net/url"

	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

// Classify walks an arbitrary error chain and returns the *Error that best
// describes it. The function is idempotent: a *Error in flows straight back
// out. nil in returns nil out.
//
// New SDK / runtime errors that need a distinct Kind should be added here so
// the entire CLI gains the correct exit code + JSON code at once.
func Classify(err error) *Error {
	if err == nil {
		return nil
	}

	// Already classified — preserve the original.
	var classified *Error
	if errors.As(err, &classified) {
		return classified
	}

	// Cancellation (SIGINT / parent context). Highest priority because
	// downstream HTTP errors often wrap it.
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return Wrap(KindCancelled, err, "cancelled")
	}

	// SDK sentinel for 401.
	if errors.Is(err, sdk.ErrorUnauthorized) {
		return AuthInvalid(err, "unauthorized: token is missing or invalid")
	}

	// SDK structured HTTP error.
	var reqErr sdk.ReqError
	if errors.As(err, &reqErr) {
		return fromReqError(reqErr)
	}

	// Command-body validation (pkg/ui's FlagError type).
	var flagErr *ui.FlagError
	if errors.As(err, &flagErr) {
		return Wrap(KindValidation, err, "%s", err.Error())
	}

	// Network-layer failures.
	var netOp *net.OpError
	if errors.As(err, &netOp) {
		return Wrap(KindNetwork, err, "network error: %s", err.Error())
	}
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		// url.Error wraps everything from dial errors to context cancels;
		// the inner error decides the Kind. Recurse to pick that up.
		if inner := Classify(urlErr.Err); inner != nil {
			// Preserve the URL context in the message.
			inner.cause = err
			return inner
		}
		return Wrap(KindNetwork, err, "request failed: %s", err.Error())
	}
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		return Wrap(KindNetwork, err, "dns lookup failed: %s", dnsErr.Error())
	}

	// Fall through: surface as an internal error so the exit code is 1.
	return Wrap(KindInternal, err, "%s", err.Error())
}

func fromReqError(r sdk.ReqError) *Error {
	msg := defaultMessageFor(r)
	switch r.StatusCode {
	case 401:
		e := Wrap(KindAuthInvalid, r, "%s", msg)
		e.HTTPStatus = 401
		return e
	case 402:
		e := Wrap(KindAuthInvalid, r, "%s", msg)
		e.HTTPStatus = 402
		return e
	case 403:
		e := Wrap(KindAuthInvalid, r, "%s", msg)
		e.HTTPStatus = 403
		return e
	case 404:
		e := Wrap(KindNotFound, r, "%s", msg)
		e.HTTPStatus = 404
		return e
	case 429:
		e := Wrap(KindRateLimited, r, "%s", msg)
		e.HTTPStatus = 429
		return e
	default:
		kind := KindInternal
		if r.StatusCode >= 400 && r.StatusCode < 500 {
			kind = KindBadRequest
		}
		e := Wrap(kind, r, "%s", msg)
		e.HTTPStatus = r.StatusCode
		return e
	}
}

func defaultMessageFor(r sdk.ReqError) string {
	if len(r.Reason.Errors) > 0 {
		return joinErrors(r.Reason.Errors)
	}
	switch r.StatusCode {
	case 401:
		return "unauthorized: token is missing or invalid"
	case 402:
		return "payment required: this endpoint is not included in your subscription"
	case 403:
		return "forbidden: token lacks permission for this resource"
	case 404:
		return "not found"
	case 429:
		return "rate limit exceeded; please retry later"
	default:
		return r.Error()
	}
}

func joinErrors(errs []string) string {
	switch len(errs) {
	case 0:
		return ""
	case 1:
		return errs[0]
	default:
		joined := errs[0]
		for _, m := range errs[1:] {
			joined += "; " + m
		}
		return joined
	}
}
