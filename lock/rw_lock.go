package main

import "sync"

type Cache struct {
	data    map[string]any
	rwMutex sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]any),
	}
}

func (c *Cache) Get(key string) (any, bool) {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()
	value, ok := c.data[key]
	return value, ok
}

func (c *Cache) Set(key string, value any) {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()
	c.data[key] = value
}
