package eviction

import (
	"container/list"
	"log"
)

type LRUEviction struct {
	queue      *list.List
	listPtrMap map[int]*list.Element
}

func NewLRUEviction() *LRUEviction {
	return &LRUEviction{queue: list.New(), listPtrMap: make(map[int]*list.Element)}
}
func (le *LRUEviction) Evict() int {
	nodeToRemove := le.queue.Front()
	k := nodeToRemove.Value.(int)
	le.Remove(k)
	log.Printf("Evicting %d eviction policy ds", k)
	return k
}

func (le *LRUEviction) NotifyAdd(k int) {
	log.Printf("Adding %d in eviction ds", k)
	if kptr, ok := le.listPtrMap[k]; ok {
		le.queue.MoveToBack(kptr)
		return
	}
	le.listPtrMap[k] = le.queue.PushBack(k)
}

func (le *LRUEviction) Remove(k int) {
	log.Printf("Removing %d from eviction ds", k)
	if kptr, ok := le.listPtrMap[k]; ok {
		le.queue.Remove(kptr)
		delete(le.listPtrMap, k)
	}
}

func (le *LRUEviction) NotifyAccessed(k int) {
	log.Printf("accessed %d in eviction ds", k)
	if kptr, ok := le.listPtrMap[k]; ok {
		le.queue.MoveToBack(kptr)
	}
}
