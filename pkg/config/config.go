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

func loadConfig() (*Config, error) {
	dir, err := Dir()
	if err != nil {
		return nil, err
	}

	viper.AddConfigPath(dir)
	viper.SetConfigName("vulncheck")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Retrieve values from the configuration
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return &config, nil
}

func saveConfig(config *Config) error {
	dir, err := Dir()
	if err != nil {
		return err
	}
	viper.AddConfigPath(dir)
	viper.SetConfigName("vulncheck")
	viper.SetConfigPermissions(0600)
	viper.SetConfigType("yaml")
	viper.Set("Token", config.Token)
	viper.Set("IndicesDir", config.IndicesDir)
	return viper.WriteConfigAs(fmt.Sprintf("%s/vulncheck.yaml", dir))
}

func Dir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	dir := fmt.Sprintf("%s/.config/vulncheck", homeDir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", err
		}
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
	config, err := loadConfig()
	if err != nil {
		config = &Config{}
	}
	config.IndicesDir = dir
	return saveConfig(config)
}

func HasConfig() bool {
	_, err := loadConfig()
	return err == nil
}

func TokenFromEnv() bool {
	return ValidToken(os.Getenv("VC_TOKEN"))
}

func Token() string {

	token := os.Getenv("VC_TOKEN")

	if token != "" && ValidToken(token) {
		return token
	}

	config, err := loadConfig()
	if err != nil {
		return ""
	}
	return config.Token
}

func HasToken() bool {
	return Token() != ""
}

func SaveToken(token string) error {
	config := &Config{
		Token: token,
	}
	return saveConfig(config)
}

func RemoveToken() error {
	config := &Config{
		Token: "",
	}
	return saveConfig(config)
}

func ValidToken(token string) bool {
	if !strings.HasPrefix(token, "vulncheck_") {
		return false
	}
	if len(token) != 74 {
		return false
	}
	return true
}

// IsCI based on https://github.com/watson/ci-info/blob/HEAD/index.js
func IsCI() bool {
	return os.Getenv("CI") != "" || // GitHub Actions, Travis CI, CircleCI, Cirrus CI, GitLab CI, AppVeyor, CodeShip, dsari
		os.Getenv("BUILD_NUMBER") != "" || // Jenkins, TeamCity
		os.Getenv("RUN_ID") != "" // TaskCluster, dsari
}
