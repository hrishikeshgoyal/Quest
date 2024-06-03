package remote

import "github.com/hrishikeshgoyal/quest/design-patterns/command/device"

type Command interface {
	execute()
}

type TvOnCommand struct {
	Tv *device.Tv
}

func NewTvOnCommand(tv *device.Tv) *TvOnCommand {
	return &TvOnCommand{
		Tv: tv,
	}
}

func (tvOnCommand *TvOnCommand) execute() {
	tvOnCommand.Tv.On()
}

type TvOffCommand struct {
	Tv *device.Tv
}

func NewTvOffCommand(tv *device.Tv) *TvOffCommand {
	return &TvOffCommand{
		Tv: tv,
	}
}

func (tvOffCommand *TvOffCommand) execute() {
	tvOffCommand.Tv.Off()
}

type AcOnCommand struct {
	Ac *device.Ac
}

func NewAcOnCommand(ac *device.Ac) *AcOnCommand {
	return &AcOnCommand{
		Ac: ac,
	}
}

func (acOnCommand *AcOnCommand) execute() {
	acOnCommand.Ac.On()
}

type AcOffCommand struct {
	Ac *device.Ac
}

func NewAcOffCommand(ac *device.Ac) *AcOffCommand {
	return &AcOffCommand{
		Ac: ac,
	}
}

func (acOffCommand *AcOffCommand) execute() {
	acOffCommand.Ac.Off()
}
