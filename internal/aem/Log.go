package aem

import (
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/cli/project"
	"os"
)

// ListLogFiles log files for instance
func ListLogFiles(i *objects.Instance) ([]os.FileInfo, error) {
	logPath, _ := project.GetLogDirLocation(*i)
	files, err := project.ReadDir(logPath)
	if err != nil {
		return []os.FileInfo{}, err
	}
	return files, err
}
