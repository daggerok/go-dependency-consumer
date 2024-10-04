package main

import (
	"fmt"
	"github.com/daggerok/go-dependency-producer/lib"
	"os"
)

func main() {
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	greeting := lib.Greet(name)
	fmt.Println(greeting)
}
