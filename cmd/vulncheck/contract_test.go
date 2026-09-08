package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// binPath points at a freshly-built vulncheck binary for the duration
// of this test package's runtime. See TestMain.
var binPath string

// isolatedHome is a temp dir set as $HOME + $XDG_CONFIG_HOME while tests
// run, so the binary never sees the developer's real ~/.config/vulncheck.
var isolatedHome string

func TestMain(m *testing.M) {
	tmp, err := os.MkdirTemp("", "vulncheck-contract-*")
	if err != nil {
		panic(err)
	}
	defer func() { _ = os.RemoveAll(tmp) }()

	// On Windows `go build -o vulncheck` writes to `vulncheck.exe`;
	// exec.Command would then fail with "file not found" on the bare name.
	// Match Go's naming so both build and later Run land on the same path.
	binName := "vulncheck"
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}
	binPath = filepath.Join(tmp, binName)
	build := exec.Command("go", "build", "-o", binPath, ".")
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		panic("failed to build binary for contract tests: " + err.Error())
	}

	isolatedHome = filepath.Join(tmp, "home")
	if err := os.MkdirAll(isolatedHome, 0o755); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

// runCLI shells out to the built binary with the given args and returns
// stdout, stderr, and the process exit code. Both token variables are
// force-cleared so tests that assert auth-required paths behave consistently.
func runCLI(t *testing.T, args ...string) (stdout, stderr string, exitCode int) {
	return runCLIEnv(t, "", args...)
}

// runCLIAuthed sets a syntactically valid VC_TOKEN so PersistentPreRunE's
// auth gate passes; use this for tests that want to reach a command's
// own validation logic (which fires AFTER the auth check).
func runCLIAuthed(t *testing.T, args ...string) (stdout, stderr string, exitCode int) {
	return runCLIEnv(t, "test-token-not-real", args...)
}

func runCLIEnv(t *testing.T, token string, args ...string) (stdout, stderr string, exitCode int) {
	return runCLIHome(t, isolatedHome, token, args...)
}

// runCLIHome is runCLIEnv with an explicit HOME, for tests that need to seed
// their own vulncheck.yaml without disturbing the package-wide isolatedHome.
func runCLIHome(t *testing.T, home, token string, args ...string) (stdout, stderr string, exitCode int) {
	t.Helper()
	return runCLIHomeEnv(t, home, nil, token, args...)
}

// runCLIHomeEnv is runCLIHome with extra environment entries appended after
// the fixture, so a test can set both token variables at once. Later entries
// win, so extra overrides the defaults below.
func runCLIHomeEnv(t *testing.T, home string, extra []string, token string, args ...string) (stdout, stderr string, exitCode int) {
	t.Helper()
	cmd := exec.Command(binPath, args...)
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	cmd.Env = append(os.Environ(),
		// HOME on Unix, USERPROFILE on Windows — both point os.UserHomeDir
		// at our isolated tree so the tests never touch the developer's
		// real ~/.config/vulncheck or %APPDATA%.
		"HOME="+home,
		"USERPROFILE="+home,
		"XDG_CONFIG_HOME="+home,
		"VC_TOKEN="+token,
		// runCLI passes token == "", which falls through to the next variable:
		// without this a developer with VULNCHECK_API_TOKEN exported has their
		// real token satisfy every auth-required test. CI never exports one,
		// so it presents as developer-machine-only flakiness.
		"VULNCHECK_API_TOKEN=",
		"NO_COLOR=1",
		// Clear CI so Interactive() logic isn't skewed by the outer test env.
		"CI=",
		"BUILD_NUMBER=",
		"RUN_ID=",
	)
	cmd.Env = append(cmd.Env, extra...)
	_ = cmd.Run()
	return outBuf.String(), errBuf.String(), cmd.ProcessState.ExitCode()
}

// mustJSON parses stdout as JSON and fails the test if it isn't valid.
func mustJSON(t *testing.T, out string) map[string]any {
	t.Helper()
	var m map[string]any
	if err := json.Unmarshal([]byte(out), &m); err != nil {
		t.Fatalf("stdout was not valid JSON: %v\nstdout: %q", err, out)
	}
	return m
}

