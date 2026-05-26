package pokecache

import (
	"sync"
	"time"
)

// Define a cache struct
type Cache struct {
	entries  map[string]cacheEntry
	interval time.Duration
	mutex    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Create a new cache
func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
		mutex:    &sync.Mutex{},
	}

	//recursively clean the cache
	go cache.reapLoop()

	return cache
}

// A basic imple,entation of LRU Cache
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
		c.mutex.Unlock()
	}
}

// Add item to cache
func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get item from cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}
