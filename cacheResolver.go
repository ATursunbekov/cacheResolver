package cacheResolver

type CacheResolver map[string]interface{}

func (cache CacheResolver) Get(key string) (interface{}, bool) {
	val, ok := cache[key]
	return val, ok
}

func (cache CacheResolver) Put(key string, val interface{}) {
	cache[key] = val
}

func (cache CacheResolver) Del(key string) {
	delete(cache, key)
}
