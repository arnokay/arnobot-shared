package mapcacher

import "sync"

type MapCacher struct {
	store map[string][]byte
	mx    sync.RWMutex
}

func New() MapCacher {
	return MapCacher{
		store: make(map[string][]byte),
	}
}

func (c *MapCacher) Get(key string) ([]byte, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.store[key]
	if !ok {
		return nil, nil
	}

	return val, nil
}

func (c *MapCacher) Set(key string, value []byte) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.store[key] = value

	return nil
}

func (c *MapCacher) Delete(key string) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.store, key)

	return nil
}

type MapCacherWithPrefix struct {
	mapCacher *MapCacher
	prefix    string
}

func WithPrefix(cacher *MapCacher, prefix string) *MapCacherWithPrefix {
	return &MapCacherWithPrefix{
		mapCacher: cacher,
		prefix:    prefix,
	}
}

func (c *MapCacherWithPrefix) Get(key string) ([]byte, error) {
	val, err := c.mapCacher.Get(c.prefix + "." + key)

	return val, err
}

func (c *MapCacherWithPrefix) Set(key string, value []byte) error {
	err := c.mapCacher.Set(c.prefix+"."+key, value)

	return err
}

func (c *MapCacherWithPrefix) Delete(key string) error {
	return c.mapCacher.Delete(c.prefix + "." + key)
}
