package main

import (
	"sync"
	"time"
)

type TTLCache struct {
	mu     *sync.RWMutex
	cache  map[string]*CacheItem
	timers map[string]*time.Timer
}

type CacheItem struct {
	Value interface{}
	TTL   time.Time
}

func NewTTLCache() *TTLCache {
	return &TTLCache{
		mu:     &sync.RWMutex{},
		cache:  make(map[string]*CacheItem),
		timers: make(map[string]*time.Timer),
	}
}

func (c *TTLCache) Set(key string, value interface{}, ttl time.Duration) {
	if ttl == 0 {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	if oldTimer, exists := c.timers[key]; exists {
		oldTimer.Stop()
	}

	expiresAt := time.Now().Add(ttl)

	c.cache[key] = &CacheItem{
		Value: value,
		TTL:   expiresAt,
	}

	deleteFunc := func() {
		c.mu.Lock()
		delete(c.cache, key)
		delete(c.timers, key)
		c.mu.Unlock()
	}

	c.timers[key] = time.AfterFunc(ttl, deleteFunc)
}

func (c *TTLCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return item.Value, ok
}
