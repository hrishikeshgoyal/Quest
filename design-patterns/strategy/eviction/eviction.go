package eviction

type EvictionAlgo interface {
	Evict() int

	NotifyAdd(k int)

	NotifyAccessed(k int)
	Remove(k int)
}
