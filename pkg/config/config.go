package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/vulncheck-oss/cli/pkg/environment"
	"log"
	"os"
	"strings"
)

type Environment struct {
	Name string
	API  string
	WEB  string
}

type Config struct {
	Token      string
	IndicesDir string
}

func Init() {
	environment.Init()
}

// newViper returns an isolated viper instance. Deliberately NOT the package
// -level global: viper.Set writes the highest-precedence "override" layer, so
// a single saveConfig on the global would make every later loadConfig in the
// process return the written values regardless of what is on disk. That made
// the save/load round-trip untestable and let a write in one command leak into
// a read in another.
func newViper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("vulncheck")
	v.SetConfigType("yaml")
	return v
}

func loadConfig() (*Config, error) {
	dir, err := Dir()
	if err != nil {
		return nil, err
	}

	v := newViper()
	v.AddConfigPath(dir)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	// Retrieve values from the configuration
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return &config, nil
}

func saveConfig(config *Config) error {
	dir, err := Dir()
	if err != nil {
		return err
	}
	v := newViper()
	v.AddConfigPath(dir)
	v.SetConfigPermissions(0600)
	v.Set("Token", config.Token)
	v.Set("IndicesDir", config.IndicesDir)
	return v.WriteConfigAs(fmt.Sprintf("%s/vulncheck.yaml", dir))
}

// mutateConfig applies fn to the config currently on disk and writes it back,
// so a change to one field never silently drops the others. saveConfig writes
// every field it knows about, so callers must never hand it a fresh &Config{}
// built from a single value.
func mutateConfig(fn func(*Config)) error {
	config, err := loadConfig()
	if err != nil {
		config = &Config{}
	}
	fn(config)
	return saveConfig(config)
}

func Dir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	dir := fmt.Sprintf("%s/.config/vulncheck", homeDir)
	// 0700: the directory holds the API token (in vulncheck.yaml, itself
	// 0600). On shared / multi-user systems any other-readable bit on the
	// parent dir lets a colocated user enumerate / race the file.
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return "", err
		}
	} else if err == nil {
		// Existing dir might have been created with the old 0755 mode
		// (or by an unrelated tool). Best-effort tighten; ignore errors
		// because we don't own the dir necessarily.
		_ = os.Chmod(dir, 0700)
	}
	return dir, nil
}

func IndicesDir() (string, error) {

	config, err := loadConfig()
	if err == nil && config.IndicesDir != "" {
		dir := config.IndicesDir
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", fmt.Errorf("failed to create indices directory: %w", err)
		}
		return dir, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	dir := fmt.Sprintf("%s/.config/vulncheck/indices", homeDir)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create indices directory: %w", err)
	}
	return dir, nil
}

func SetIndicesDir(dir string) error {
	return mutateConfig(func(c *Config) { c.IndicesDir = dir })
}

func HasConfig() bool {
	_, err := loadConfig()
	return err == nil
}

// EnvToken is the legacy token variable, consulted first. See Resolve.
const EnvToken = "VC_TOKEN"

// EnvTokenAPI is the name shared with the VulnCheck SDKs and MCP server, and the
// one user-facing copy should teach. EnvToken keeps precedence over it.
const EnvTokenAPI = "VULNCHECK_API_TOKEN"

// TokenSource identifies where the active token came from. The string values
// are the agent-facing contract emitted by `auth status --json`.
type TokenSource string

const (
	// SourceNone means no usable token was found in either location.
	SourceNone TokenSource = ""
	// SourceEnv means the token came from one of the token environment
	// variables; Resolution.EnvVar reports which.
	SourceEnv TokenSource = "env"
	// SourceConfig means the token came from vulncheck.yaml.
	SourceConfig TokenSource = "config"
)

