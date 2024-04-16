package main

import (
	"io"
	"net"
	"os/exec"
	"strings"
	"syscall"
	"testing"
)

func TestTelnetClient(t *testing.T) {
	// Start a temporary TCP server
	listener, err := net.Listen("tcp", "127.0.0.1:0") // Use a random port
	if err != nil {
		t.Fatal("Error starting TCP server:", err)
	}
	defer listener.Close()

	// Accept a connection in a goroutine
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		// Echo back whatever is received
		io.Copy(conn, conn)
	}()

	// Start the telnet client as a subprocess
	port := listener.Addr().(*net.TCPAddr).Port

	cmd := exec.Command("go", "run", "main.go", "--timeout=5s", "127.0.0.1", string(rune(port)))
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal("Error creating stdin pipe:", err)
	}
	defer stdin.Close()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal("Error creating stdout pipe:", err)
	}
	defer stdout.Close()

	err = cmd.Start()
	if err != nil {
		t.Fatal("Error starting telnet client:", err)
	}
	defer func() {
		cmd.Process.Signal(syscall.SIGINT)
		cmd.Wait()
	}()

	// Send input to the telnet client
	input := "test\n"
	io.WriteString(stdin, input)

	// Read output from the telnet client
	output := make([]byte, 1024)
	n, err := stdout.Read(output)
	if err != nil && err != io.EOF {
		t.Fatal("Error reading output from telnet client:", err)
	}

	// Check if the output matches the input
	if strings.TrimSpace(string(output[:n])) != input {
		t.Errorf("Expected output %q, got %q", input, string(output[:n]))
	}
}

func TestMainFunction(t *testing.T) {
	// Test the main function by running it with invalid arguments
	cmd := exec.Command("go", "run", "main.go")
	err := cmd.Run()
	if err == nil {
		t.Error("Expected error for missing arguments, got nil")
	}
}
