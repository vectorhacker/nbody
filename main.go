package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/vectorhacker/nbody/nbody"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	bodies := nbody.NewSystem()
	fmt.Printf("%.9f\n", bodies.Energy())
	for i := 0; i < n; i++ {
		bodies.Advance(0.01)
	}
	fmt.Printf("%.9f\n", bodies.Energy())
}
