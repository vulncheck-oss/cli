package environment

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	originalVcEnv := os.Getenv("VC_ENV")
	originalVcApi := os.Getenv("VC_API")
	originalVcWeb := os.Getenv("VC_WEB")
	defer func() {
		_ = os.Setenv("VC_ENV", originalVcEnv)
		_ = os.Setenv("VC_API", originalVcApi)
		_ = os.Setenv("VC_WEB", originalVcWeb)
		Env = Environments[0]
	}()

	tests := []struct {
		name        string
		vcEnv       string
		expectedEnv string
		expectedAPI string
		expectedWEB string
	}{
		{
			name:        "Production environment",
			vcEnv:       "production",
			expectedEnv: "production",
			expectedAPI: "https://api.vulncheck.com",
			expectedWEB: "https://console.vulncheck.com",
		},
		{
			name:        "Production alias",
			vcEnv:       "prod",
			expectedEnv: "production",
			expectedAPI: "https://api.vulncheck.com",
			expectedWEB: "https://console.vulncheck.com",
		},
		{
			name:        "Development environment",
			vcEnv:       "development",
			expectedEnv: "development",
			expectedAPI: "http://localhost:8000",
			expectedWEB: "http://localhost:3000",
		},
		{
			name:        "Development alias",
			vcEnv:       "dev",
			expectedEnv: "development",
			expectedAPI: "http://localhost:8000",
			expectedWEB: "http://localhost:3000",
		},
		{
			name:        "Local alias",
			vcEnv:       "local",
			expectedEnv: "development",
			expectedAPI: "http://localhost:8000",
			expectedWEB: "http://localhost:3000",
		},
		{
			name:        "Custom environment",
			vcEnv:       "custom",
			expectedEnv: "custom",
			expectedAPI: "",
			expectedWEB: "",
		},
		{
			name:        "Unknown environment defaults to production",
			vcEnv:       "unknown",
			expectedEnv: "production",
			expectedAPI: "https://api.vulncheck.com",
			expectedWEB: "https://console.vulncheck.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = os.Unsetenv("VC_API")
			_ = os.Unsetenv("VC_WEB")
			Env = Environments[0]

			_ = os.Setenv("VC_ENV", tt.vcEnv)
			Init()

			if Env.Name != tt.expectedEnv {
				t.Errorf("Expected Env.Name to be '%s', got '%s'", tt.expectedEnv, Env.Name)
			}
			if Env.API != tt.expectedAPI {
				t.Errorf("Expected Env.API to be '%s', got '%s'", tt.expectedAPI, Env.API)
			}
			if Env.WEB != tt.expectedWEB {
				t.Errorf("Expected Env.WEB to be '%s', got '%s'", tt.expectedWEB, Env.WEB)
			}
		})
	}
}

func TestInitWithOverrides(t *testing.T) {
	originalVcEnv := os.Getenv("VC_ENV")
	originalVcApi := os.Getenv("VC_API")
	originalVcWeb := os.Getenv("VC_WEB")
	defer func() {
		_ = os.Setenv("VC_ENV", originalVcEnv)
		_ = os.Setenv("VC_API", originalVcApi)
		_ = os.Setenv("VC_WEB", originalVcWeb)
		Env = Environments[0]
	}()

	tests := []struct {
		name        string
		vcEnv       string
		vcApi       string
		vcWeb       string
		expectedAPI string
		expectedWEB string
	}{
		{
			name:        "API override only",
			vcEnv:       "production",
			vcApi:       "https://custom-api.com",
			vcWeb:       "",
			expectedAPI: "https://custom-api.com",
			expectedWEB: "https://console.vulncheck.com",
		},
		{
			name:        "WEB override only",
			vcEnv:       "production",
			vcApi:       "",
			vcWeb:       "https://custom-web.com",
			expectedAPI: "https://api.vulncheck.com",
			expectedWEB: "https://custom-web.com",
		},
		{
			name:        "Both API and WEB override",
			vcEnv:       "development",
			vcApi:       "https://staging-api.com",
			vcWeb:       "https://staging-web.com",
			expectedAPI: "https://staging-api.com",
			expectedWEB: "https://staging-web.com",
		},
		{
			name:        "Custom environment with overrides",
			vcEnv:       "custom",
			vcApi:       "https://custom-api.example.com",
			vcWeb:       "https://custom-web.example.com",
			expectedAPI: "https://custom-api.example.com",
			expectedWEB: "https://custom-web.example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Env = Environments[0]

			_ = os.Setenv("VC_ENV", tt.vcEnv)
			if tt.vcApi != "" {
				_ = os.Setenv("VC_API", tt.vcApi)
			} else {
				_ = os.Unsetenv("VC_API")
			}
			if tt.vcWeb != "" {
				_ = os.Setenv("VC_WEB", tt.vcWeb)
			} else {
				_ = os.Unsetenv("VC_WEB")
			}

			Init()

			if Env.API != tt.expectedAPI {
				t.Errorf("Expected Env.API to be '%s', got '%s'", tt.expectedAPI, Env.API)
			}
			if Env.WEB != tt.expectedWEB {
				t.Errorf("Expected Env.WEB to be '%s', got '%s'", tt.expectedWEB, Env.WEB)
			}
		})
	}
}
