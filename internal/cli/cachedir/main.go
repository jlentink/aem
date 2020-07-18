package cachedir

import (
	"fmt"
	"github.com/jlentink/aem/internal/cli/project"
)

var doCache = true
const cacheDir = ".aem"

func Init() {
	migrate()
}

func Disable(){
	doCache = false
}

func createCacheDir() {
	if !project.Exists(getCacheRoot()) {
		_, err := project.CreateDir(getCacheRoot())
		if err != nil {
			processErr = err
		}
	}
}

func getCacheRoot() string {
	homedir, err := project.HomeDir()
	if err != nil {
		processErr = err
		return ``
	}
	return fmt.Sprintf("%s/%s", homedir, cacheDir)
}
