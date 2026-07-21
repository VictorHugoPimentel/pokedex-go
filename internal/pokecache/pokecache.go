package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]CacheEntry
	mux sync.Mutex
}

type CacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(internal time.Duration) Cache {
	c := Cache{
		cache: make(map[string]CacheEntry),
		mux: sync.Mutex{},
	}
	go c.ReapLoop(internal)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = CacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

func (c *Cache) ReapLoop(internal time.Duration) {
	ticker := time.NewTicker(internal)
	for range ticker.C {
		c.reap(internal)
	}
}

func (c *Cache) reap(internal time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	timeAgo := time.Now().UTC().Add(-internal)
	for k, v := range c.cache {
		if v.createdAt.Before(timeAgo) {
			delete(c.cache, k)
		}
	}
}
