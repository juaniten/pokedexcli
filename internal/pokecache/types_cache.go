package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:  map[string]cacheEntry{},
		mu:       &sync.Mutex{},
		interval: interval,
	}
	c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) (val []byte, found bool) {
	entry, found := c.entries[key]
	val = nil
	if found {
		val = entry.val
	}
	return val, found
}

func (c *Cache) reapLoop() {

}
