package cachedir

import "github.com/jlentink/aem/internal/cli/project"

func migrate(){
	initiated = true
	if project.Exists(getCacheRoot()) && !project.IsFile(getCacheRoot()){
		return
	}
	if project.Exists(getCacheRoot()) && project.IsFile(getCacheRoot()){
		project.Rename(getCacheRoot(), getCacheRoot() +".bak")
		createCacheDir()
		project.Rename(getCacheRoot() +".bak",getCacheRoot() + "/projects.toml")
	}
}