// Resolution is the outcome of resolving a token from all sources. It is the
// single place precedence is decided; Token/HasToken/TokenFromEnv are thin
// wrappers so they can never drift apart.
type Resolution struct {
	// Token is the token that will actually be sent to the API, or "" if none.
	Token string
	// Source says which location Token came from.
	Source TokenSource
	// Shadowed reports that a token environment variable is overriding a
	// *different* token saved in vulncheck.yaml. Only true when both are
	// present and their values differ — identical values cannot confuse
	// anyone, and CI (env set, no config file) is never shadowed, so this
	// stays quiet in automation.
	Shadowed bool
	// ConfigToken is the token stored in vulncheck.yaml, whether or not it
	// won. Used to report shadowing; never render it to the user.
	ConfigToken string
	// EnvVar names the variable Token came from, or "" when Source is not
	// SourceEnv. Messages must use this rather than a fixed constant, or they
	// will name the wrong variable to half the users who see them.
	EnvVar string
	// EnvVarsSet names every token variable holding a usable value, in
	// precedence order; EnvVarsSet[0] == EnvVar. Hints must clear all of them —
	// clearing only the winner hands the win to the next.
	EnvVarsSet []string
}

// FromEnv reports whether the active token came from the environment.
func (r Resolution) FromEnv() bool { return r.Source == SourceEnv }

// UnsetHint names every variable to clear, as an embeddable fragment. Shell
// neutral: `unset` is not a command on Windows, which the CLI ships binaries for.
func (r Resolution) UnsetHint() string {
	switch n := len(r.EnvVarsSet); n {
	case 0:
		return ""
	case 1:
		return "clear " + r.EnvVarsSet[0] + " from your environment"
	default:
		return "clear " + strings.Join(r.EnvVarsSet[:n-1], ", ") +
			" and " + r.EnvVarsSet[n-1] + " from your environment"
	}
}

// Resolve determines the active token. Precedence: EnvToken, EnvTokenAPI, then
// vulncheck.yaml; a value failing ValidToken is absent in every position.
// The only place the CLI reads a token environment variable. Reading one
// elsewhere duplicates the ValidToken rule and the precedence order.
func Resolve() Resolution {
	var res Resolution

	if config, err := loadConfig(); err == nil && ValidToken(config.Token) {
		res.ConfigToken = config.Token
	}

	// Keep scanning past the winner — see EnvVarsSet.
	for _, name := range []string{EnvToken, EnvTokenAPI} {
		env := os.Getenv(name)
		if !ValidToken(env) {
			continue
		}
		res.EnvVarsSet = append(res.EnvVarsSet, name)
		if res.Source == SourceEnv {
			continue
		}
		res.Token, res.Source, res.EnvVar = env, SourceEnv, name
		res.Shadowed = res.ConfigToken != "" && res.ConfigToken != env
	}
	if res.Source == SourceEnv {
		return res
	}

	if res.ConfigToken != "" {
		res.Token = res.ConfigToken
		res.Source = SourceConfig
	}
	return res
}

func TokenFromEnv() bool {
	return Resolve().FromEnv()
}

func Token() string {
	return Resolve().Token
}

func HasToken() bool {
	return Resolve().Token != ""
}

// SaveToken writes token to vulncheck.yaml, preserving every other setting.
func SaveToken(token string) error {
	return mutateConfig(func(c *Config) { c.Token = token })
}

// RemoveToken clears the saved token, preserving every other setting.
// Note this only clears the *config file*; if a token environment variable is
// set the caller remains authenticated. Check Resolve().FromEnv() before
// claiming otherwise.
func RemoveToken() error {
	return mutateConfig(func(c *Config) { c.Token = "" })
}

func ValidToken(token string) bool {
	if strings.TrimSpace(token) == "" {
		return false
	}
	/* TODO: update this to work with JWT as well
	if !strings.HasPrefix(token, "vulncheck_") {
		return false
	}
	if len(token) != 74 {
		return false
	}
	*/
	return true
}

// IsCI based on https://github.com/watson/ci-info/blob/HEAD/index.js
func IsCI() bool {
	return os.Getenv("CI") != "" || // GitHub Actions, Travis CI, CircleCI, Cirrus CI, GitLab CI, AppVeyor, CodeShip, dsari
		os.Getenv("BUILD_NUMBER") != "" || // Jenkins, TeamCity
		os.Getenv("RUN_ID") != "" // TaskCluster, dsari
}
