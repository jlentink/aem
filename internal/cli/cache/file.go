package cache

import (
	"github.com/BurntSushi/toml"
	"github.com/jlentink/aem/internal/cli/project"
)

func (c *Cache) IsCached(URI string) bool {
	for _, f := range c.getCachedFiles() {
		for _, u := range f.URI {
			if u == URI {
				return true
			}
		}
	}
	return false
}

func (c *Cache) getCachedFiles() []CachedFile {
	cf := cachedFiles{}
	if project.Exists(c.root + "/" + filesCacheFile) {
		toml.DecodeFile(c.root + "/" +filesCacheFile, &cf) // nolint: errcheck
	}
	return cf.File
}

//
//import (
//	"bytes"
//	"crypto/md5"
//	"fmt"
//	"github.com/BurntSushi/toml"
//	"github.com/jlentink/aem/internal/cli/cachedir"
//	"github.com/jlentink/aem/internal/cli/project"
//	"github.com/jlentink/aem/internal/output"
//	"github.com/spf13/afero"
//	"github.com/thoas/go-funk"
//	"io"
//	"io/ioutil"
//	"os"
//	"path/filepath"
//	"time"
//)
//
//const filesCacheFile = "files.toml"
//
//var fs afero.Fs
//
//type cachedFiles struct {
//	File []cachedFile
//}
//
//type cachedFile struct {
//	MD5 string
//	URI []string
//	Date time.Time
//	OriginalName string
//}
//
//func SetFS() {
//	fs = afero.NewOsFs()
//}
//
//func (c *Cache) IsCached(URI string) bool {
//	for _, f := range c.getCachedFiles() {
//		for _, u := range f.URI {
//			if u == URI {
//				return true
//			}
//		}
//	}
//	return false
//}
//
//func (c *Cache) AddCacheFile(URI, hash, location string) []cachedFile {
//	files := c.getCachedFiles()
//
//	files = append(files, cachedFile{
//		MD5: hash,
//		Date: time.Now(),
//		URI: []string{URI},
//		OriginalName: filepath.Base(location),
//	})
//	writeCachedfilesConfig(files)
//	return files
//}
//
//
//func CacheFile(URI, path string, forceRefresh bool) (bool, error){
//	if IsCached(URI) && forceRefresh {
//		removeUri(URI)
//	}
//
//	if IsCached(URI) {
//		return true, fmt.Errorf("file is already cached for URL")
//	}
//
//	if !project.IsFile(path) {
//		return false, fmt.Errorf("to be cached file is not find")
//	}
//
//	md5Hash, err := md5Sum(path)
//	if err != nil {
//		return false, fmt.Errorf("could not create a MD5Sum")
//	}
//
//	if isCachedMD5(md5Hash){
//		addUriToHash(URI, md5Hash)
//		return true, nil
//	}
//	if !project.IsFile(getFilesRoot() + "/" + 	md5Hash) {
//		f, _ := afero.Af //project.Open(URI)
//		defer f.Close()
//		//bar := progressbar.NewOptions64(f.)
//		project.Copy(path, getFilesRoot() + "/" + 	md5Hash)
//	}
//	AddCacheFile(URI, md5Hash, path)
//	return true, nil
//}
//
//func getFileSize(f afero.File) int64 {
//	fi, err := f.Stat()
//	if err != nil {
//		return 0
//	}
//	return fi.Size()
//}
//
//func GetCachedFile(URI, destination string){
//
//}
//
//func getFilesRoot() string{
//	return cachedir.getCacheRoot() + "/" + cachedir.filesDir
//}
//
//func (c *cache.Cache) getCachedFiles() []cachedFile {
//	cf := cachedFiles{}
//	if project.Exists(cachedir.getCacheRoot() + "/" + filesCacheFile) {
//		toml.DecodeFile(cachedir.getCacheRoot() + "/" +filesCacheFile, &cf) // nolint: errcheck
//	}
//	return cf.File
//}
//
//func isCachedMD5(hash string) bool {
//	for _, f := range getCachedFiles() {
//		if f.MD5 == hash {
//			return true
//		}
//	}
//	return false
//}
//
//func addUriToHash(uri, hash string){
//	files := getCachedFiles()
//	for i, f := range files {
//		if f.MD5 == hash {
//			files[i].URI = append(files[i].URI, uri)
//			writeCachedfilesConfig(files)
//			return
//		}
//	}
//
//}
//func removeUri(uri string){
//	files := getCachedFiles()
//	for i, f := range files {
//		for y, u := range f.URI {
//			if u == uri {
//				files[i].URI = funk.DropString(f.URI, y)
//			}
//		}
//	}
//	writeCachedfilesConfig(files)
//}
//
//func md5Sum(path string) (string, error){
//	f, err := os.Open(path)
//	if err != nil {
//		return ``, err
//	}
//	defer f.Close()
//
//	h := md5.New()
//	if _, err := io.Copy(h, f); err != nil {
//		return ``, err
//	}
//	return fmt.Sprintf("%x", h.Sum(nil)), nil
//}
//
//func writeCachedfilesConfig(files []cachedFile) {
//	data := cachedFiles{File: files}
//	buf := new(bytes.Buffer)
//	err := toml.NewEncoder(buf).Encode(data)
//	if err != nil {
//		output.Printf(output.VERBOSE, "Error encoding projects file: %s", err.Error())
//		return
//	}
//
//	err = ioutil.WriteFile(cachedir.getCacheRoot() + "/" +filesCacheFile, buf.Bytes(), 0644)
//	if err != nil {
//		output.Printf(output.VERBOSE, "Error writing cachedFile file: %s", err.Error())
//		return
//	}
//}