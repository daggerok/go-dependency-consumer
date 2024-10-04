package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	// Run the tests
	exitCode := m.Run()
	os.Exit(exitCode) // Exit with the exit code of the tests
}

func TestGreetMK(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "MK")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run main.go: %v", err)
	}
	expected := "Hello, MK!\n" // '\n' here is important!
	if string(out) != expected {
		t.Fatalf("Expected: %v, got: %v", expected, string(out))
	}
}

func TestGreetAnonymous(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run main.go: %v", err)
	}
	expected := "Hello, anonymous!\n" // '\n' here is important!
	if string(out) != expected {
		t.Fatalf("Expected: %v, got: %v", expected, string(out))
	}
}
