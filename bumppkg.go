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

// parseVersion splits a version string into its components
func parseVersion(version string) (major, minor, patch int, preRelease string, err error) {
	// Split version into main part and pre-release part
	parts := strings.SplitN(version, "-", 2)
	mainVersion := parts[0]
	if len(parts) > 1 {
		preRelease = parts[1]
	}

	// Split main version into major.minor.patch
	versionParts := strings.Split(mainVersion, ".")
	if len(versionParts) != 3 {
		return 0, 0, 0, "", fmt.Errorf("invalid version format: %s", version)
	}

	major, err = strconv.Atoi(versionParts[0])
	if err != nil {
		return 0, 0, 0, "", fmt.Errorf("invalid major version: %s", versionParts[0])
	}

	minor, err = strconv.Atoi(versionParts[1])
	if err != nil {
		return 0, 0, 0, "", fmt.Errorf("invalid minor version: %s", versionParts[1])
	}

	patch, err = strconv.Atoi(versionParts[2])
	if err != nil {
		return 0, 0, 0, "", fmt.Errorf("invalid patch version: %s", versionParts[2])
	}

	return major, minor, patch, preRelease, nil
}

// formatVersion formats version components into a version string
func formatVersion(major, minor, patch int, preRelease string) string {
	version := fmt.Sprintf("%d.%d.%d", major, minor, patch)
	if preRelease != "" {
		version = fmt.Sprintf("%s-%s", version, preRelease)
	}
	return version
}

// BumpMajor increments the major version number and resets minor and patch to 0
func BumpMajor(filename string) error {
	pkg, err := ReadPkgJson(filename)
	if err != nil {
		return err
	}

	major, _, _, _, err := parseVersion(pkg.Version)
	if err != nil {
		return err
	}

	pkg.Version = formatVersion(major+1, 0, 0, "")
	return WritePkgJson(filename, pkg)
}

// BumpMinor increments the minor version number and resets patch to 0
func BumpMinor(filename string) error {
	pkg, err := ReadPkgJson(filename)
	if err != nil {
		return err
	}

	major, minor, _, _, err := parseVersion(pkg.Version)
	if err != nil {
		return err
	}

	pkg.Version = formatVersion(major, minor+1, 0, "")
	return WritePkgJson(filename, pkg)
}

// BumpPatch increments the patch version number
func BumpPatch(filename string) error {
	pkg, err := ReadPkgJson(filename)
	if err != nil {
		return err
	}

	major, minor, patch, _, err := parseVersion(pkg.Version)
	if err != nil {
		return err
	}

	pkg.Version = formatVersion(major, minor, patch+1, "")
	return WritePkgJson(filename, pkg)
}

// AddPreRelease adds or updates the pre-release identifier
func AddPreRelease(filename string, identifier string) error {
	pkg, err := ReadPkgJson(filename)
	if err != nil {
		return err
	}

	major, minor, patch, _, err := parseVersion(pkg.Version)
	if err != nil {
		return err
	}

	pkg.Version = formatVersion(major, minor, patch, identifier)
	return WritePkgJson(filename, pkg)
}

// RemovePreRelease removes the pre-release identifier
func RemovePreRelease(filename string) error {
	pkg, err := ReadPkgJson(filename)
	if err != nil {
		return err
	}

	major, minor, patch, _, err := parseVersion(pkg.Version)
	if err != nil {
		return err
	}

	pkg.Version = formatVersion(major, minor, patch, "")
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
