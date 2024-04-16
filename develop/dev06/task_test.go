package main

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {
	input := "1\t2\t3\n4\t5\t6\n7\t8\t9\n"

	tests := []struct {
		args        []string
		expectedOut string
	}{
		{[]string{"-f", "1", "-d", "\t"}, "1\n4\n7\n"},
		{[]string{"-f", "2", "-d", "\t"}, "2\n5\n8\n"},
		{[]string{"-f", "1,3", "-d", "\t"}, "1\t3\n4\t6\n7\t9\n"},
		{[]string{"-s", "-d", "\t"}, "1\t2\t3\n4\t5\t6\n7\t8\t9\n"},
	}

	for _, test := range tests {
		t.Run(strings.Join(test.args, " "), func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = append([]string{"cmd"}, test.args...)

			tempFile, err := ioutil.TempFile("", "test_stdout.txt")
			if err != nil {
				t.Fatal("Error creating temp file:", err)
			}
			defer os.Remove(tempFile.Name())
			oldStdout := os.Stdout
			os.Stdout = tempFile

			oldStdin := os.Stdin
			r, w, _ := os.Pipe()
			os.Stdin = r

			io.WriteString(w, input)
			w.Close()
			main()

			os.Stdin = oldStdin
			os.Stdout = oldStdout

			tempFile.Close()
			content, err := ioutil.ReadFile(tempFile.Name())
			if err != nil {
				t.Fatal("Error reading temp file:", err)
			}
			if string(content) != test.expectedOut {
				t.Errorf("Expected output: %q, but got: %q", test.expectedOut, string(content))
			}
		})
	}
}
