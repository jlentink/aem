package setup

import (
	"os/exec"
	"runtime"
	"strings"
)

// Requirement of bin
const (
	Required = int(iota)
	Optional
)

var (
	general = []Description{
		{Bin: "java", Description: "needed to run AEM and run-oak. (Crucial)", Required: Required},
		{Bin: "vlt", Description: "needed for all vlt actions", Required: Optional},
		{Bin: "aemsync", Description: "needed to sync frontend code with AEM instance", Required: Optional},
		{Bin: "lazybones", Description: "Lazybones templates for Adobe Experience Manager", Required: Optional},
		{Bin: "git", Description: "Git version control for versioning", Required: Optional},
		{Bin: "mvn", Description: "Maven build tool", Required: Optional},
	}
	darwin = []Description{
		{Bin: "kill", Description: "needed to run stop. (Crucial)", Required: Required},
		{Bin: "open", Description: "needed to open the browser with the correct URL.", Required: Optional},
	}
	linux = []Description{
		{Bin: "kill", Description: "needed to run stop. (Crucial)", Required: Required},
	}
	windows = []Description{
		{Bin: "rundll32", Description: "needed to open the browser with the correct URL.", Required: Optional},
	}
)

func osCombine() []Description {
	switch strings.ToLower(runtime.GOOS) {
	case "darwin":
		return append(general, darwin...)
	case "linux":
		return append(general, linux...)
	case "windows":
		return append(general, windows...)
	default:
		return general
	}
}

func check(bin Description) Description {
	_, err := exec.LookPath(bin.Bin)
	bin.Found = nil == err
	return bin
}

// Check bin if it is available
func Check() []Description {
	binaries := osCombine()
	for i, bin := range binaries {
		binaries[i] = check(bin)
	}
	return binaries
}
