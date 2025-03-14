# go-bump-pkg
> A tools for bumping package.json or other version files.

## installation
```sh
go get -u github.com/afeiship/go-bump-pkg
```

## methods
- `ReadPkgJson(filePath string) (pkg *bumppkg.Package, err error)` - Read package.json file and return a `Package` object.
- `BumpMajor(filePath string) error` - Bump major version.
- `BumpMinor(filePath string) error` - Bump minor version.
- `BumpPatch(filePath string) error` - Bump patch version.
- `AddPreRelease(filePath string, preRelease string) error` - Add pre-release identifier.
- `RemovePreRelease(filePath string) error` - Remove pre-release identifier.
- `GetVersion(filePath string) (version string, err error)` - Get version from package.json file.

## usage
```go
package main

import (
    "fmt"
    "log"
    "github.com/afeiship/go-bump-pkg"
)

func main() {
    // Read current version
    pkg, err := bumppkg.ReadPkgJson("package.json")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Current version: %s\n", pkg.Version)

    // Bump major version (1.2.3 -> 2.0.0)
    if err := bumppkg.BumpMajor("package.json"); err != nil {
        log.Fatal(err)
    }

    // Bump minor version (1.2.3 -> 1.3.0)
    if err := bumppkg.BumpMinor("package.json"); err != nil {
        log.Fatal(err)
    }

    // Bump patch version (1.2.3 -> 1.2.4)
    if err := bumppkg.BumpPatch("package.json"); err != nil {
        log.Fatal(err)
    }

    // Add pre-release identifier (1.2.3 -> 1.2.3-beta)
    if err := bumppkg.AddPreRelease("package.json", "beta"); err != nil {
        log.Fatal(err)
    }

    // Remove pre-release identifier (1.2.3-beta -> 1.2.3)
    if err := bumppkg.RemovePreRelease("package.json"); err != nil {
        log.Fatal(err)
    }
	
	// get version
	version, err := bumppkg.GetVersion("package.json")
	if err != nil { log.Fatal(err) }
	fmt.Printf("Version: %s\n", version)
}
```