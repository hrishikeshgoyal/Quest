package store

type Store interface {
	Add(k int, v string)

	Get(k int) (string, error)
	Remove(k int)
	IsFull() bool
}
