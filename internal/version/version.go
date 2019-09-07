package version

import "fmt"

var (
	//Main app version
	Main = "1.0.0"
	// Build app build
	Build = "no-build-hash"
)

// GetVersion Get version for application
func GetVersion() string {
	return Main
}

// GetBuild Get build hash for application
func GetBuild() string {
	return Build
}

// DisplayVersion returns version string for application
func DisplayVersion(v, m bool) string {
	if m {
		return fmt.Sprintf("%s\n", GetVersion())
	} else if v {
		return fmt.Sprintf("AEMcli (https://github.com/jlentink/aem)\nVersion: %s\nBuilt: %s\n", GetVersion(), GetBuild())
	}
	return fmt.Sprintf("AEMcli\nVersion: %s\n", GetVersion())
}
