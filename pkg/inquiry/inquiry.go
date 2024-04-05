package inquiry

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vulncheck-oss/cli/pkg/environment"
	"github.com/vulncheck-oss/cli/pkg/session"
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

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

type Response struct {
	Benchmark float64 `json:"_benchmark"`
	Message   string  `json:"message"`
	Data      Inquiry `json:"data"`
}

var Port = ":8678"

func ListenForHash() (string, error) {
	return ListenFor("inquiry", func(w http.ResponseWriter, r *http.Request, hash string) {
		if hash == "denied" {
			return // denied
		}
		if err := UpdateInquiry(hash); err != nil {
			fmt.Println(err)
		}
		redirect := fmt.Sprintf("%s/inquiry/%s", environment.Env.WEB, hash)
		http.Redirect(w, r, redirect, http.StatusFound)
	})
}

func ListenForToken() (string, error) {
	return ListenFor("token", func(w http.ResponseWriter, r *http.Request, token string) {
		var redirect string
		if token == "denied" {
			redirect = fmt.Sprintf("%s/token#cli-d", environment.Env.WEB)
		} else {
			redirect = fmt.Sprintf("%s/token#cli-s", environment.Env.WEB)
		}
		http.Redirect(w, r, redirect, http.StatusFound)
	})
}

func ListenFor(path string, action func(http.ResponseWriter, *http.Request, string)) (string, error) {
	var value string
	done := make(chan bool)

	server := &http.Server{Addr: Port}

	http.HandleFunc("/"+path+"/", func(w http.ResponseWriter, r *http.Request) {
		value = strings.TrimPrefix(r.URL.Path, "/"+path+"/")
		action(w, r, value)
		done <- true
	})

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println(err)
		}
	}()

	timer := time.AfterFunc(30*time.Second, func() {
		Shutdown(server)
	})
	select {
	case <-done:
		timer.Stop()
		Shutdown(server)
	case <-timer.C:
		Shutdown(server)
	}

	return value, nil
}

func Shutdown(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println(err)
	}
}

// UpdateInquiry update the inquiry passing ComputerName and user agent
func UpdateInquiry(hash string) error {
	var responseJSON *Response
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

// IsPortAvailable checks if a port is available by trying to listen on it
func IsPortAvailable(port string) bool {
	ln, err := net.Listen("tcp", port)

	if err != nil {
		return false
	}

	_ = ln.Close()
	return true
}
