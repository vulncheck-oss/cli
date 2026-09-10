package config

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

// isolateHome points config.Dir() at a fresh temp directory and returns it.
//
// os.UserHomeDir reads USERPROFILE on Windows and HOME elsewhere, so setting
// only HOME leaves Windows resolving the runner's real profile — every test in
// the package then shares one config file and they contaminate each other.
// Always set both.
func isolateHome(t *testing.T) string {
	t.Helper()
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	return home
}

// writeConfigToken seeds a config file under an isolated home. Callers use
// Resolve() to observe it.
func writeConfigToken(t *testing.T, token string) {
	t.Helper()
	dir := filepath.Join(isolateHome(t), ".config", "vulncheck")
	if err := os.MkdirAll(dir, 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "vulncheck.yaml"),
		[]byte("token: "+token+"\n"), 0600); err != nil {
		t.Fatal(err)
	}
}

func TestResolvePrecedence(t *testing.T) {
	tests := []struct {
		name        string
		legacy      string // VC_TOKEN
		documented  string // VULNCHECK_API_TOKEN
		configToken string

		wantToken      string
		wantSource     TokenSource
		wantShadowed   bool
		wantEnvVar     string
		wantUnsetsBoth bool // UnsetTargets must name both variables
	}{
		{
			name:       "legacy only",
			legacy:     "env_token",
			wantToken:  "env_token",
			wantSource: SourceEnv,
			wantEnvVar: EnvToken,
		},
		{
			name:       "documented only",
			documented: "api_token",
			wantToken:  "api_token",
			wantSource: SourceEnv,
			wantEnvVar: EnvTokenAPI,
		},
		{
			// The design decision: the documented name is a fallback, so no
			// existing setup silently changes which credential it uses.
			name:       "legacy wins over documented",
			legacy:     "env_token",
			documented: "api_token",
			wantToken:  "env_token",
			wantSource: SourceEnv,
			wantEnvVar: EnvToken,
			// Unsetting only the winner hands the win to the other one.
			wantUnsetsBoth: true,
		},
		{
			// Same value, but unsetting one still leaves the other winning,
			// so the hint must clear both here too.
			name:           "both set to the same value",
			legacy:         "same_token",
			documented:     "same_token",
			wantToken:      "same_token",
			wantSource:     SourceEnv,
			wantEnvVar:     EnvToken,
			wantUnsetsBoth: true,
		},
		{
			// A whitespace-only value was never going to be sent, so it must
			// stay out of the unset hint.
			name:       "whitespace documented beside a valid legacy",
			legacy:     "env_token",
			documented: "   ",
			wantToken:  "env_token",
			wantSource: SourceEnv,
			wantEnvVar: EnvToken,
		},
		{
			name:       "blank legacy falls through to documented",
			legacy:     "   ",
			documented: "api_token",
			wantToken:  "api_token",
			wantSource: SourceEnv,
			wantEnvVar: EnvTokenAPI,
		},
		{
			// A Kubernetes secret mounted from a file, or a sourced .env,
			// arrives with the newline still attached. Untrimmed it reaches
			// net/http, which refuses the Authorization header outright.
			name:       "surrounding whitespace is trimmed off the env token",
			documented: "  api_token\n",
			wantToken:  "api_token",
			wantSource: SourceEnv,
			wantEnvVar: EnvTokenAPI,
		},
		{
			// Same credential, different spelling. Reporting a conflict here
			// would send the user unsetting a variable for no reason.
			name:        "whitespace-only difference from config is not shadowing",
			legacy:      "same_token\n",
			configToken: "same_token",
			wantToken:   "same_token",
			wantSource:  SourceEnv,
			wantEnvVar:  EnvToken,
		},
		{
			name:        "config only",
			configToken: "file_token",
			wantToken:   "file_token",
			wantSource:  SourceConfig,
		},
		{
			name:         "legacy shadows a different config token",
			legacy:       "env_token",
			configToken:  "file_token",
			wantToken:    "env_token",
			wantSource:   SourceEnv,
			wantShadowed: true,
			wantEnvVar:   EnvToken,
		},
		{
			name:         "documented shadows a different config token",
			documented:   "api_token",
			configToken:  "file_token",
			wantToken:    "api_token",
			wantSource:   SourceEnv,
			wantShadowed: true,
			wantEnvVar:   EnvTokenAPI,
		},
		{
			// The state UnsetTargets was written for: both variables set
			// *and* a different token in the config file. Clearing only the
			// winner hands the win to the other variable, so the hint has to
			// name both while still reporting the config file as shadowed.
			name:           "both env vars set beside a differing config token",
			legacy:         "env_token",
			documented:     "api_token",
			configToken:    "file_token",
			wantToken:      "env_token",
			wantSource:     SourceEnv,
			wantShadowed:   true,
			wantEnvVar:     EnvToken,
			wantUnsetsBoth: true,
		},
		{
			name:        "identical values are not shadowing",
			legacy:      "same_token",
			configToken: "same_token",
			wantToken:   "same_token",
			wantSource:  SourceEnv,
			// No possible confusion, so no warning should ever fire.
			wantShadowed: false,
			wantEnvVar:   EnvToken,
		},
		{
			name:       "neither source",
			wantToken:  "",
			wantSource: SourceNone,
		},
		{
			name:        "blank env falls through to config",
			legacy:      "   ",
			configToken: "file_token",
			wantToken:   "file_token",
			wantSource:  SourceConfig,
		},
		{
			name:        "whitespace-only config token is treated as absent",
			configToken: `"   "`,
			wantToken:   "",
			wantSource:  SourceNone,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.configToken != "" {
				writeConfigToken(t, tt.configToken)
			} else {
				isolateHome(t)
			}
			// Set both unconditionally: neither may leak in from the
			// developer's own environment.
			t.Setenv(EnvToken, tt.legacy)
			t.Setenv(EnvTokenAPI, tt.documented)

			got := Resolve()
			if got.Token != tt.wantToken {
				t.Errorf("Token = %q, want %q", got.Token, tt.wantToken)
			}
			if got.Source != tt.wantSource {
				t.Errorf("Source = %q, want %q", got.Source, tt.wantSource)
			}
			if got.Shadowed != tt.wantShadowed {
				t.Errorf("Shadowed = %v, want %v", got.Shadowed, tt.wantShadowed)
			}
			if got.EnvVar != tt.wantEnvVar {
				t.Errorf("EnvVar = %q, want %q", got.EnvVar, tt.wantEnvVar)
			}
			// EnvVar is named exactly when the env supplied the token; a
			// message that names a variable the token did not come from is
			// worse than naming none.
			if got.FromEnv() != (got.EnvVar != "") {
				t.Errorf("FromEnv()=%v disagrees with EnvVar=%q", got.FromEnv(), got.EnvVar)
			}
			var wantUnset []string
			switch {
			case tt.wantUnsetsBoth:
				wantUnset = []string{EnvToken, EnvTokenAPI}
			case tt.wantEnvVar != "":
				wantUnset = []string{tt.wantEnvVar}
			}
			if !slices.Equal(got.EnvVarsSet, wantUnset) {
				t.Errorf("EnvVarsSet = %v, want %v", got.EnvVarsSet, wantUnset)
			}
			// The winner must be the head of the list, or a caller reading
			// EnvVarsSet[0] names a variable the token did not come from.
			if len(got.EnvVarsSet) > 0 && got.EnvVarsSet[0] != got.EnvVar {
				t.Errorf("EnvVarsSet[0] = %q but EnvVar = %q", got.EnvVarsSet[0], got.EnvVar)
			}
			// The hint has to name every variable: a user who clears only some
			// of them still has one winning, and gets refused all over again.
			for _, name := range got.EnvVarsSet {
				if !strings.Contains(got.UnsetHint(), name) {
					t.Errorf("UnsetHint() = %q, does not name %s", got.UnsetHint(), name)
				}
			}
		})
	}
}

