package cache

import "time"

// CachedFile Cached file
type CachedFile struct {
	MD5 string
	URI []string
	Date time.Time
	OriginalName string
}
