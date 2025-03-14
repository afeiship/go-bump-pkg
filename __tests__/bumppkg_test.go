package bumppkg__test

import (
	"encoding/json"
	bumppkg "github.com/afeiship/go-bump-pkg"
	"os"
	"testing"
)

const testFile = "./test_package.json"

func setupTestFile(t *testing.T) {
	pkg := bumppkg.PkgJson{
		Name:        "test-package",
		Version:     "1.2.3",
		Description: "Test package",
		Private:     true,
		License:     "MIT",
		Scripts:     map[string]string{"test": "go test"},
	}

	data, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal test package.json: %v", err)
	}

	if err := os.WriteFile(testFile, data, 0644); err != nil {
		t.Fatalf("Failed to write test package.json: %v", err)
	}
}

func cleanupTestFile(t *testing.T) {
	if err := os.Remove(testFile); err != nil && !os.IsNotExist(err) {
		t.Fatalf("Failed to cleanup test package.json: %v", err)
	}
}

func TestReadPkgJson(t *testing.T) {
	setupTestFile(t)
	defer cleanupTestFile(t)

	pkg, err := bumppkg.ReadPkgJson(testFile)
	if err != nil {
		t.Fatalf("Failed to read package.json: %v", err)
	}

	if pkg.Name != "test-package" {
		t.Errorf("Expected name 'test-package', got '%s'", pkg.Name)
	}
	if pkg.Version != "1.2.3" {
		t.Errorf("Expected version '1.2.3', got '%s'", pkg.Version)
	}
}

func TestBumpMajor(t *testing.T) {
	setupTestFile(t)
	defer cleanupTestFile(t)

	if err := bumppkg.BumpMajor(testFile); err != nil {
		t.Fatalf("Failed to bump major version: %v", err)
	}

	pkg, err := bumppkg.ReadPkgJson(testFile)
	if err != nil {
		t.Fatalf("Failed to read package.json after bump: %v", err)
	}

	if pkg.Version != "2.0.0" {
		t.Errorf("Expected version '2.0.0', got '%s'", pkg.Version)
	}
}

func TestBumpMinor(t *testing.T) {
	setupTestFile(t)
	defer cleanupTestFile(t)

	if err := bumppkg.BumpMinor(testFile); err != nil {
		t.Fatalf("Failed to bump minor version: %v", err)
	}

	pkg, err := bumppkg.ReadPkgJson(testFile)
	if err != nil {
		t.Fatalf("Failed to read package.json after bump: %v", err)
	}

	if pkg.Version != "1.3.0" {
		t.Errorf("Expected version '1.3.0', got '%s'", pkg.Version)
	}
}

func TestBumpPatch(t *testing.T) {
	setupTestFile(t)
	defer cleanupTestFile(t)

	if err := bumppkg.BumpPatch(testFile); err != nil {
		t.Fatalf("Failed to bump patch version: %v", err)
	}

	pkg, err := bumppkg.ReadPkgJson(testFile)
	if err != nil {
		t.Fatalf("Failed to read package.json after bump: %v", err)
	}

	if pkg.Version != "1.2.4" {
		t.Errorf("Expected version '1.2.4', got '%s'", pkg.Version)
	}
}
