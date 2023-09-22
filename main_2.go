package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	Width  = 40
	Height = 20
)

type Cell struct {
	Alive bool
}

type Grid struct {
	Cells [Width][Height]Cell
}

func (g *Grid) initialize() {
	for x := 0; x < Width; x++ {
		for y := 0; y < Height; y++ {
			g.Cells[x][y] = Cell{Alive: rand.Intn(2) == 1}
		}
	}
}

func (g *Grid) draw() {
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			if g.Cells[x][y].Alive {
				fmt.Print("■")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (g *Grid) update() {
	newGrid := Grid{}

	for x := 0; x < Width; x++ {
		for y := 0; y < Height; y++ {
			aliveNeighbors := g.countAliveNeighbors(x, y)
			if g.Cells[x][y].Alive {
				newGrid.Cells[x][y].Alive = aliveNeighbors == 2 || aliveNeighbors == 3
			} else {
				newGrid.Cells[x][y].Alive = aliveNeighbors == 3
			}
		}
	}

	g.Cells = newGrid.Cells
}

func (g *Grid) updateConc() {
	newGrid := Grid{}
	var wg sync.WaitGroup
	wg.Add(Width)

	for x := 0; x < Width; x++ {
		go func(x int) {
			defer wg.Done()
			for y := 0; y < Height; y++ {
				aliveNeighbors := g.countAliveNeighbors(x, y)
				if g.Cells[x][y].Alive {
					newGrid.Cells[x][y].Alive = aliveNeighbors == 2 || aliveNeighbors == 3
				} else {
					newGrid.Cells[x][y].Alive = aliveNeighbors == 3
				}
			}
		}(x)
	}

	wg.Wait()
	g.Cells = newGrid.Cells
}

func (g *Grid) countAliveNeighbors(x, y int) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && ny >= 0 && nx < Width && ny < Height {
				if g.Cells[nx][ny].Alive {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	rand.Seed(time.Now().UnixNano())

	grid := Grid{}
	grid.initialize()

	for {
		fmt.Print("\x0c") // Clear the screen
		grid.draw()
		grid.update()
		time.Sleep(100 * time.Millisecond)
	}
}
