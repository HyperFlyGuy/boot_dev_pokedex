package pokecache

import(
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	mu sync.Mutex
	cache map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()         // Acquire lock
	defer c.mu.Unlock() // Ensure unlock when method returns
	entry := cacheEntry{createdAt: time.Now(), val: val}
	c.cache[key] = entry
}

func (c *Cache) Get(key string) ([]byte,bool){
	c.mu.Lock()         // Acquire lock
	defer c.mu.Unlock() // Ensure unlock when method returns
	value, ok := c.cache[key]
	if ok {
		return value.val, true
	} else{
			return nil, false
	}
}

func (c *Cache) reapLoop(i time.Duration){
	ticker := time.NewTicker(i)
	for range ticker.C {
		c.mu.Lock()         // Acquire lock
		for key, entry := range c.cache{
			now := time.Now()
			if (now.Sub(entry.createdAt)) > i {
				delete(c.cache, key)
			}
	 	}
		c.mu.Unlock() // Ensure unlock when method returns
	}
}
