package eviction

import "container/list"

type Fifo struct {
	queue *list.List
}

func NewFifoPolicy() *Fifo {
	return &Fifo{
		list.New(),
	}
}

func (f *Fifo) Evict() int {
	e := f.queue.Front()
	f.Remove(e.Value.(int))
	return e.Value.(int)
}

func (f *Fifo) NotifyAdd(k int) {
	for i := f.queue.Front(); i != nil; i = i.Next() {
		if i.Value.(int) == k {
			return
		}
	}
	f.queue.PushBack(k)
}

func (f *Fifo) Remove(k int) {
	for i := f.queue.Front(); i != nil; i = i.Next() {
		if i.Value.(int) == k {
			f.queue.Remove(i)
		}
	}
}

func (f *Fifo) NotifyAccessed(k int) {

}
