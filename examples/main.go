package main

import (
	"math/rand"
	"time"

	"github.com/mokelab-go/maze"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	maze := maze.New(20, 20)
	maze.Generate(r)
	maze.Print()
}
