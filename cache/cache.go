package pokecache

import "time"

func NewCache() Cache {
	return Cache{
		entries: make(map[string][]string),
	}
}

type Cache struct {
	entries  map[string][]string
	interval time.Duration
}

func (c *Cache) Add(key string, value []string) {
	c.entries[key] = value
}

func (c *Cache) Get(key string) ([]string, bool) {
	e, ok := c.entries[key]
	return e, ok
}

func (c *Cache) NumKeys() int {
	return len(c.entries)
}
