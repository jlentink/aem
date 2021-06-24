package version

import "fmt"

var (
	version = "snapshot"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

// GetVersion Get version for application
func GetVersion() string {
	return version
}

// GetBuild Get build hash for application
func GetBuild() string {
	return commit
}

// DisplayVersion returns version string for application
func DisplayVersion(v, m bool) string {
	if m {
		return fmt.Sprintf("%s\n", version)
	} else if v {
		return fmt.Sprintf("AEMcli (https://github.com/jlentink/aem)\n" +
							"Version: %s\n" +
							"Built: %s\n" +
							"Date: %s\n", version, commit, date)
	}
	return fmt.Sprintf("AEMcli\n" +
		"Version: %s\n", version)
}
