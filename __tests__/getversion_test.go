package bumppkg__test

import (
	"os"
	"testing"

	"github.com/afeiship/go-bump-pkg"
)

func TestGetVersion(t *testing.T) {
	// Create a temporary test file
	testFile := "test_package.json"
	testContent := `{
		"name": "test-package",
		"version": "1.2.3"
	}`

	// Write test content to file
	err := os.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile)

	// Test successful version retrieval
	version, err := bumppkg.GetVersion(testFile)
	if err != nil {
		t.Fatalf("Failed to get version: %v", err)
	}

	if version != "1.2.3" {
		t.Errorf("Expected version '1.2.3', got '%s'", version)
	}

	// Test error handling for non-existent file
	_, err = bumppkg.GetVersion("nonexistent.json")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}

	// Test error handling for invalid JSON
	invalidFile := "invalid.json"
	err = os.WriteFile(invalidFile, []byte("invalid json"), 0644)
	if err != nil {
		t.Fatalf("Failed to create invalid test file: %v", err)
	}
	defer os.Remove(invalidFile)

	_, err = bumppkg.GetVersion(invalidFile)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}