// The three public accessors are wrappers over Resolve; assert they cannot
// disagree with it the way the previous independent implementations could.
func TestAccessorsAgreeWithResolve(t *testing.T) {
	writeConfigToken(t, `"   "`) // whitespace: previously HasToken=true, CheckAuth=false
	t.Setenv(EnvToken, "")
	t.Setenv(EnvTokenAPI, "")

	res := Resolve()
	if HasToken() != (res.Token != "") {
		t.Errorf("HasToken()=%v but Resolve().Token=%q", HasToken(), res.Token)
	}
	if Token() != res.Token {
		t.Errorf("Token()=%q but Resolve().Token=%q", Token(), res.Token)
	}
	if TokenFromEnv() != res.FromEnv() {
		t.Errorf("TokenFromEnv()=%v but Resolve().FromEnv()=%v", TokenFromEnv(), res.FromEnv())
	}
	if res.EnvVar != "" {
		t.Errorf("EnvVar=%q with no env token set, want empty", res.EnvVar)
	}
	if HasToken() {
		t.Error("a whitespace-only config token must not count as a token")
	}
}

// saveConfig must not leak values into later loadConfig calls via viper's
// override layer, and loadConfig must reflect what is actually on disk.
func TestLoadConfigReadsDisk(t *testing.T) {
	home := isolateHome(t)

	if err := saveConfig(&Config{Token: "from_save"}); err != nil {
		t.Fatal(err)
	}

	p := filepath.Join(home, ".config", "vulncheck", "vulncheck.yaml")
	if err := os.WriteFile(p, []byte("token: from_disk\nindicesdir: \"\"\n"), 0600); err != nil {
		t.Fatal(err)
	}

	got, err := loadConfig()
	if err != nil {
		t.Fatal(err)
	}
	if got.Token != "from_disk" {
		t.Errorf("loadConfig() token = %q, want %q (viper override layer leaked)", got.Token, "from_disk")
	}
}

