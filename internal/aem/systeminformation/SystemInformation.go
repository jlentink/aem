package systeminformation

import (
	"encoding/json"
	"github.com/jlentink/aem/internal/aem"
	"github.com/jlentink/aem/internal/aem/objects"
)

const (
	// URLSystemInformation System information url
	URLSystemInformation = "/libs/granite/operations/content/systemoverview/export.json"
)

// GetSystemInformation gets system information from running instance
func GetSystemInformation(i *objects.Instance) (*objects.SystemInformation, error) {
	body, err := aem.GetFromInstance(i, URLSystemInformation)
	if err != nil {
		return nil, err
	}
	sysInfo := objects.SystemInformation{}
	err = json.Unmarshal(body, &sysInfo)
	if err != nil {
		return nil, err
	}

	if len(sysInfo.SystemInformation.Windows) > 0 {
		sysInfo.SystemInformation.CurrentOS = "Windows " + sysInfo.SystemInformation.Windows
	} else if len(sysInfo.SystemInformation.Linux) > 0 {
		sysInfo.SystemInformation.CurrentOS = "Linux " + sysInfo.SystemInformation.Linux
	} else if len(sysInfo.SystemInformation.MacOSX) > 0 {
		sysInfo.SystemInformation.CurrentOS = "MacOS " + sysInfo.SystemInformation.MacOSX
	}

	return &sysInfo, nil
}
