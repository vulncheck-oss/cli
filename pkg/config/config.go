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

// EnvToken is the environment variable consulted before the config file.
const EnvToken = "VC_TOKEN"

// TokenSource identifies where the active token came from. The string values
// are the agent-facing contract emitted by `auth status --json`.
type TokenSource string

const (
	// SourceNone means no usable token was found in either location.
	SourceNone TokenSource = ""
	// SourceEnv means the token came from the VC_TOKEN environment variable.
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
	// Shadowed reports that VC_TOKEN is overriding a *different* token saved
	// in vulncheck.yaml. Only true when both are present and their values
	// differ — identical values cannot confuse anyone, and CI (env set, no
	// config file) is never shadowed, so this stays quiet in automation.
	Shadowed bool
	// ConfigToken is the token stored in vulncheck.yaml, whether or not it
	// won. Used to report shadowing; never render it to the user.
	ConfigToken string
}

// FromEnv reports whether the active token came from the environment.
func (r Resolution) FromEnv() bool { return r.Source == SourceEnv }

// Resolve determines the active token. Precedence: VC_TOKEN, then the
// token in vulncheck.yaml. A value that fails ValidToken is treated as absent
// in both positions, so a blank or whitespace entry in either place falls
// through to the next source rather than being sent to the API.
func Resolve() Resolution {
	var res Resolution

	if config, err := loadConfig(); err == nil && ValidToken(config.Token) {
		res.ConfigToken = config.Token
	}

	if env := os.Getenv(EnvToken); ValidToken(env) {
		res.Token = env
		res.Source = SourceEnv
		res.Shadowed = res.ConfigToken != "" && res.ConfigToken != env
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
// Note this only clears the *config file*; if VC_TOKEN is set the caller
// remains authenticated. Check Resolve().FromEnv() before claiming otherwise.
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
