package cachemutex

import (
	"sync"
)

type T int

type Cache struct {
	items map[string]T
	mtx   sync.RWMutex
}

var (
	cache *Cache
)

func New() *Cache {
	cache = &Cache{
		items: make(map[string]T),
	}

	return cache
}

func (c *Cache) Add(key string, val T) {
	c.mtx.Lock()
	c.items[key] = val
	c.mtx.Unlock()
}

func (c *Cache) Get(key string) T {
	c.mtx.RLock()
	val := c.items[key]
	c.mtx.RUnlock()
	return val
}
