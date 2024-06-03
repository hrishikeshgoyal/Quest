package main

import (
	"fmt"
	"github.com/hrishikeshgoyal/quest/design-patterns/strategy/cache"
	"github.com/hrishikeshgoyal/quest/design-patterns/strategy/eviction"
)

func main() {
	fmt.Println("-----------LRU Cache---------")
	sc := cache.NewSimpleCache()
	sc.Add(1, "")
	sc.Add(2, "")
	sc.Get(2)
	sc.Add(3, "")
	sc.Get(1)
	sc.Add(4, "")

	fmt.Println("-----------FIFO Cache---------")
	sc = cache.NewSimpleCache()
	sc.SetPolicy(eviction.NewFifoPolicy())
	sc.Add(1, "")
	sc.Add(2, "")
	sc.Get(2)
	sc.Add(3, "")
	sc.Get(1)
	sc.Add(4, "")
}
