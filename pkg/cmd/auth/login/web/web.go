package web

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/charmbracelet/huh/spinner"
	"github.com/octoper/go-ray"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/environment"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"net/http"
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
	ui.Info("Attempting to launch vulncheck.com in your browser...")

	var errorResponse error

	_ = spinner.New().
		Style(ui.Pantone).
		Title(" Awaiting Verification...").Action(func() {

		if err := browser.OpenURL(fmt.Sprintf("%s/inquiry/new", environment.Env.WEB)); err != nil {
			errorResponse = err
			return
		}

		hash, err := ListenForHash()
		if err != nil {
			errorResponse = err
			return
		}

		ray.Ray(hash)

	}).Run()

	if errorResponse != nil {
		return errorResponse
	}

	return nil
}

var Server *http.Server = &http.Server{Addr: ":8080"}

func ListenForHash() (string, error) {

	var hash string

	http.HandleFunc("/inquiry/", func(w http.ResponseWriter, r *http.Request) {
		hash := strings.TrimPrefix(r.URL.Path, "/inquiry/")
		if err := UpdateInquiry(hash); err != nil {
			fmt.Println(err)
		}
		redirect := fmt.Sprintf("%s/inquiry/%s", environment.Env.WEB, hash)
		http.Redirect(w, r, redirect, http.StatusMovedPermanently)
	})

	if err := Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return "", err
	}

	return hash, nil
}

func UpdateInquiry(hash string) error {
	var responseJSON *InquiryResponse
	response, err := session.Connect("").
		Form("name", GetName()).
		Request("PUT", fmt.Sprintf("/inquiry/%s", hash))

	if err != nil {
		return err
	}
	defer response.Body.Close()
	_ = json.NewDecoder(response.Body).Decode(&responseJSON)
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

func Hash() string {
	hasher := md5.New()
	hasher.Write([]byte(time.Now().String()))
	return hex.EncodeToString(hasher.Sum(nil))
}
