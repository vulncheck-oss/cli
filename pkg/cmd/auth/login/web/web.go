package web

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/huh/spinner"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/environment"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/login"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

/**
step 1. generate an inquiry.
step 2. prompt the user to visit the inquiry URL.
step 3. loop and sleep waiting for an inquiry response.
*/

type Inquiry struct {
	Hash       string
	Token      string
	Name       string
	IP         string
	Agent      string
	Location   string
	Coordinate string
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type InquiryResponse struct {
	Benchmark float64 `json:"_benchmark"`
	Message   string  `json:"message"`
	Data      Inquiry `json:"data"`
}

type InquiryPingResponse struct {
	Benchmark float64 `json:"_benchmark"`
	Data      Inquiry `json:"data"`
}

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "web",
		Short: i18n.C.AuthLoginWeb,
		RunE:  CmdWeb,
	}
}

func CmdWeb(cmd *cobra.Command, args []string) error {
	var responseJSON *InquiryResponse
	response, err := session.Connect(config.Token()).Form("name", GetName()).Request("POST", "/inquiry")
	if err != nil {
		return err
	}
	defer response.Body.Close()
	_ = json.NewDecoder(response.Body).Decode(&responseJSON)

	ui.Info("Attempting to launch vulncheck.com in your browser...")
	if err := browser.OpenURL(fmt.Sprintf("%s/inquiry/%s", environment.Env.WEB, responseJSON.Data.Hash)); err != nil {
		return err
	}

	var errorResponse error
	var pingResponse *InquiryPingResponse

	_ = spinner.New().
		Style(ui.Pantone).
		Title(" Awaiting Verification...").Action(func() {

		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		timeout := time.After(30 * time.Second)

		for {
			select {
			case <-ticker.C:
				var responsePing *InquiryPingResponse
				response, err := session.Connect(config.Token()).Request("GET", fmt.Sprintf("/inquiry/ping/%s", responseJSON.Data.Hash))
				if err != nil {
					errorResponse = err
					return
				}
				defer response.Body.Close()
				_ = json.NewDecoder(response.Body).Decode(&responsePing)
				if config.ValidToken(responsePing.Data.Token) {
					pingResponse = responsePing
					return
				}
			case <-timeout:
				return
			}
		}

	}).Run()

	if errorResponse != nil {
		return errorResponse
	}

	if pingResponse != nil {
		return login.SaveToken(pingResponse.Data.Token)
	}
	return nil
}

// GetName returns the ComputerName and/or hostname of the machine
func GetName() string {
	var out []byte
	var err error

	if strings.HasPrefix(runtime.GOOS, "darwin") {
		out, err = exec.Command("scutil", "--get", "ComputerName").Output()
	} else {
		out, err = exec.Command("hostname").Output()
	}
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(out))
}
