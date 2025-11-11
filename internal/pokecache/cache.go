package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]CacheEntry
	mu *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time 
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	entries := make(map[string]CacheEntry)
	cache := Cache{
		entries: entries,
		mu: &sync.Mutex{},
	}

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
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
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.entries[key]
	if !exists {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(t time.Duration) {
	// Cleans up old entries past expiration
	now := time.Now()
	c.mu.Lock()
	for key, value := range c.entries {
		if now.Sub(value.createdAt) >= t {
			delete(c.entries, key)
		}
	}
	c.mu.Unlock()
}