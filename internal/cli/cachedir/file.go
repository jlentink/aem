package cachedir

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

type cachedFiles struct {
	File cachedFile
}

type cachedFile struct {
	MD5 string
	URL []string
}

func filesRoot() string{
	return fmt.Sprintf("%s/%s", getCacheRoot(), "files")
}
func IsCached(url string) bool {
	return true
}

func CacheFile(url, path string) {

}

func GetCachedFile(url, destination string){

}

func md5Sum(path string) (string, error){
	f, err := os.Open(path)
	if err != nil {
		return ``, err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return ``, err
	}
	return string(h.Sum(nil)), nil
}