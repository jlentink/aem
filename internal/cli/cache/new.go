package cache

// New Cache
func New() Cache {
	return NewWithRoot("")
}

// NewWithRoot cache root
func NewWithRoot(root string) Cache {
	cache := Cache {
		root: root,
	}
	cache.createCacheDir()
	return cache
}

