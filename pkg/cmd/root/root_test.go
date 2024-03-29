package root

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"strings"
	"testing"
)

func Test_AuthCommand(t *testing.T) {
	actual, root := setRootActual("auth")
	if err := root.Execute(); err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if strings.HasPrefix(i18n.C.AuthShort, actual.String()) {
		t.Errorf("expected %s but got %s", i18n.C.AuthShort, actual.String())
	}
}

func Test_AuthStatusCommand(t *testing.T) {
	actual, root := setRootActual("auth", "status", "--help")
	if err := root.Execute(); err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if strings.HasPrefix(i18n.C.AuthStatusShort, actual.String()) {
		t.Errorf("expected %s but got %s", i18n.C.AuthStatusShort, actual.String())
	}

}

func Test_AuthLoginCommand(t *testing.T) {
	actual, root := setRootActual("auth", "login", "--help")
	if err := root.Execute(); err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if strings.HasPrefix(i18n.C.AuthLoginShort, actual.String()) {
		t.Errorf("expected %s but got %s", i18n.C.AuthLoginShort, actual.String())
	}
}

func Test_AuthLogoutCommand(t *testing.T) {
	actual, root := setRootActual("auth", "logout", "--help")
	if err := root.Execute(); err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if strings.HasPrefix(i18n.C.AuthLogoutShort, actual.String()) {
		t.Errorf("expected %s but got %s", i18n.C.AuthLogoutShort, actual.String())
	}

}

func Test_IndicesCommand(t *testing.T) {
	actual, root := setRootActual("indices")
	if err := root.Execute(); err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if strings.HasPrefix(i18n.C.IndicesShort, actual.String()) {
		t.Errorf("expected %s but got %s", i18n.C.IndicesShort, actual.String())
	}
}

func Test_IndexCommand(t *testing.T) {
	actual, root := setRootActual("index")
	if err := root.Execute(); err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if strings.HasPrefix(i18n.C.IndexShort, actual.String()) {
		t.Errorf("expected %s but got %s", i18n.C.IndexShort, actual.String())
	}
}

func Test_BackupCommand(t *testing.T) {
	actual, root := setRootActual("backup")
	if err := root.Execute(); err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if strings.HasPrefix(i18n.C.BackupShort, actual.String()) {
		t.Errorf("expected %s but got %s", i18n.C.BackupShort, actual.String())
	}
}

func Test_CpeCommand(t *testing.T) {
	actual, root := setRootActual("cpe", "--help")
	if err := root.Execute(); err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if strings.HasPrefix(i18n.C.CpeShort, actual.String()) {
		t.Errorf("expected %s but got %s", i18n.C.CpeShort, actual.String())
	}
}

func Test_PurlCommand(t *testing.T) {
	actual, root := setRootActual("purl", "--help")
	if err := root.Execute(); err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if strings.HasPrefix(i18n.C.PurlShort, actual.String()) {
		t.Errorf("expected %s but got %s", i18n.C.PurlShort, actual.String())
	}
}

func setRootActual(args ...string) (*bytes.Buffer, *cobra.Command) {
	actual := new(bytes.Buffer)
	root := NewCmdRoot()
	root.SetOut(actual)
	root.SetErr(actual)
	var argsArray []string
	for _, arg := range args {
		argsArray = append(argsArray, arg)
	}
	root.SetArgs(argsArray)
	return actual, root
}