// ----- version -----

func TestContractVersionJSONShape(t *testing.T) {
	stdout, _, exit := runCLI(t, "version", "--json")
	if exit != 0 {
		t.Fatalf("exit = %d, want 0", exit)
	}
	m := mustJSON(t, stdout)
	if sv, ok := m["schema_version"].(float64); !ok || sv != 1 {
		t.Errorf("schema_version = %v, want 1", m["schema_version"])
	}
	for _, key := range []string{"version", "changelog_url"} {
		if _, ok := m[key]; !ok {
			t.Errorf("version payload missing %q", key)
		}
	}
}

// ----- commands (capability discovery) -----

func TestContractCommandsDumpIsUsableAndStable(t *testing.T) {
	stdout, _, exit := runCLI(t, "commands")
	if exit != 0 {
		t.Fatalf("exit = %d, want 0", exit)
	}
	m := mustJSON(t, stdout)
	if sv, ok := m["schema_version"].(float64); !ok || sv != 1 {
		t.Fatalf("schema_version = %v, want 1", m["schema_version"])
	}
	root, ok := m["root"].(map[string]any)
	if !ok {
		t.Fatal("root object missing")
	}
	if root["name"] != "vulncheck" {
		t.Errorf("root.name = %v, want vulncheck", root["name"])
	}

	// Root's inherited flags must expose the global agentic surface.
	inherited, _ := root["inherited_flags"].([]any)
	got := map[string]bool{}
	for _, f := range inherited {
		if fm, ok := f.(map[string]any); ok {
			got[fm["name"].(string)] = true
		}
	}
	for _, name := range []string{"json", "quiet", "no-color", "no-interactive", "help"} {
		if !got[name] {
			t.Errorf("root.inherited_flags missing %q; got %v", name, got)
		}
	}

	// Every top-level command we care about is present.
	subs, _ := root["subcommands"].([]any)
	names := map[string]bool{}
	for _, s := range subs {
		if sm, ok := s.(map[string]any); ok {
			names[sm["name"].(string)] = true
		}
	}
	for _, want := range []string{"scan", "cpe", "purl", "auth", "token", "indices", "index", "advisory", "offline", "backup", "version"} {
		if !names[want] {
			t.Errorf("commands dump missing top-level %q; got %v", want, names)
		}
	}
}

// ----- auth status (no token, exit 0 with JSON body) -----

func TestContractAuthStatusJSONNoToken(t *testing.T) {
	stdout, _, exit := runCLI(t, "auth", "status", "--json")
	if exit != 0 {
		t.Fatalf("exit = %d, want 0 (agents dispatch on .authenticated)", exit)
	}
	m := mustJSON(t, stdout)
	if sv, _ := m["schema_version"].(float64); sv != 1 {
		t.Errorf("schema_version = %v, want 1", m["schema_version"])
	}
	if v, ok := m["authenticated"].(bool); !ok || v {
		t.Errorf("authenticated = %v, want false", m["authenticated"])
	}
	if _, ok := m["reason"].(string); !ok {
		t.Error("expected .reason to explain why not authenticated")
	}
}

// ----- error envelope: auth_required -----

func TestContractAuthRequiredErrorEnvelope(t *testing.T) {
	stdout, stderr, exit := runCLI(t, "token", "list", "--json")
	if exit != 3 {
		t.Fatalf("exit = %d, want 3 (auth_required)\nstderr: %s", exit, stderr)
	}
	m := mustJSON(t, stdout)
	if sv, _ := m["schema_version"].(float64); sv != 1 {
		t.Errorf("schema_version = %v, want 1", m["schema_version"])
	}
	err, ok := m["error"].(map[string]any)
	if !ok {
		t.Fatal("expected .error object")
	}
	if err["code"] != "auth_required" {
		t.Errorf("error.code = %v, want auth_required", err["code"])
	}
	if _, ok := err["message"].(string); !ok {
		t.Error("error.message should be a string")
	}
}

// ----- error envelope: validation -----

