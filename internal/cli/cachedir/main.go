package cachedir

import (
	"fmt"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/output"
)

var doCache = true
const cacheDir = ".aem"
const filesDir = "files"
var initiated = false

// Init Cache
func Init() {
	migrate()
	createCacheDir()
}

// Disable cache
func Disable(){
	doCache = false
}

func createCacheDir() {
	if !project.Exists(getCacheRoot()) {
		_, err := project.CreateDir(getCacheRoot())
		if err != nil {
			output.Printf(output.VERBOSE, "Could not create cacheDir %s", err.Error())
		}
	}
	if !project.Exists(getCacheRoot() + "/" + filesDir) {
		_, err := project.CreateDir(getCacheRoot() + "/" + filesDir)
		if err != nil {
			output.Printf(output.VERBOSE, "Could not create cacheDir %s", err.Error())
		}
	}
}

func getCacheRoot() string {
	homedir, err := project.HomeDir()
	if err != nil {
		output.Printf(output.VERBOSE, "Could not find homedir %s", err.Error())
		return ``
	}
	return fmt.Sprintf("%s/%s", homedir, cacheDir)
}
