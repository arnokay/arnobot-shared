package cache

type Cacher interface {
	CacheGetter
	CacheSetter
	CacheDeleter
}

type CacheSetter interface {
	Set(key string, value []byte) error
}

type CacheGetter interface {
	Get(key string) ([]byte, error)
}

type CacheDeleter interface {
	Delete(key string) error
}
