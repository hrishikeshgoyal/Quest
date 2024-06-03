package main

const (
	DesStarted    string = "Started"
	DesDeveloping        = "Developing"
	DesLaunched          = "Launched"
	DesDeprecated        = "Deprecated"
)

// software states are
// 1 started -> 2 developing <-> 3. launched -> 4. deprecated
type SoftwareState interface {
	Develop() SoftwareState
	Launch() SoftwareState
	Deprecate() SoftwareState
}

type Started struct {
	Des string
}

func NewStarted() *Started {
	return &Started{
		DesStarted,
	}
}

type Developing struct {
	Des string
}

func NewDeveloping() *Developing {
	return &Developing{
		DesDeveloping,
	}
}

type Launched struct {
	Des string
}

func NewLaunched() *Launched {
	return &Launched{
		DesLaunched,
	}
}

type Deprecated struct {
	Des string
}

func NewDeprecated() *Deprecated {
	return &Deprecated{
		DesDeprecated,
	}
}
