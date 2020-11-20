package cache

import "time"

type CachedFile struct {
	MD5 string
	URI []string
	Date time.Time
	OriginalName string
}
