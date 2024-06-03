package device

import "fmt"

type Ac struct {
	isRunning bool
}

func (t *Ac) On() {
	fmt.Println("Turning AC ON")
	t.isRunning = true

}
func (t *Ac) Off() {
	fmt.Println("Turning tv Off")
	t.isRunning = false
}

func (t *Ac) Print() {
	fmt.Println("Ac:", t)
}
