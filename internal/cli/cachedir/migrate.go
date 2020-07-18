package cachedir

import "github.com/jlentink/aem/internal/cli/project"

func migrate(){
	if project.Exists(getCacheRoot()) && !project.IsFile(getCacheRoot()){
		project.Rename(getCacheRoot(), getCacheRoot() +".bak")
	}
	createCacheDir()
	project.Rename(getCacheRoot() +".bak",getCacheRoot() + "/projects.toml")
}
