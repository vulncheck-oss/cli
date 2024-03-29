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
	root.SetArgs([]string{"auth"})
	if err := root.Execute(); err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if strings.HasPrefix(i18n.C.AuthShort, actual.String()) {
		t.Errorf("expected %s but got %s", i18n.C.AuthShort, actual.String())
	}
}

func Test_IndicesCommand(t *testing.T) {
	actual, root := setRootActual("indices")
	root.SetArgs([]string{"indices"})
	if err := root.Execute(); err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if strings.HasPrefix(i18n.C.IndicesShort, actual.String()) {
		t.Errorf("expected %s but got %s", i18n.C.IndicesShort, actual.String())
	}
}

func Test_IndexCommand(t *testing.T) {
	actual, root := setRootActual("index")
	root.SetArgs([]string{"index"})
	if err := root.Execute(); err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if strings.HasPrefix(i18n.C.IndexShort, actual.String()) {
		t.Errorf("expected %s but got %s", i18n.C.IndexShort, actual.String())
	}
}

func Test_BackupCommand(t *testing.T) {
	actual, root := setRootActual("backup")
	root.SetArgs([]string{"backup"})
	if err := root.Execute(); err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if strings.HasPrefix(i18n.C.BackupShort, actual.String()) {
		t.Errorf("expected %s but got %s", i18n.C.BackupShort, actual.String())
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
