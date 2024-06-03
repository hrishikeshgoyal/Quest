package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var obj Object
var once sync.Once

type Object struct {
	number int
}

func NewObject() Object {
	once.Do(func() {
		obj = Object{number: rand.Int()}
	})
	return obj
}

func main() {
	fmt.Printf("object.number: %v\n", NewObject())
	fmt.Printf("object.number: %v\n", NewObject())
	fmt.Printf("object.number: %v\n", NewObject())
}
