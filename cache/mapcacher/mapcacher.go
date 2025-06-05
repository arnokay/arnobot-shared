package mapcacher

type MapCacher struct {
	store map[string]string
}

func New() MapCacher {
	return MapCacher{
		store: make(map[string]string),
	}
}

func (c *MapCacher) Get(key string) (string, error) {
	val, ok := c.store[key]
	if !ok {
		return "", nil
	}

	return val, nil
}

func (c *MapCacher) Set(key, value string) error {
	c.store[key] = value

	return nil
}

type MapCacherWithPrefix struct {
	mapCacher *MapCacher
	prefix    string
}

func WithPrefix(cacher *MapCacher, prefix string) MapCacherWithPrefix {
	return MapCacherWithPrefix{
		mapCacher: cacher,
		prefix:    prefix,
	}
}

func (c *MapCacherWithPrefix) Get(key string) (string, error) {
	val, err := c.mapCacher.Get(c.prefix + ":" + key)

	return val, err
}

func (c *MapCacherWithPrefix) Set(key, value string) error {
	err := c.mapCacher.Set(c.prefix+":"+key, value)

	return err
}
