package device

import "fmt"

type Tv struct {
	isRunning bool
}

func (t *Tv) On() {
	fmt.Println("Turning tv ON")
	t.isRunning = true
}
func (t *Tv) Off() {
	fmt.Println("Turning tv Off")
	t.isRunning = false
}

func (t *Tv) Print() {
	fmt.Println("Tv:", t)
}
