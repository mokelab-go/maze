package main

import (
	"fmt"
	"math/rand"
	"time"
)

type maze struct {
	width        int
	height       int
	actualWidth  int
	actualHeight int
	data         []int
}

func main() {

	rand.Seed(time.Now().UnixNano())
	//rand.Seed(1)
	maze := makeMaze(10, 10)
	maze.generate()
	maze.print()
}

func makeMaze(w, h int) *maze {
	actualWidth := w*2 + 1
	actualHeight := h*2 + 1
	out := &maze{
		width:        w,
		height:       h,
		actualWidth:  actualWidth,
		actualHeight: actualHeight,
		data:         make([]int, actualWidth*actualHeight),
	}
	for x := 0; x < actualWidth; x++ {
		out.data[x] = 1
		out.data[x+actualWidth*(actualHeight-1)] = 1
	}
	for y := 0; y < actualWidth; y++ {
		out.data[y*actualWidth] = 1
		out.data[(actualWidth-1)+y*actualWidth] = 1
	}

	return out
}

func (m *maze) generate() {
	step := 2
	for {
		if !m.hasEmpty() {
			break
		}
		x := rand.Intn(m.width)*2 + 2
		y := rand.Intn(m.height)*2 + 2
		if m.data[x+y*m.actualWidth] != 0 {
			continue
		}
		stack := make([]int, 0)
		_ = stack
		for {
			m.data[x+y*m.actualWidth] = step
			directions := generateDirections()
			done := false
			moved := false
			for _, direction := range directions {
				if direction == 0 {
					// upper
					if m.data[x+(y-1)*m.actualWidth] > 0 {
						continue
					}
					if m.data[x+(y-2)*m.actualWidth] == step {
						continue
					}
					if m.data[x+(y-2)*m.actualWidth] > 0 {
						m.data[x+(y-1)*m.actualWidth] = step
						done = true
						break
					}
					m.data[x+(y-1)*m.actualWidth] = step
					y -= 2
					moved = true
					break
				} else if direction == 1 {
					// left
					if m.data[x-1+y*m.actualWidth] > 0 {
						continue
					}
					if m.data[x-2+y*m.actualWidth] == step {
						continue
					}
					if m.data[x-2+y*m.actualWidth] > 0 {
						m.data[x-1+y*m.actualWidth] = step
						done = true
						break
					}
					m.data[x-1+y*m.actualWidth] = step
					x -= 2
					moved = true
					break
				} else if direction == 2 {
					// right
					if m.data[x+1+y*m.actualWidth] > 0 {
						continue
					}
					if m.data[x+2+y*m.actualWidth] == step {
						continue
					}
					if m.data[x+2+y*m.actualWidth] > 0 {
						m.data[x+1+y*m.actualWidth] = step
						done = true
						break
					}
					m.data[x+1+y*m.actualWidth] = step
					x += 2
					moved = true
					break
				} else {
					// under
					if m.data[x+(y+1)*m.actualWidth] > 0 {
						continue
					}
					if m.data[x+(y+2)*m.actualWidth] == step {
						continue
					}
					if m.data[x+(y+2)*m.actualWidth] > 0 {
						m.data[x+(y+1)*m.actualWidth] = step
						done = true
						break
					}
					m.data[x+(y+1)*m.actualWidth] = step
					y += 2
					moved = true
					break
				}
			}
			if done {
				break
			}
			if !moved {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				x = v % m.actualWidth
				y = v / m.actualWidth
			}
		}
		step++
	}
}

func (m *maze) hasEmpty() bool {
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			if m.data[x*2+2+(y*2+2)*m.actualWidth] == 0 {
				return true
			}
		}
	}
	return false
}

func (m *maze) print() {
	for y := 0; y < m.actualHeight; y++ {
		for x := 0; x < m.actualWidth; x++ {
			if m.data[x+y*m.actualWidth] >= 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func generateDirections() []int {
	out := []int{0, 1, 2, 3}
	rand.Shuffle(len(out), func(i, j int) {
		out[i], out[j] = out[j], out[i]
	})
	return out
}
