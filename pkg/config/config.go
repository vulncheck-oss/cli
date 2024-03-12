package config

import (
	"github.com/spf13/viper"
	"github.com/vulncheck-oss/cli/pkg/environment"
)

type Environment struct {
	Name string
	API  string
	WEB  string
}

func Init() {
	environment.Init()
	viper.SetDefault("Environment", "production")
}
