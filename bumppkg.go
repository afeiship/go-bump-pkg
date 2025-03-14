package bumppkg

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BumpType string

// bump type enums
const (
	Major BumpType = "major"
	Minor BumpType = "minor"
	Patch BumpType = "patch"
)

// PkgJson represents the structure of a package.json file
// Example:
//
//	pkg := PkgJson{
//		Name: "my-package",
//		Version: "1.0.0",
//		Description: "My awesome package",
//		Private: false,
//		License: "MIT",
//		Scripts: map[string]string{"test": "go test"},
//	}
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

// BumpBy increments the version number by the specified bump type
//
// Example:
//
//		err := BumpBy("package.json", Major)
//		// Version will be "2.0.0"
//		err := BumpBy("package.json", Minor)
//		// Version will be "1.3.0"
//	 err := BumpBy("package.json", Patch)
func BumpBy(filename string, bumpType BumpType) (string, error) {
	switch bumpType {
	case Major:
		return BumpMajor(filename)
	case Minor:
		return BumpMinor(filename)
	case Patch:
		return BumpPatch(filename)
	default:
		return "", fmt.Errorf("invalid bump type: %s", bumpType)
	}
}

// BumpMajor increments the major version number and resets minor and patch to 0
//
// Example:
//
//	// If package.json has version "1.2.3"
//	err := BumpMajor("package.json")
//	// Version will be "2.0.0"
func BumpMajor(filename string) (string, error) {
	pkg, err := ReadPkgJson(filename)
	if err != nil {
		return "", err
	}

	major, _, _, _, err := parseVersion(pkg.Version)
	if err != nil {
		return "", err
	}

	pkg.Version = formatVersion(major+1, 0, 0, "")
	if err := WritePkgJson(filename, pkg); err != nil {
		return "", err
	}
	return pkg.Version, nil
}

// BumpMinor increments the minor version number and resets patch to 0
//
// Example:
//
//	// If package.json has version "1.2.3"
//	err := BumpMinor("package.json")
//	// Version will be "1.3.0"
func BumpMinor(filename string) (string, error) {
	pkg, err := ReadPkgJson(filename)
	if err != nil {
		return "", err
	}

	major, minor, _, _, err := parseVersion(pkg.Version)
	if err != nil {
		return "", err
	}

	pkg.Version = formatVersion(major, minor+1, 0, "")
	if err := WritePkgJson(filename, pkg); err != nil {
		return "", err
	}
	return pkg.Version, nil
}

// BumpPatch increments the patch version number
//
// Example:
//
//	// If package.json has version "1.2.3"
//	err := BumpPatch("package.json")
//	// Version will be "1.2.4"
func BumpPatch(filename string) (string, error) {
	pkg, err := ReadPkgJson(filename)
	if err != nil {
		return "", err
	}

	major, minor, patch, _, err := parseVersion(pkg.Version)
	if err != nil {
		return "", err
	}

	pkg.Version = formatVersion(major, minor, patch+1, "")
	if err := WritePkgJson(filename, pkg); err != nil {
		return "", err
	}
	return pkg.Version, nil
}

// AddPreRelease adds or updates the pre-release identifier
//
// Example:
//
//	// If package.json has version "1.2.3"
//	err := AddPreRelease("package.json", "beta")
//	// Version will be "1.2.3-beta"
func AddPreRelease(filename string, identifier string) (string, error) {
	pkg, err := ReadPkgJson(filename)
	if err != nil {
		return "", err
	}

	major, minor, patch, _, err := parseVersion(pkg.Version)
	if err != nil {
		return "", err
	}

	pkg.Version = formatVersion(major, minor, patch, identifier)
	if err := WritePkgJson(filename, pkg); err != nil {
		return "", err
	}
	return pkg.Version, nil
}

// RemovePreRelease removes the pre-release identifier
//
// Example:
//
//	// If package.json has version "1.2.3-beta"
//	err := RemovePreRelease("package.json")
//	// Version will be "1.2.3"
func RemovePreRelease(filename string) (string, error) {
	pkg, err := ReadPkgJson(filename)
	if err != nil {
		return "", err
	}

	major, minor, patch, _, err := parseVersion(pkg.Version)
	if err != nil {
		return "", err
	}

	pkg.Version = formatVersion(major, minor, patch, "")
	if err := WritePkgJson(filename, pkg); err != nil {
		return "", err
	}
	return pkg.Version, nil
}

// ReadPkgJson reads and parses a package.json file
//
// Example:
//
//	pkg, err := ReadPkgJson("package.json")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Current version: %s\n", pkg.Version)
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
//
// Example:
//
//	pkg := &PkgJson{
//		Name: "my-package",
//		Version: "1.0.0",
//	}
//	err := WritePkgJson("package.json", pkg)
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
