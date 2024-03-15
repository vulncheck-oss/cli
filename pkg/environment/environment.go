package environment

import "os"

type Environment struct {
	Name   string
	Values []string
	API    string
	WEB    string
}

var Environments = []Environment{
	{
		Name:   "production",
		Values: []string{"production", "prod"},
		API:    "https://api.vulncheck.com",
		WEB:    "https://vulncheck.com",
	},
	{
		Name:   "staging",
		Values: []string{"staging", "stage"},
		API:    "https://api.staging.vulncheck.com",
		WEB:    "https://staging.vulncheck.com",
	},
	{
		Name:   "development",
		Values: []string{"development", "dev", "local"},
		API:    "http://localhost:8000",
		WEB:    "http://localhost:3000",
	},
}

// TODO: move this to [0] before initial release - and remove staging from the array
var Env = Environments[2]

func Init() {
	envVar := os.Getenv("VC_ENV")
	for _, env := range Environments {
		for _, value := range env.Values {
			if value == envVar {
				Env = env
				return
			}
		}
	}
}
