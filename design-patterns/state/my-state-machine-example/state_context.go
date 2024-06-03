package main

import "fmt"

// Software
type SoftwareStateContext struct {
	current SoftwareState
}

func NewSoftwareStateContext() *SoftwareStateContext {
	return &SoftwareStateContext{
		NewStarted(),
	}
}

func (c SoftwareStateContext) print() {
	fmt.Printf("%v\n", c.current)
}

func (s *SoftwareStateContext) Develop() {
	s.current = s.current.Develop()
}

func (s *SoftwareStateContext) Launch() {
	s.current = s.current.Launch()
}

func (s *SoftwareStateContext) Deprecate() {
	s.current = s.current.Deprecate()
}

func (s *Started) Develop() SoftwareState {
	fmt.Println("changing to develop")
	return NewDeveloping()
}

func (s *Started) Launch() SoftwareState {
	fmt.Println("Not allowed")
	return s
}

func (s *Started) Deprecate() SoftwareState {
	fmt.Println("Deprecating")
	return NewDeprecated()
}

func (s *Developing) Develop() SoftwareState {
	fmt.Println("keep developing")
	return s
}

func (s *Developing) Launch() SoftwareState {
	fmt.Println("Launching")
	return NewLaunched()
}

func (s *Developing) Deprecate() SoftwareState {
	fmt.Println("Deprecating")
	return NewDeprecated()
}

func (s *Launched) Develop() SoftwareState {
	fmt.Println("changing to develop")
	return NewDeveloping()
}

func (s *Launched) Launch() SoftwareState {
	fmt.Println("Not allowed")
	return s
}

func (s *Launched) Deprecate() SoftwareState {
	fmt.Println("Deprecating")
	return NewDeprecated()
}

func (s *Deprecated) Develop() SoftwareState {
	fmt.Println("Not allowed")
	return s
}

func (s *Deprecated) Launch() SoftwareState {
	fmt.Println("Not allowed")
	return s
}

func (s *Deprecated) Deprecate() SoftwareState {
	fmt.Println("Already deprecated")
	return s
}