func TestContractValidationErrorEnvelope(t *testing.T) {
	// cpe requires a positional arg; the auth check runs first so we
	// need to satisfy it with a syntactically valid token before we
	// reach the arg validation.
	stdout, _, exit := runCLIAuthed(t, "cpe", "--json")
	if exit != 2 {
		t.Fatalf("exit = %d, want 2 (validation)", exit)
	}
	m := mustJSON(t, stdout)
	if sv, _ := m["schema_version"].(float64); sv != 1 {
		t.Errorf("schema_version = %v, want 1", m["schema_version"])
	}
	err, _ := m["error"].(map[string]any)
	if err == nil || err["code"] != "validation" {
		t.Errorf("error.code = %v, want validation", err)
	}
}

// ----- stdout / stderr discipline: nothing but JSON on stdout in --json mode -----

func TestContractJSONModeStdoutIsPureJSON(t *testing.T) {
	// Trigger the auth_required path — this exercises PersistentPreRunE's
	// auth check + the error envelope render.
	stdout, _, _ := runCLI(t, "token", "list", "--json")
	if !strings.HasPrefix(stdout, "{") {
		t.Fatalf("stdout should be pure JSON, got: %q", stdout)
	}
	if _, err := jsonRoundtrip(stdout); err != nil {
		t.Fatalf("stdout should parse as a single JSON document: %v\nstdout: %q", err, stdout)
	}
}

// ----- -h shorthand works at root AND subcommand level -----

func TestContractShortHelpFlag(t *testing.T) {
	for _, args := range [][]string{
		{"-h"},
		{"scan", "-h"},
		{"token", "list", "-h"},
	} {
		stdout, stderr, exit := runCLI(t, args...)
		if exit != 0 {
			t.Fatalf("%v: exit = %d, want 0", args, exit)
		}
		out := stdout + stderr
		if !strings.Contains(out, "Usage:") {
			t.Errorf("%v: expected usage output, got: %q", args, out)
		}
	}
}

// ----- --no-interactive gates interactive commands cleanly -----

func TestContractNoInteractiveRefusesTokenBrowse(t *testing.T) {
	stdout, _, exit := runCLIAuthed(t, "token", "browse", "--no-interactive", "--json")
	if exit != 2 {
		t.Fatalf("exit = %d, want 2 (validation)", exit)
	}
	m := mustJSON(t, stdout)
	err, _ := m["error"].(map[string]any)
	if err == nil || err["code"] != "validation" {
		t.Errorf("error.code = %v, want validation", err)
	}
}

func jsonRoundtrip(s string) (any, error) {
	var v any
	err := json.Unmarshal([]byte(s), &v)
	return v, err
}

// ----- scan --sbom-only --json (closes vulncheck-oss/cli#208) -----

// TestContractScanSbomOnlyJSON locks in the envelope shape emitted by
// scan when --sbom-only is set. This path does not hit the API, so we
// can drive it against a synthetic project directory without network.
func TestContractScanSbomOnlyJSON(t *testing.T) {
	proj := t.TempDir()
	if err := os.WriteFile(filepath.Join(proj, "go.mod"), []byte("module example.com/x\ngo 1.21\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	stdout, stderr, exit := runCLIAuthed(t, "scan", proj, "--sbom-only", "--json")
	if exit != 0 {
		t.Fatalf("exit = %d, want 0\nstderr: %s", exit, stderr)
	}
	m := mustJSON(t, stdout)
	if sv, _ := m["schema_version"].(float64); sv != 1 {
		t.Errorf("schema_version = %v, want 1", m["schema_version"])
	}
	if v, _ := m["sbom_only"].(bool); !v {
		t.Errorf("sbom_only = %v, want true", m["sbom_only"])
	}
	// Nothing else should appear on stdout — no vulnerabilities key when
	// the scan explicitly skipped the vuln lookup.
	if _, present := m["vulnerabilities"]; present {
		t.Errorf("sbom-only response should not include a vulnerabilities key; got %v", m)
	}
}

// ----- environment token guardrails -----

// homeWithToken returns a fresh HOME containing a vulncheck.yaml that holds
// the given token, so tests can create the "logged in AND an environment
// token set" state.
func homeWithToken(t *testing.T, token string) string {
	t.Helper()
	home := t.TempDir()
	dir := filepath.Join(home, ".config", "vulncheck")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "vulncheck.yaml"),
		[]byte("token: "+token+"\nindicesdir: \"\"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	return home
}

// unauthorizedAPI starts a stub API that rejects every request and returns the
// VC_API entry pointing the CLI at it. Any test whose fixture token reaches a
// *verification* call needs this: the tokens are syntactically valid, so the
// 401 has to come from a server, and without the override that server is
// production.
func unauthorizedAPI(t *testing.T) string {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"message":"unauthorized"}`))
	}))
	t.Cleanup(srv.Close)
	return "VC_API=" + srv.URL
}

// authorizedAPI starts a stub API that accepts every request with a /me
// payload, and returns the VC_API entry pointing the CLI at it. Needed by the
// tests that assert on *successful* auth output.
func authorizedAPI(t *testing.T) string {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"data":{"Name":"Test User","Email":"test@example.com"}}`))
	}))
	t.Cleanup(srv.Close)
	return "VC_API=" + srv.URL
}

