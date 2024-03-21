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
	Token string
}

func Init() {
	environment.Init()
}

func loadConfig() (*Config, error) {
	dir, err := configDir()
	if err != nil {
		return nil, err
	}

	viper.AddConfigPath(dir)
	viper.SetConfigName("vc")
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
	dir, err := configDir()
	if err != nil {
		return err
	}
	viper.AddConfigPath(dir)
	viper.SetConfigName("vc")
	viper.SetConfigType("yaml")
	viper.Set("Token", config.Token)
	return viper.WriteConfigAs(fmt.Sprintf("%s/vc.yaml", dir))
}

func configDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	dir := fmt.Sprintf("%s/.config/vc", homeDir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", err
		}
	}
	return dir, nil
}

func HasConfig() bool {

	_, err := loadConfig()
	if err != nil {
		return false
	}
	return true
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
