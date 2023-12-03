package main

import (
	"math/rand"
	"time"

	"github.com/mokelab-go/maze"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//rand.Seed(1)
	maze := maze.New(20, 20)
	maze.Generate()
	maze.Print()
}
