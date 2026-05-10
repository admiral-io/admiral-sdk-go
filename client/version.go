package client

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
)

const modulePath = "go.admiral.io/sdk"

// SchemaVersion is the version of the admiral-protos schema this SDK was
// generated from. Servers report this via the Capabilities endpoint and
// clients use it to enforce skew policy. Format follows `git describe`:
// "v1.24.0" on a tag, "v1.23.0-4-gabc1234" between tags, "abc1234" if no
// tags exist on admiral-protos.
const SchemaVersion = "v0.1.0-3-gda1e5e3"

// Version returns the module version (e.g., "v1.2.3" or "dev").
// When used as a dependency, returns the version from go.sum.
// When running in development, returns "dev".
func Version() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		// Check if we're a dependency with a real version
		for _, dep := range info.Deps {
			if dep.Path == modulePath && isRealVersion(dep.Version) {
				return dep.Version
			}
		}
		// Check if we're the main module with a real version
		if info.Main.Path == modulePath && isRealVersion(info.Main.Version) {
			return info.Main.Version
		}
	}
	return "dev"
}

// isRealVersion returns true if the version is a real release version,
// not a development or pseudo-version.
func isRealVersion(v string) bool {
	if v == "" || v == "(devel)" {
		return false
	}
	// Filter out pseudo-versions (v0.0.0-timestamp-commit)
	if strings.HasPrefix(v, "v0.0.0-") {
		return false
	}
	return true
}

// ClientUserAgent returns a User-Agent string for gRPC/HTTP requests.
func ClientUserAgent() string {
	return fmt.Sprintf("%s/%s (%s/%s; %s)",
		modulePath, Version(), runtime.GOOS, runtime.GOARCH, runtime.Version())
}
