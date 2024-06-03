package store

import (
	"fmt"
	"log"
)

type HashMapStore struct {
	store    map[int]string
	capacity int
	size     int
}

func NewHashMapStore() *HashMapStore {
	return &HashMapStore{
		store:    make(map[int]string, 3),
		capacity: 3,
		size:     0,
	}
}

func (c *HashMapStore) Add(k int, v string) {
	log.Printf("Adding %d, %s to hashmap store", k, v)
	if _, ok := c.store[k]; ok {
		log.Printf("Key %d already exists in hashmap store", k)
		c.store[k] = v
		return
	}

	if c.size == c.capacity {
		log.Printf("hashmap store is full, first remove element then store")
		return
	}
	c.size++
	c.store[k] = v
}

func (c *HashMapStore) Remove(k int) {
	log.Printf("Removing %d from hashmap store", k)
	delete(c.store, k)
	c.size--
}

func (c *HashMapStore) IsFull() bool {
	return c.size == c.capacity
}

func (c *HashMapStore) Get(k int) (string, error) {
	log.Printf("Get %d from store", k)
	if v, ok := c.store[k]; ok {
		return v, nil
	}
	return "", fmt.Errorf("Not Found")
}
