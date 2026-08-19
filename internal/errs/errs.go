// Package errs defines the CLI's structured error taxonomy.
//
// Agents and scripts consume the CLI's exit code + stable JSON error code
// to decide what to do next. Both come from this package:
//
//   - Kind is a stable string identifier (e.g. "auth_required") that
//     survives across releases. Add new Kinds; never re-purpose old ones.
//   - ExitCode maps each Kind to a small integer in the range [1, 7],
//     stable across releases. Exit code 130 is reserved for SIGINT (POSIX).
//
// Error wraps the underlying cause but carries the classified Kind + the
// agent-facing message. Classify(err) walks an arbitrary error chain and
// produces a *Error; commands that already know the right kind can build
// one directly via Validation/AuthRequired/etc.
package errs

import "fmt"

// Kind is the agent-facing error identifier. The string value is the
// stable API contract — once shipped, do not rename.
type Kind string

const (
	// KindInternal is the catch-all for unexpected errors. Exit 1.
	KindInternal Kind = "internal"
	// KindValidation covers bad CLI arguments, missing required flags,
	// or any pre-flight input problem the user can fix. Exit 2.
	KindValidation Kind = "validation"
	// KindAuthRequired means no token was supplied at all. Exit 3.
	KindAuthRequired Kind = "auth_required"
	// KindAuthInvalid means a token was supplied but the server rejected
	// it (401, 403, or session.CheckAuth failure). Exit 3.
	KindAuthInvalid Kind = "auth_invalid"
	// KindNotFound is HTTP 404 / "no such index" / "token id unknown". Exit 4.
	KindNotFound Kind = "not_found"
	// KindRateLimited is HTTP 429. Exit 5.
	KindRateLimited Kind = "rate_limited"
	// KindNetwork covers DNS failures, connection refused, request timeouts. Exit 6.
	KindNetwork Kind = "network"
	// KindBadRequest is HTTP 4xx that is not 401/403/404/429 — typically a
	// malformed query the user can fix. Exit 2 (alongside validation).
	KindBadRequest Kind = "bad_request"
	// KindCancelled means the user (or a parent process) sent SIGINT/SIGTERM
	// or cancelled the context. Exit 130 to match POSIX (128 + SIGINT).
	KindCancelled Kind = "cancelled"
)

// ExitCode is the stable process exit code mapped to a Kind.
func (k Kind) ExitCode() int {
	switch k {
	case KindValidation, KindBadRequest:
		return 2
	case KindAuthRequired, KindAuthInvalid:
		return 3
	case KindNotFound:
		return 4
	case KindRateLimited:
		return 5
	case KindNetwork:
		return 6
	case KindCancelled:
		return 130
	default:
		return 1
	}
}

// Error is the CLI's classified error type. It implements `error` and can
// wrap an underlying cause via Unwrap so callers can still use errors.Is
// against SDK sentinels.
type Error struct {
	Kind       Kind
	Message    string
	HTTPStatus int
	// Hint is optional remediation context that explains *why* the caller hit
	// this error, when the message alone is not actionable. Rendered after the
	// message on stderr and as `error.hint` in the JSON envelope. Additive:
	// callers that don't know the field simply ignore it.
	Hint  string
	cause error
}

// WithHint attaches remediation context and returns the receiver so it can be
// chained onto a constructor. No-op on nil.
func (e *Error) WithHint(format string, args ...any) *Error {
	if e == nil {
		return nil
	}
	e.Hint = fmt.Sprintf(format, args...)
	return e
}

// Error satisfies the error interface.
func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return e.Message
}

// Unwrap returns the original cause, if any.
func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

// ExitCode returns the process exit code associated with this error's Kind.
func (e *Error) ExitCode() int {
	if e == nil {
		return 0
	}
	return e.Kind.ExitCode()
}

// New builds a classified error without an underlying cause.
func New(kind Kind, format string, args ...any) *Error {
	return &Error{
		Kind:    kind,
		Message: fmt.Sprintf(format, args...),
	}
}

// Wrap builds a classified error that preserves the underlying cause.
// errors.Is(wrapped, sentinel) keeps working through the chain.
func Wrap(kind Kind, cause error, format string, args ...any) *Error {
	msg := fmt.Sprintf(format, args...)
	if msg == "" && cause != nil {
		msg = cause.Error()
	}
	return &Error{
		Kind:    kind,
		Message: msg,
		cause:   cause,
	}
}

// Validation is the common shortcut for command-body argument checks.
func Validation(format string, args ...any) *Error {
	return New(KindValidation, format, args...)
}

// AuthRequired indicates no token is configured.
func AuthRequired(message string) *Error {
	return New(KindAuthRequired, "%s", message)
}

// AuthInvalid indicates the configured token was rejected by the server.
func AuthInvalid(cause error, message string) *Error {
	e := Wrap(KindAuthInvalid, cause, "%s", message)
	e.HTTPStatus = 401
	return e
}

// NotFound indicates the resource the user asked for doesn't exist.
func NotFound(format string, args ...any) *Error {
	return New(KindNotFound, format, args...)
}

// RateLimited indicates the caller is being throttled by the API.
func RateLimited(cause error) *Error {
	e := Wrap(KindRateLimited, cause, "rate limit exceeded; please retry later")
	e.HTTPStatus = 429
	return e
}