// authorizedAPIExpecting is authorizedAPI plus an assertion on the exact
// Authorization header received. The permissive stub cannot tell a trimmed
// token from a padded one: net/http refuses a header value containing a
// newline outright, but a trailing space travels fine and the stub would
// accept it. Proving the trim means looking at what went over the wire.
func authorizedAPIExpecting(t *testing.T, wantToken string) string {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if got := req.Header.Get("Authorization"); got != "Bearer "+wantToken {
			t.Errorf("Authorization = %q, want %q", got, "Bearer "+wantToken)
		}
		_, _ = w.Write([]byte(`{"data":{"Name":"Test User","Email":"test@example.com"}}`))
	}))
	t.Cleanup(srv.Close)
	return "VC_API=" + srv.URL
}

// A stale VC_TOKEN silently overrode a saved token, so `auth login` verified
// the pasted value, skipped the save, and still reported success. It must now
// refuse rather than pretend.
func TestContractAuthLoginRefusesWhenEnvTokenSet(t *testing.T) {
	home := homeWithToken(t, "vulncheck_saved_token_value")

	// --json implies non-interactive, which would trip the CI guard first, so
	// exercise the plain path and read stderr.
	_, stderr, exit := runCLIHome(t, home, "vulncheck_env_token_value", "auth", "login", "token")
	if exit != 2 {
		t.Fatalf("exit = %d, want 2 (validation)\nstderr: %s", exit, stderr)
	}
	if !strings.Contains(stderr, "VC_TOKEN") {
		t.Errorf("stderr should name VC_TOKEN as the cause; got %q", stderr)
	}
	if !strings.Contains(stderr, "clear VC_TOKEN from your environment") {
		t.Errorf("stderr should tell the user how to stop VC_TOKEN winning; got %q", stderr)
	}

	// The saved token must be untouched — refusing is not a licence to write.
	b, err := os.ReadFile(filepath.Join(home, ".config", "vulncheck", "vulncheck.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), "vulncheck_saved_token_value") {
		t.Errorf("saved token should be unchanged after a refused login; file = %q", b)
	}
}

// `auth logout` used to revoke the *env* token server-side while deleting a
// different token from disk, then report success even though the caller was
// still authenticated. It must refuse instead of touching anything.
func TestContractAuthLogoutRefusesWhenEnvTokenSet(t *testing.T) {
	home := homeWithToken(t, "vulncheck_saved_token_value")

	_, stderr, exit := runCLIHome(t, home, "vulncheck_env_token_value", "auth", "logout")
	if exit != 2 {
		t.Fatalf("exit = %d, want 2 (validation)\nstderr: %s", exit, stderr)
	}
	if !strings.Contains(stderr, "VC_TOKEN") {
		t.Errorf("stderr should name VC_TOKEN; got %q", stderr)
	}

	b, err := os.ReadFile(filepath.Join(home, ".config", "vulncheck", "vulncheck.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), "vulncheck_saved_token_value") {
		t.Errorf("logout must not clear the saved token while refusing; file = %q", b)
	}
}

// auth status reports the source and the shadowing state so an agent can
// explain why a freshly saved credential appears to do nothing.
func TestContractAuthStatusReportsShadowing(t *testing.T) {
	home := homeWithToken(t, "vulncheck_saved_token_value")

	stdout, _, exit := runCLIHomeEnv(t, home, []string{unauthorizedAPI(t)},
		"vulncheck_env_token_value", "auth", "status", "--json")
	if exit != 0 {
		t.Fatalf("exit = %d, want 0", exit)
	}
	m := mustJSON(t, stdout)
	if m["token_source"] != "env" {
		t.Errorf("token_source = %v, want env", m["token_source"])
	}
	if v, _ := m["token_shadowed"].(bool); !v {
		t.Errorf("token_shadowed = %v, want true", m["token_shadowed"])
	}
	if m["token_env_var"] != "VC_TOKEN" {
		t.Errorf("token_env_var = %v, want VC_TOKEN", m["token_env_var"])
	}
}

// The CLI must accept VULNCHECK_API_TOKEN and report which variable it read.
// Runs through the verify-failed branch, where naming it matters most.
func TestContractAuthStatusAcceptsDocumentedEnvVar(t *testing.T) {
	stdout, _, exit := runCLIHomeEnv(t, t.TempDir(),
		[]string{"VULNCHECK_API_TOKEN=vulncheck_api_token_value", unauthorizedAPI(t)},
		"", "auth", "status", "--json")
	if exit != 0 {
		t.Fatalf("exit = %d, want 0", exit)
	}
	m := mustJSON(t, stdout)
	if m["token_source"] != "env" {
		t.Errorf("token_source = %v, want env", m["token_source"])
	}
	if m["token_env_var"] != "VULNCHECK_API_TOKEN" {
		t.Errorf("token_env_var = %v, want VULNCHECK_API_TOKEN", m["token_env_var"])
	}
}

// VC_TOKEN keeps winning, so no existing setup silently changes which
// credential it authenticates with when VULNCHECK_API_TOKEN is also present.
func TestContractLegacyEnvVarTakesPrecedence(t *testing.T) {
	stdout, _, exit := runCLIHomeEnv(t, t.TempDir(),
		[]string{"VULNCHECK_API_TOKEN=vulncheck_api_token_value", unauthorizedAPI(t)},
		"vulncheck_env_token_value", "auth", "status", "--json")
	if exit != 0 {
		t.Fatalf("exit = %d, want 0", exit)
	}
	if m := mustJSON(t, stdout); m["token_env_var"] != "VC_TOKEN" {
		t.Errorf("token_env_var = %v, want VC_TOKEN (legacy must keep winning)", m["token_env_var"])
	}
}

// Refusals must name the variable actually in use, not a hardcoded VC_TOKEN:
// advice to clear VC_TOKEN would leave a VULNCHECK_API_TOKEN user just as stuck.
func TestContractAuthLoginRefusalNamesDocumentedEnvVar(t *testing.T) {
	home := homeWithToken(t, "vulncheck_saved_token_value")

	_, stderr, exit := runCLIHomeEnv(t, home,
		[]string{"VULNCHECK_API_TOKEN=vulncheck_api_token_value"},
		"", "auth", "login", "token")
	if exit != 2 {
		t.Fatalf("exit = %d, want 2 (validation)\nstderr: %s", exit, stderr)
	}
	if !strings.Contains(stderr, "clear VULNCHECK_API_TOKEN from your environment") {
		t.Errorf("hint must name the variable in use; got %q", stderr)
	}
}

// With both variables set, a hint naming only the winner sends the user in a
// loop: clearing VC_TOKEN hands the win to VULNCHECK_API_TOKEN and the next
// attempt is refused all over again. Assert the joined form specifically — a
// test that merely looked for "VC_TOKEN" would pass against that bug.
func TestContractAuthLoginHintNamesEveryEnvToken(t *testing.T) {
	home := homeWithToken(t, "vulncheck_saved_token_value")

	_, stderr, exit := runCLIHomeEnv(t, home,
		[]string{"VULNCHECK_API_TOKEN=vulncheck_api_token_value"},
		"vulncheck_env_token_value", "auth", "login", "token")
	if exit != 2 {
		t.Fatalf("exit = %d, want 2 (validation)\nstderr: %s", exit, stderr)
	}
	if !strings.Contains(stderr, "clear VC_TOKEN and VULNCHECK_API_TOKEN from your environment") {
		t.Errorf("hint must name both variables at once; got %q", stderr)
	}
}

// token_shadowed must stay absent when only VC_TOKEN is set — that is the
// normal CI shape and it must not look like a misconfiguration.
func TestContractAuthStatusNoShadowingInCIShape(t *testing.T) {
	stdout, _, exit := runCLIHomeEnv(t, t.TempDir(), []string{unauthorizedAPI(t)},
		"vulncheck_env_token_value", "auth", "status", "--json")
	if exit != 0 {
		t.Fatalf("exit = %d, want 0", exit)
	}
	m := mustJSON(t, stdout)
	if _, present := m["token_shadowed"]; present {
		t.Errorf("token_shadowed should be omitted when nothing is shadowed; got %v", m)
	}
}

// The auth_required envelope gains a hint naming the token source, so the
// "re-login changes nothing" loop is self-diagnosing.
func TestContractAuthErrorEnvelopeHint(t *testing.T) {
	// No token anywhere: no source to name, so no hint.
	stdout, _, exit := runCLIHome(t, t.TempDir(), "", "indices", "list", "--json")
	if exit != 3 {
		t.Errorf("exit = %d, want 3 (auth)", exit)
	}
	m := mustJSON(t, stdout)
	errObj, ok := m["error"].(map[string]any)
	if !ok {
		t.Fatalf("expected an error envelope, got %v", m)
	}
	if _, present := errObj["hint"]; present {
		t.Errorf("hint should be omitted when there is no token source to name; got %v", errObj)
	}
}

// The auth_required hint must name whichever variable supplied the rejected
// token, so a user who exported only the documented name is not sent looking
// for a VC_TOKEN they never set.
func TestContractAuthErrorEnvelopeNamesDocumentedEnvVar(t *testing.T) {
	stdout, _, exit := runCLIHomeEnv(t, t.TempDir(),
		[]string{"VULNCHECK_API_TOKEN=vulncheck_api_token_value", unauthorizedAPI(t)},
		"", "indices", "list", "--json")

	// The exit code is what an agent branches on before it ever parses the
	// envelope; a rejected token must not arrive as a generic failure.
	if exit != 3 {
		t.Errorf("exit = %d, want 3 (auth)", exit)
	}
	m := mustJSON(t, stdout)
	errObj, ok := m["error"].(map[string]any)
	if !ok {
		t.Fatalf("expected an error envelope, got %v", m)
	}
	if errObj["code"] != "auth_invalid" {
		t.Errorf("code = %v, want auth_invalid", errObj["code"])
	}
	hint, _ := errObj["hint"].(string)
	if !strings.Contains(hint, "VULNCHECK_API_TOKEN") {
		t.Errorf("hint should name the variable the token came from; got %q", hint)
	}
	// No config file was written here, so a hint contrasting the environment
	// with one sends the user looking for a file that does not exist.
	if strings.Contains(hint, "config file") {
		t.Errorf("hint must not point at a config file that was never written; got %q", hint)
	}
}

// ----- both token variables set -----

// The two-variable state is reported by the unset command naming both, never
// by prose about the variable that lost. Prose would have to know which won,
// whether the other differs, whether a config file exists and whether that
// differs — a state space that produced a false "is being ignored" warning for
// the one setup this fallback exists to serve: a single credential exported
// under both names to feed the CLI and an SDK.
func TestContractBothEnvVarsSetStayQuiet(t *testing.T) {
	for _, tt := range []struct{ name, documented string }{
		{"identical tokens", "vulncheck_legacy_token"},
		{"differing tokens", "vulncheck_documented_token"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			stdout, stderr, exit := runCLIHomeEnv(t, t.TempDir(),
				[]string{"VULNCHECK_API_TOKEN=" + tt.documented, authorizedAPI(t)},
				"vulncheck_legacy_token", "auth", "status", "--json")
			if exit != 0 {
				t.Fatalf("exit = %d, want 0\nstderr: %s", exit, stderr)
			}
			m := mustJSON(t, stdout)
			// The variable in use is named; the one that lost is not a field.
			if m["token_env_var"] != "VC_TOKEN" {
				t.Errorf("token_env_var = %v, want VC_TOKEN", m["token_env_var"])
			}
			for _, gone := range []string{"token_env_var_ignored", "token_env_var_conflict"} {
				if _, present := m[gone]; present {
					t.Errorf("%s must not be part of the contract; got %v", gone, m)
				}
			}
			if strings.Contains(stderr, "ignored") {
				t.Errorf("no prose about the losing variable; got %q", stderr)
			}
		})
	}
}

// With both variables set, the remediation has to clear both — clearing only
// the winner hands the win to the other and refuses the next attempt all over
// again. The command therefore names a variable the sentence does not, so it
// has to say why, or it reads as the CLI having picked the wrong variable.
func TestContractUnsetHintNamesAndExplainsBothVars(t *testing.T) {
	_, stderr, exit := runCLIHomeEnv(t, t.TempDir(),
		[]string{"VULNCHECK_API_TOKEN=vulncheck_documented_token", authorizedAPI(t)},
		"vulncheck_legacy_token", "auth", "login", "token", "vulncheck_pasted_token")
	if exit != 2 {
		t.Fatalf("exit = %d, want 2 (validation)", exit)
	}
	if !strings.Contains(stderr, "clear VC_TOKEN and VULNCHECK_API_TOKEN from your environment") {
		t.Errorf("hint must name both variables at once; got %q", stderr)
	}
	// A working SDK setup is not a failure to be fixed by deleting the
	// credential the SDK needs.
	if strings.Contains(stderr, "would have no effect") && !strings.Contains(stderr, "nothing to do") {
		t.Errorf("refusal must not read as a broken setup; got %q", stderr)
	}
}

// A token carrying a trailing newline used to reach net/http, which refuses
// the Authorization header — surfacing as code "internal" (exit 1) naming no
// variable, so an agent branching on the auth codes treated a fixable
// credential problem as a CLI bug. A trailing space was worse: it survives
// the header and comes back as a bare 401 on a perfectly good token.
func TestContractWhitespacePaddedTokenIsUsable(t *testing.T) {
	for _, tt := range []struct{ name, padded string }{
		{"trailing newline", "vulncheck_api_token_value\n"},
		{"surrounding spaces", "  vulncheck_api_token_value  "},
	} {
		t.Run(tt.name, func(t *testing.T) {
			stdout, _, exit := runCLIHomeEnv(t, t.TempDir(),
				[]string{
					"VULNCHECK_API_TOKEN=" + tt.padded,
					authorizedAPIExpecting(t, "vulncheck_api_token_value"),
				},
				"", "auth", "status", "--json")
			if exit != 0 {
				t.Fatalf("exit = %d, want 0", exit)
			}
			m := mustJSON(t, stdout)
			if m["authenticated"] != true {
				t.Errorf("authenticated = %v, want true (padding should be trimmed); got %v",
					m["authenticated"], m)
			}
		})
	}
}

// The GitHub Actions help is the one place the CLI hands the user a snippet to
// paste. Naming only VULNCHECK_API_TOKEN sends anyone holding the VC_TOKEN
// secret that vulncheck-oss/action documents into a dead end: a missing secret
// expands to an empty string, so they land right back on "No token found".
func TestContractGitHubActionsHelpNamesBothSecrets(t *testing.T) {
	_, stderr, exit := runCLIHomeEnv(t, t.TempDir(),
		[]string{"GITHUB_ACTIONS=true"}, "", "indices", "list")
	if exit != 3 {
		t.Fatalf("exit = %d, want 3 (auth)", exit)
	}
	for _, want := range []string{"VULNCHECK_API_TOKEN", "VC_TOKEN"} {
		if !strings.Contains(stderr, want) {
			t.Errorf("Actions help must name %s; got %q", want, stderr)
		}
	}
}
