package offline

import (
	"bytes"
	"testing"
)

func TestCommand(t *testing.T) {
	// Get the offline command
	offlineCmd := Command()

	// Test that the command has the correct properties
	if offlineCmd.Use != "offline <command>" {
		t.Errorf("Expected Use: 'offline <command>', got: '%s'", offlineCmd.Use)
	}

	if offlineCmd.Short != "Offline commands" {
		t.Errorf("Expected Short: 'Offline commands', got: '%s'", offlineCmd.Short)
	}

	// Create a map of actual subcommands for easy lookup
	actualSubcommands := make(map[string]bool)
	for _, cmd := range offlineCmd.Commands() {
		actualSubcommands[cmd.Name()] = true
	}

	// Verify mandatory base subcommands are present
	mandatoryCommands := []string{
		"sync",
		"ipintel",
		"purl",
		"cpe",
		"status",
	}

	for _, cmdName := range mandatoryCommands {
		if !actualSubcommands[cmdName] {
			t.Errorf("Expected mandatory subcommand '%s' not found", cmdName)
		}
	}

	// Count how many ipintel alias commands we have
	// We expect 3 from the ipintel.AliasCommands() based on timeframes (3d, 10d, 30d)
	ipintelAliasCount := 0
	for cmdName := range actualSubcommands {
		if cmdName != "ipintel" && len(cmdName) > 7 && cmdName[:8] == "ipintel-" {
			ipintelAliasCount++
		}
	}

	if ipintelAliasCount != 3 {
		t.Errorf("Expected 3 ipintel alias commands, found %d", ipintelAliasCount)
	}

	// Test the Run function (should execute Help)
	buf := new(bytes.Buffer)
	offlineCmd.SetOut(buf)
	offlineCmd.SetErr(buf)

	// Execute the command without subcommands or arguments
	// This should trigger the Run function which calls cmd.Help()
	offlineCmd.SetArgs([]string{})
	err := offlineCmd.Execute()
	if err != nil {
		t.Errorf("Expected no error when executing command, got: %v", err)
	}

	// Check that the help output was generated
	if buf.Len() == 0 {
		t.Error("Expected help output, but buffer is empty")
	}
}

// TestSubcommandExecution tests that each subcommand can be properly accessed
func TestSubcommandExecution(t *testing.T) {
	offlineCmd := Command()

	// Get all subcommand names dynamically
	var subcommandNames []string
	for _, cmd := range offlineCmd.Commands() {
		subcommandNames = append(subcommandNames, cmd.Name())
	}

	// Test accessing each subcommand
	for _, subcmd := range subcommandNames {
		// Set up a command execution that just accesses the subcommand but doesn't execute it
		offlineCmd.SetArgs([]string{subcmd, "--help"})

		// We don't need to check the output, just that no panic occurs
		// and the subcommand is found
		err := offlineCmd.Execute()
		if err != nil {
			t.Errorf("Error accessing subcommand '%s': %v", subcmd, err)
		}
	}
}
