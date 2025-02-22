package cacheResolver

import (
	"sync"
	"time"
)

type CacheResolver struct {
	mut   sync.Mutex
	cache map[string]interface{}
}

func NewCacheResolver() *CacheResolver {
	return &CacheResolver{
		cache: make(map[string]interface{}),
	}
}

func (c *CacheResolver) Get(key string) (interface{}, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()
	val, ok := c.cache[key]
	return val, ok
}

func (c *CacheResolver) Put(key string, val interface{}) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.cache[key] = val
}

func (c *CacheResolver) Del(key string) {
	c.mut.Lock()
	defer c.mut.Unlock()
	delete(c.cache, key)
}

func (c *CacheResolver) Set(key string, value interface{}, ttl time.Duration) {
	c.Put(key, value)

	go func() {
		time.Sleep(ttl)
		c.Del(key)
	}()
}
