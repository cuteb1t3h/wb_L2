package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestGrep(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Записываем тестовые данные во временный файл
	data := []byte("Hello, World!\nThis is a test file.\nAnother line with a pattern.\nOne more line.\n")
	if _, err := tempFile.Write(data); err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		args        []string
		expectedOut string
		expectedErr string
	}{
		{[]string{"-n", "pattern", tempFile.Name()}, "3:Another line with a pattern.\n", ""},
		{[]string{"-c", "pattern", tempFile.Name()}, "Match count: 1\n", ""},
		{[]string{"-B", "1", "pattern", tempFile.Name()}, "This is a test file.\nAnother line with a pattern.\n", ""},
	}

	for _, test := range tests {
		t.Run(strings.Join(test.args, " "), func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = test.args

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			main()

			w.Close()
			os.Stdout = oldStdout

			// Сравниваем вывод с ожидаемым результатом
			out, _ := ioutil.ReadAll(r)
			if string(out) != test.expectedOut {
				t.Errorf("Expected output: %q, but got: %q", test.expectedOut, string(out))
			}
		})
	}
}
