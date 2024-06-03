package main

import (
	"github.com/hrishikeshgoyal/quest/design-patterns/command/device"
	"github.com/hrishikeshgoyal/quest/design-patterns/command/remote"
)

func main() {
	tv := &device.Tv{}
	c := remote.NewTvOnCommand(tv)
	b := remote.Button{c}

	tv.Print()
	b.Press()
	tv.Print()

	cc := remote.NewTvOffCommand(tv)
	b = remote.Button{cc}
	tv.Print()
	b.Press()
	tv.Print()

	ac := &device.Ac{}
	a := remote.NewAcOnCommand(ac)
	b = remote.Button{a}
	ac.Print()
	b.Press()
	ac.Print()

	aa := remote.NewAcOffCommand(ac)
	b = remote.Button{aa}
	ac.Print()
	b.Press()
	ac.Print()

}
