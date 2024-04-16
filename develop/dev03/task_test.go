package main

import (
	"io/ioutil"
	"os"
	"sort"
	"testing"
)

func TestMainFunction(t *testing.T) {
	inputFile := createTempFile(t, "example.txt", "2 b\n3 c\n1 a\n3 c\n2 b\n")
	defer os.Remove(inputFile.Name())

	outputFile := createTempFile(t, "res.txt", "")
	os.Args = []string{"cmd", "-input", inputFile.Name(), "-output", outputFile.Name(), "-k", "1", "-n", "-r", "-u"}
	main()

	expectedOutput := "3 c\n2 b\n1 a\n"
	actualOutput := readTempFile(t, outputFile)
	if actualOutput != expectedOutput {
		t.Errorf("Output file content is incorrect. Expected: %s, got: %s", expectedOutput, actualOutput)
	}
}

func createTempFile(t *testing.T, fileName, content string) *os.File {
	file, err := ioutil.TempFile("", fileName)
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer file.Close()

	if _, err := file.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	return file
}

func readTempFile(t *testing.T, file *os.File) string {
	content, err := ioutil.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("Failed to read temporary file: %v", err)
	}
	return string(content)
}

func TestSorting(t *testing.T) {
	lines := []string{"2 b", "3 c", "1 a", "3 c", "2 b"}
	expected := []string{"1 a", "2 b", "2 b", "3 c", "3 c"}

	sortingFunction := func() {
		sort.Strings(lines)
	}

	assertSorted(t, sortingFunction, lines, expected)
}

func assertSorted(t *testing.T, sortingFunction func(), lines []string, expected []string) {
	sortingFunction()

	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("Expected line %q at index %d, but got %q", expected[i], i, line)
		}
	}
}
