package main

import (
	"bytes"
	"encoding/json"
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
// stdout, stderr, and the process exit code. VC_TOKEN is force-cleared
// so tests that assert auth-required paths behave consistently.
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
		"NO_COLOR=1",
		// Clear CI so Interactive() logic isn't skewed by the outer test env.
		"CI=",
		"BUILD_NUMBER=",
		"RUN_ID=",
	)
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

// ----- VC_TOKEN shadowing guardrail -----

// homeWithToken returns a fresh HOME containing a vulncheck.yaml that holds
// the given token, so tests can create the "logged in AND VC_TOKEN set" state.
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
	if !strings.Contains(stderr, "unset") {
		t.Errorf("stderr should tell the user to unset VC_TOKEN; got %q", stderr)
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

	stdout, _, exit := runCLIHome(t, home, "vulncheck_env_token_value", "auth", "status", "--json")
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
}

// token_shadowed must stay absent when only VC_TOKEN is set — that is the
// normal CI shape and it must not look like a misconfiguration.
func TestContractAuthStatusNoShadowingInCIShape(t *testing.T) {
	stdout, _, exit := runCLIHome(t, t.TempDir(), "vulncheck_env_token_value", "auth", "status", "--json")
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
	stdout, _, _ := runCLIHome(t, t.TempDir(), "", "indices", "list", "--json")
	m := mustJSON(t, stdout)
	errObj, ok := m["error"].(map[string]any)
	if !ok {
		t.Fatalf("expected an error envelope, got %v", m)
	}
	if _, present := errObj["hint"]; present {
		t.Errorf("hint should be omitted when there is no token source to name; got %v", errObj)
	}
}
