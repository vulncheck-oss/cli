package config

import (
	"os"
	"path/filepath"
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
		name         string
		env          string
		configToken  string
		wantToken    string
		wantSource   TokenSource
		wantShadowed bool
	}{
		{
			name:       "env only",
			env:        "env_token",
			wantToken:  "env_token",
			wantSource: SourceEnv,
		},
		{
			name:        "config only",
			configToken: "file_token",
			wantToken:   "file_token",
			wantSource:  SourceConfig,
		},
		{
			name:         "env shadows a different config token",
			env:          "env_token",
			configToken:  "file_token",
			wantToken:    "env_token",
			wantSource:   SourceEnv,
			wantShadowed: true,
		},
		{
			name:        "identical values are not shadowing",
			env:         "same_token",
			configToken: "same_token",
			wantToken:   "same_token",
			wantSource:  SourceEnv,
			// No possible confusion, so no warning should ever fire.
			wantShadowed: false,
		},
		{
			name:       "neither source",
			wantToken:  "",
			wantSource: SourceNone,
		},
		{
			name:        "blank env falls through to config",
			env:         "   ",
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
			t.Setenv(EnvToken, tt.env)

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
		})
	}
}

// The three public accessors are wrappers over Resolve; assert they cannot
// disagree with it the way the previous independent implementations could.
func TestAccessorsAgreeWithResolve(t *testing.T) {
	writeConfigToken(t, `"   "`) // whitespace: previously HasToken=true, CheckAuth=false
	t.Setenv(EnvToken, "")

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
