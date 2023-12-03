package maze

import (
	"fmt"
	"math/rand"
)

const (
	directionUp = iota
	directionLeft
	directionRight
	directionDown
)

type Maze interface {
	// Generate generates a maze
	Generate(r *rand.Rand)

	// Print prints the maze
	Print()
}

type maze struct {
	cellWidth  int
	cellHeight int
	width      int
	height     int
	data       []int
}

// New creates a new maze
func New(w, h int) Maze {
	actualWidth := w*2 + 1
	actualHeight := h*2 + 1
	out := &maze{
		cellWidth:  w,
		cellHeight: h,
		width:      actualWidth,
		height:     actualHeight,
		data:       make([]int, actualWidth*actualHeight),
	}
	return out
}

func (m *maze) Generate(r *rand.Rand) {
	m.setEdgeWalls()
	cells := generateCells(r, m.cellWidth, m.cellHeight)
	for c, cell := range cells {
		x := (cell%m.cellWidth)*2 + 2
		y := (cell/m.cellWidth)*2 + 2
		if m.isWall(x, y) {
			continue
		}
		// wall id 1 = edge wall
		wallId := c + 2
		stack := newStack()
		for {
			m.set(x, y, wallId)
			directions := generateDirections(r)
			done := false
			moved := false
			for _, direction := range directions {
				if direction == directionUp {
					if m.isWall(x, y-1) {
						continue
					}
					if m.get(x, y-2) == wallId {
						continue
					}
					// connect to other wall?
					if m.isWall(x, y-2) {
						m.set(x, y-1, wallId)
						done = true
						break
					}
					m.set(x, y-1, wallId)
					stack.push(x, y)
					y -= 2
					moved = true
					break
				} else if direction == directionLeft {
					if m.isWall(x-1, y) {
						continue
					}
					if m.get(x-2, y) == wallId {
						continue
					}
					// connect to other wall?
					if m.isWall(x-2, y) {
						m.set(x-1, y, wallId)
						done = true
						break
					}
					m.set(x-1, y, wallId)
					stack.push(x, y)
					x -= 2
					moved = true
					break
				} else if direction == directionRight {
					// right
					if m.isWall(x+1, y) {
						continue
					}
					if m.get(x+2, y) == wallId {
						continue
					}
					// connect to other wall?
					if m.isWall(x+2, y) {
						m.set(x+1, y, wallId)
						done = true
						break
					}
					m.set(x+1, y, wallId)
					stack.push(x, y)
					x += 2
					moved = true
					break
				} else {
					// directionDown
					if m.isWall(x, y+1) {
						continue
					}
					if m.get(x, y+2) == wallId {
						continue
					}
					// connect to other wall?
					if m.isWall(x, y+2) {
						m.set(x, y+1, wallId)
						done = true
						break
					}
					m.set(x, y+1, wallId)
					stack.push(x, y)
					y += 2
					moved = true
					break
				}
			}
			if done {
				break
			}
			if !moved {
				x, y = stack.pop()
			}
		}
	}
}

func (m *maze) Print() {
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			if m.data[x+y*m.width] >= 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (m *maze) get(x, y int) int {
	return m.data[x+y*m.width]
}

func (m *maze) set(x, y, value int) {
	m.data[x+y*m.width] = value
}

func (m *maze) isWall(x, y int) bool {
	return m.get(x, y) > 0
}

func (m *maze) setEdgeWalls() {
	for x := 0; x < m.width; x++ {
		m.set(x, 0, 1)
		m.set(x, m.height-1, 1)
	}
	for y := 0; y < m.height; y++ {
		m.set(0, y, 1)
		m.set(m.width-1, y, 1)
	}
}

func generateCells(r *rand.Rand, w, h int) []int {
	out := make([]int, w*h)
	for i := 0; i < len(out); i++ {
		out[i] = i
	}
	r.Shuffle(len(out), func(i, j int) {
		out[i], out[j] = out[j], out[i]
	})
	return out
}

func generateDirections(r *rand.Rand) []int {
	out := []int{directionUp, directionLeft, directionRight, directionDown}
	r.Shuffle(len(out), func(i, j int) {
		out[i], out[j] = out[j], out[i]
	})
	return out
}

type stack struct {
	x []int
	y []int
}

func newStack() *stack {
	return &stack{
		x: make([]int, 0),
		y: make([]int, 0),
	}
}

func (s *stack) push(x, y int) {
	s.x = append(s.x, x)
	s.y = append(s.y, y)
}

func (s *stack) pop() (int, int) {
	x := s.x[len(s.x)-1]
	y := s.y[len(s.y)-1]
	s.x = s.x[:len(s.x)-1]
	s.y = s.y[:len(s.y)-1]
	return x, y
}
