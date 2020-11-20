package cache

func New() Cache {
	return NewWithRoot("")
}

func NewWithRoot(root string) Cache {
	cache := Cache {
		root: root,
	}
	cache.createCacheDir()
	return cache
}

