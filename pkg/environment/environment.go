package environment

import "os"

type Environment struct {
	Name   string
	Values []string
	API    string
	WEB    string
}

// Example usage: VC_ENV=custom VC_API=https://custom-api.com VC_WEB=https://custom-web.com vulncheck ...
var Environments = []Environment{
	{
		Name:   "production",
		Values: []string{"production", "prod"},
		API:    "https://api.vulncheck.com",
		WEB:    "https://console.vulncheck.com",
	},
	{
		Name:   "development",
		Values: []string{"development", "dev", "local"},
		API:    "http://localhost:8000",
		WEB:    "http://localhost:3000",
	},
	{
		Name:   "custom",
		Values: []string{"custom"},
		API:    "",
		WEB:    "",
	},
}

var Env = Environments[0]

func Init() {
	envVar := os.Getenv("VC_ENV")
	for _, env := range Environments {
		for _, value := range env.Values {
			if value == envVar {
				Env = env
				break
			}
		}
	}
	if apiOverride := os.Getenv("VC_API"); apiOverride != "" {
		Env.API = apiOverride
	}
	if webOverride := os.Getenv("VC_WEB"); webOverride != "" {
		Env.WEB = webOverride
	}
}
