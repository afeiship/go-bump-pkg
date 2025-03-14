package bumppkg

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PkgJson represents the structure of a package.json file
type PkgJson struct {
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Description string            `json:"description"`
	Private     bool              `json:"private"`
	License     string            `json:"license"`
	Scripts     map[string]string `json:"scripts"`
}

// BumpMajor increments the major version number and resets minor and patch to 0
func BumpMajor(filename string) error {
	pkg, err := ReadPkgJson(filename)
	if err != nil {
		return err
	}

	parts := strings.Split(pkg.Version, ".")
	if len(parts) != 3 {
		return fmt.Errorf("invalid version format: %s", pkg.Version)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid major version: %s", parts[0])
	}

	pkg.Version = fmt.Sprintf("%d.0.0", major+1)
	return WritePkgJson(filename, pkg)
}

// BumpMinor increments the minor version number and resets patch to 0
func BumpMinor(filename string) error {
	pkg, err := ReadPkgJson(filename)
	if err != nil {
		return err
	}

	parts := strings.Split(pkg.Version, ".")
	if len(parts) != 3 {
		return fmt.Errorf("invalid version format: %s", pkg.Version)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid major version: %s", parts[0])
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid minor version: %s", parts[1])
	}

	pkg.Version = fmt.Sprintf("%d.%d.0", major, minor+1)
	return WritePkgJson(filename, pkg)
}

// BumpPatch increments the patch version number
func BumpPatch(filename string) error {
	pkg, err := ReadPkgJson(filename)
	if err != nil {
		return err
	}

	parts := strings.Split(pkg.Version, ".")
	if len(parts) != 3 {
		return fmt.Errorf("invalid version format: %s", pkg.Version)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid major version: %s", parts[0])
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid minor version: %s", parts[1])
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return fmt.Errorf("invalid patch version: %s", parts[2])
	}

	pkg.Version = fmt.Sprintf("%d.%d.%d", major, minor, patch+1)
	return WritePkgJson(filename, pkg)
}

// ReadPkgJson reads and parses a package.json file
func ReadPkgJson(filename string) (*PkgJson, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read package.json: %v", err)
	}

	var pkg PkgJson
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, fmt.Errorf("failed to parse package.json: %v", err)
	}

	return &pkg, nil
}

// WritePkgJson writes the package.json content back to file
func WritePkgJson(filename string, pkg *PkgJson) error {
	data, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal package.json: %v", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write package.json: %v", err)
	}

	return nil
}
