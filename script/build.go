// Build tasks for the VulnCheck CLI project.
//
// Usage:  go run script/build.go [<tasks>...] [<env>...]
//
// Known tasks are:
//
//   bin/gh:
//     Builds the main executable.
//     Supported environment variables:
//     - VC_VERSION: determined from source by default
//     - VC_OAUTH_CLIENT_ID
//     - VC_OAUTH_CLIENT_SECRET
//     - SOURCE_DATE_EPOCH: enables reproducible builds
//     - GO_LDFLAGS
//
//   manpages:
//     Builds the man pages under `share/man/man1/`.
//
//   clean:
//     Deletes all built files.
//

package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var tasks = map[string]func(string) error{
	"bin/vc": func(exe string) error {
		ldflags := os.Getenv("GO_LDFLAGS")
		ldflags = fmt.Sprintf("-X github.com/vulncheck-oss/cli/pkg/build.Version=%s %s", version(), ldflags)
		ldflags = fmt.Sprintf("-X github.com/vulncheck-oss/cli/pkg/build.Date=%s %s", date(), ldflags)

		return run("go", "build", "-trimpath", "-ldflags", ldflags, "-o", exe, "./cmd/vc")
	},
	"manpages": func(_ string) error {
		return run("go", "run", "./cmd/gen-docs", "--man-page", "--doc-path", "./share/man/man1/")
	},
	"clean": func(_ string) error {
		return rmrf("bin", "share")
	},
}

var self string

func main() {
	args := os.Args[:1]
	for _, arg := range os.Args[1:] {
		if idx := strings.IndexRune(arg, '='); idx >= 0 {
			os.Setenv(arg[:idx], arg[idx+1:])
		} else {
			args = append(args, arg)
		}
	}

	if len(args) < 2 {
		if isWindowsTarget() {
			args = append(args, filepath.Join("bin", "vc.exe"))
		} else {
			args = append(args, "bin/vc")
		}
	}

	self = filepath.Base(args[0])
	if self == "build" {
		self = "build.go"
	}

	for _, task := range args[1:] {
		t := tasks[normalizeTask(task)]
		if t == nil {
			fmt.Fprintf(os.Stderr, "Don't know how to build task `%s`.\n", task)
			os.Exit(1)
		}

		err := t(task)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprintf(os.Stderr, "%s: building task `%s` failed.\n", self, task)
			os.Exit(1)
		}
	}
}

func version() string {
	if versionEnv := os.Getenv("VC_VERSION"); versionEnv != "" {
		return versionEnv
	}
	if desc, err := cmdOutput("git", "describe", "--tags"); err == nil {
		return desc
	}
	rev, _ := cmdOutput("git", "rev-parse", "--short", "HEAD")
	return rev
}

func date() string {
	t := time.Now()
	if sourceDate := os.Getenv("SOURCE_DATE_EPOCH"); sourceDate != "" {
		if sec, err := strconv.ParseInt(sourceDate, 10, 64); err == nil {
			t = time.Unix(sec, 0)
		}
	}
	return t.Format("2006-01-02")
}

func cmdOutput(args ...string) (string, error) {
	exe, err := exec.LookPath(args[0])
	if err != nil {
		return "", err
	}
	cmd := exec.Command(exe, args[1:]...)
	cmd.Stderr = io.Discard
	out, err := cmd.Output()
	return strings.TrimSuffix(string(out), "\n"), err
}

func run(args ...string) error {
	exe, err := exec.LookPath(args[0])
	if err != nil {
		return err
	}
	announce(args...)
	cmd := exec.Command(exe, args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func rmrf(targets ...string) error {
	args := append([]string{"rm", "-rf"}, targets...)
	announce(args...)
	for _, target := range targets {
		if err := os.RemoveAll(target); err != nil {
			return err
		}
	}
	return nil
}

func announce(args ...string) {
	fmt.Println(shellInspect(args))
}

func shellInspect(args []string) string {
	fmtArgs := make([]string, len(args))
	for i, arg := range args {
		if strings.ContainsAny(arg, " \t'\"") {
			fmtArgs[i] = fmt.Sprintf("%q", arg)
		} else {
			fmtArgs[i] = arg
		}
	}
	return strings.Join(fmtArgs, " ")
}

func normalizeTask(t string) string {
	return filepath.ToSlash(strings.TrimSuffix(t, ".exe"))
}

func isWindowsTarget() bool {
	if os.Getenv("GOOS") == "windows" {
		return true
	}
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}
