package cache

import "time"

type Cache struct {
	cache map[string]CacheEntry
}

type CacheEntry struct{
	val []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache{
	c := Cache{
		cache: make(map[string]CacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = CacheEntry{
		val: val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheEntry, ok := c.cache[key]
	return cacheEntry.val, ok
}

func (c *Cache) reap(interval time.Duration) {
	MinutesAgo := time.Now().UTC().Add(-interval)
	for key, value := range c.cache {
		if value.createdAt.Before(MinutesAgo) {
			delete(c.cache, key)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}