package pokecache

import (
	"fmt"
	"sync"
	"time"
)

func NewCache(timeout time.Duration) Cache {
	ticker := time.NewTicker(timeout)

	var mutex sync.Mutex

	c := Cache{
		entries: make(map[string]cacheEntry),
		mu:      &mutex,
	}

	go func() {
		for {
			select {
			case <-ticker.C:
				c.pruneCache(timeout)
			}
		}
	}()

	return c
}

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		value:     value,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	v, ok := c.entries[key]
	c.mu.Unlock()
	return v.value, ok
}

func (c *Cache) Remove(key string) {
	c.mu.Lock()
	delete(c.entries, key)
	c.mu.Unlock()
}
func (c *Cache) NumKeys() int {
	return len(c.entries)
}

func (c *Cache) pruneCache(timeout time.Duration) {
	for k, v := range c.entries {
		t := time.Now()
		if t.Sub(v.createdAt) > timeout {
			c.Remove(k)
			fmt.Printf("removed %v after %v\n", k, timeout)
		}
	}
}
