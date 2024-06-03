package cache

import (
	"github.com/hrishikeshgoyal/quest/design-patterns/strategy/eviction"
	"github.com/hrishikeshgoyal/quest/design-patterns/strategy/store"
	"log"
)

type SimpleCache struct {
	store        store.Store
	evictionAlgo eviction.EvictionAlgo
}

func NewSimpleCache() *SimpleCache {
	return &SimpleCache{
		store:        store.NewHashMapStore(),
		evictionAlgo: eviction.NewLRUEviction(),
	}
}

func (s *SimpleCache) SetPolicy(e eviction.EvictionAlgo) {
	s.evictionAlgo = e
}

func (c *SimpleCache) Add(k int, v string) {
	log.Printf("Adding %d, %s to cache", k, v)
	if c.store.IsFull() {
		keyToEvict := c.evictionAlgo.Evict()
		c.store.Remove(keyToEvict)
	}
	c.store.Add(k, v)
	c.evictionAlgo.NotifyAdd(k)
}

func (c *SimpleCache) Remove(k int) {
	log.Printf("Removing %d from cache store", k)
	c.store.Remove(k)
	c.evictionAlgo.Remove(k)
}

func (c *SimpleCache) Get(k int) (string, error) {
	log.Printf("Getting %d from cache", k)
	v, err := c.store.Get(k)
	if err != nil {
		return v, err
	}
	c.evictionAlgo.NotifyAccessed(k)
	return v, nil
}
