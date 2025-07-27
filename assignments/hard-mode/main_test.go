package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMainWithExistingFile(t *testing.T) {
	// Create temp file
	content := "test content"
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	// Run program
	cmd := exec.Command("go", "run", "main.go", tmpfile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	if string(output) != content {
		t.Errorf("Expected output %q, got %q", content, string(output))
	}
}

func TestMainWithNonExistentFile(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "nonexistent.txt")
	output, _ := cmd.CombinedOutput()

	if !strings.Contains(string(output), "Error opening file") {
		t.Errorf("Expected error message about opening file, got %q", string(output))
	}
}

func TestMainWithNoArgs(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	output, _ := cmd.CombinedOutput()

	if !strings.Contains(string(output), "index out of range") {
		t.Errorf("Expected index out of range error, got %q", string(output))
	}
}