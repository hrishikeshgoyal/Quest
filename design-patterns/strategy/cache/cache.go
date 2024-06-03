package cache

import (
	"github.com/hrishikeshgoyal/quest/design-patterns/strategy/eviction"
)

type Cache interface {
	Add(k int, v string)

	Get(k int) (string, error)
	Remove(k int)
	SetPolicy(e eviction.EvictionAlgo)
}
