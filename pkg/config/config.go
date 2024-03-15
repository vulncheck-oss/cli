package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/vulncheck-oss/cli/pkg/environment"
	"log"
	"os"
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
	viper.SetDefault("Environment", "production")
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

func SaveToken(token string) (string, error) {
	dir, err := configDir()
	if err != nil {
		return "", nil
	}
	viper.AddConfigPath(dir)
	viper.SetConfigName("vc")
	viper.SetConfigType("yaml")
	viper.Set("Token", token)
	err = viper.WriteConfigAs(fmt.Sprintf("%s/vc.yaml", dir))
	if err != nil {
		return "", err
	}
	return dir, nil
}
