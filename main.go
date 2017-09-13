package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/vectorhacker/nbody/nbody"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	switch e := err.(type) {
	case *strconv.NumError:
		log.Warnf("%s, using maximum 32 bit integer value of %d", e, math.MaxInt32)
		n = math.MaxInt32
	default:
		if err != nil {
			log.Fatal(e)
		}
	}

	bodies := nbody.Default()
	fmt.Printf("%.9f\n", bodies.Energy())
	for i := 0; i < n; i++ {
		bodies.Advance(0.01)
	}
	fmt.Printf("%.9f\n", bodies.Energy())
}
