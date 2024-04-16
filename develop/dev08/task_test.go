package main_test

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestShellCommands(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{"pwd", "/"},
		{"cd /tmp && pwd", "/tmp"},
		{"echo Hello, World!", "Hello, World!\n"},
		{"ls", ""},
		{"echo 123 && echo 456", "123\n456\n"},
		{"ls /nonexistent", "ls: cannot access '/nonexistent': No such file or directory\n"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			cmd := exec.Command("go", "run", ".")
			cmd.Stdin = strings.NewReader(test.input + "\n")
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &out

			err := cmd.Run()
			if err != nil {
				t.Fatalf("Failed to execute shell program: %v", err)
			}

			if got := out.String(); got != test.expectedOutput {
				t.Errorf("Expected output:\n%q\nGot:\n%q", test.expectedOutput, got)
			}
		})
	}
}
