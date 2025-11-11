package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entries map[string]CacheEntry
	mu sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time 
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	entries := make(map[string]CacheEntry)
	cache := Cache{
		entries: entries,
	}

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for t := range ticker.C {
			fmt.Println(t)
			cache.reapLoop(interval)
		}
	}()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = CacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// Returns true if found, else false
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.entries[key]
	if !exists {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(t time.Duration) {
	// invoked in new cache, cleans up old entries past intervals
	now := time.Now()
	c.mu.Lock()
	for key, value := range c.entries {
		if now.Sub(value.createdAt) >= t {
			delete(c.entries, key)
		}
	}
	c.mu.Unlock()
}