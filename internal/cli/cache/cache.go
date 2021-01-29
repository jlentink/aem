package cache

import (
	"fmt"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/spf13/afero"
)

const cacheDir = ".aem"
const filesDir = "files"
const filesCacheFile = "files.toml"

// Cache system
type Cache struct {
	fs afero.Fs
	root string
}

func (c *Cache) getCacheRoot() (string, error) {
	homedir, err := project.HomeDir()
	if err != nil {
		return ``, err
	}
	c.root = fmt.Sprintf("%s/%s", homedir, cacheDir)
	return c.root, nil
}

func (c *Cache) createFs() {
	if c.root == "" {
		c.getCacheRoot()
	}
	c.fs = afero.NewBasePathFs(afero.NewOsFs(), c.root)
}

func (c *Cache) createCacheDir() error {
	root, err := c.getCacheRoot()
	if err != nil {
		return err
	}

	exists, err := afero.DirExists(c.fs, root + "/" + filesDir)
	if err != nil {
		return err
	}

	if !exists {
		err := c.fs.MkdirAll(root + "/" + filesDir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}