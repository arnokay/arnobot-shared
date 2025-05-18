package pkg

type Cacher interface {
	CacheGetter
	CacheSetter
}

type CacheSetter interface {
	Set(key string, value string) error
}

type CacheGetter interface {
	Get(key string) (string, error)
}
