package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDownloadPage(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test HTML content"))
	}))
	defer server.Close()

	// Create a temporary file for testing
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal("Error creating temp file:", err)
	}
	defer os.Remove(tmpfile.Name())

	// Call downloadPage with the mock server URL and the temporary file name
	err = downloadPage(server.URL, tmpfile.Name())
	if err != nil {
		t.Fatalf("Error downloading page: %v", err)
	}

	// Read the content of the temporary file
	content, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("Error reading temp file: %v", err)
	}

	// Verify that the content matches the expected HTML content
	expectedContent := "Test HTML content"
	if string(content) != expectedContent {
		t.Errorf("Expected content %q, got %q", expectedContent, string(content))
	}
}
