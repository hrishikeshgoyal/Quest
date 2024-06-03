package main

import (
	"fmt"
	"github.com/hrishikeshgoyal/quest/design-patterns/functional-options/server"
	"time"
)

func main() {
	s := server.NewServer(server.WithHost("1.2.3.4"), server.WithPort(4444), server.WithTimeout(5*time.Second))
	fmt.Printf("server: %v", s)
}
