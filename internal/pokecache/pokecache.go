package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	ExpireAt  time.Time
	Val []byte
}

type Cache struct {
	Entry map[string]cacheEntry
	mutex sync.RWMutex
}

var cacheCleaners = make(map[*Cache]chan struct{})
var cleanersMutex sync.Mutex

func NewCache(clearInterval  ...time.Duration) *Cache {
	cache := &Cache{
		Entry: make(map[string]cacheEntry),
	}

	if len(clearInterval) > 0 && clearInterval[0] > 0 {
		interval := clearInterval[0]
		stopChan := startAutoClear(cache, interval)
		
		// Store the stop channel for later cleanup
		cleanersMutex.Lock()
		cacheCleaners[cache] = stopChan
		cleanersMutex.Unlock()
	}
	return cache
}

func (c *Cache) Add(key string, val []byte, ) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	now := time.Now()
	ttl := 5 * time.Millisecond
	c.Entry[key] = cacheEntry{
		CreatedAt: now,
		ExpireAt:  now.Add(ttl),
		Val: val}
}

func (c *Cache) Get(key string) ([]byte,bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	val, ok := c.Entry[key]

	if !ok {
	return nil, false
	}

	return val.Val, true
}

func (c *Cache) reapLoop() {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	c.Entry = make(map[string]cacheEntry)
}

func startAutoClear(c *Cache, interval time.Duration) chan struct{} {
	stopChan := make(chan struct{})
	ticker := time.NewTicker(interval)
	
	go func() {
		for {
			select {
			case <-ticker.C:
				c.reapLoop()
			case <-stopChan:
				ticker.Stop()
				return
			}
		}
	}()
	
	return stopChan
}

func StopAutoClear(cache *Cache) {
	cleanersMutex.Lock()
	defer cleanersMutex.Unlock()
	
	if stopChan, exists := cacheCleaners[cache]; exists {
		stopChan <- struct{}{}
		close(stopChan)
		delete(cacheCleaners, cache)
	}
}