// SaveToken / RemoveToken change one field; they must not drop the others.
func TestTokenWritesPreserveIndicesDir(t *testing.T) {
	home := isolateHome(t)
	dir := filepath.Join(home, ".config", "vulncheck")
	if err := os.MkdirAll(dir, 0700); err != nil {
		t.Fatal(err)
	}
	p := filepath.Join(dir, "vulncheck.yaml")
	if err := os.WriteFile(p, []byte("token: old\nindicesdir: /data/my-indices\n"), 0600); err != nil {
		t.Fatal(err)
	}

	if err := SaveToken("new_token"); err != nil {
		t.Fatal(err)
	}
	c, err := loadConfig()
	if err != nil {
		t.Fatal(err)
	}
	if c.IndicesDir != "/data/my-indices" {
		t.Errorf("after SaveToken, IndicesDir = %q, want %q", c.IndicesDir, "/data/my-indices")
	}
	if c.Token != "new_token" {
		t.Errorf("after SaveToken, Token = %q, want %q", c.Token, "new_token")
	}

	if err := RemoveToken(); err != nil {
		t.Fatal(err)
	}
	c, err = loadConfig()
	if err != nil {
		t.Fatal(err)
	}
	if c.IndicesDir != "/data/my-indices" {
		t.Errorf("after RemoveToken, IndicesDir = %q, want %q", c.IndicesDir, "/data/my-indices")
	}
	if c.Token != "" {
		t.Errorf("after RemoveToken, Token = %q, want empty", c.Token)
	}
}

// A padded paste saved verbatim fails every run afterwards, and unlike an
// environment variable the user cannot see the padding to fix it.
func TestSaveTokenTrimsWhitespace(t *testing.T) {
	isolateHome(t)

	if err := SaveToken("  vulncheck_padded_token\n"); err != nil {
		t.Fatal(err)
	}
	c, err := loadConfig()
	if err != nil {
		t.Fatal(err)
	}
	if c.Token != "vulncheck_padded_token" {
		t.Errorf("saved token = %q, want it trimmed", c.Token)
	}
}

// The hint is embedded in sentences, so it must read as a fragment and must
// not name a shell: `unset` is not a command in any Windows shell, and the CLI
// ships Windows binaries.
func TestUnsetHintPhrasing(t *testing.T) {
	tests := []struct {
		names []string
		want  string
	}{
		{nil, ""},
		{[]string{EnvToken}, "clear VC_TOKEN from your environment"},
		{[]string{EnvToken, EnvTokenAPI}, "clear VC_TOKEN and VULNCHECK_API_TOKEN from your environment"},
	}
	for _, tt := range tests {
		if got := (Resolution{EnvVarsSet: tt.names}).UnsetHint(); got != tt.want {
			t.Errorf("UnsetHint() with %v = %q, want %q", tt.names, got, tt.want)
		}
	}
}
