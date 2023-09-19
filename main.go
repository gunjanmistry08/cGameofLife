package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type Grid [][]uint8
type Cell bool

func main() {
	rand.Seed(time.Now().UnixNano())
	n := flag.Int("n", 5, "Value of \"n\"")

	flag.Parse()
	fmt.Println("Hello World", *n)
	grid := initGrid(n, true)
	var generation int
	for {
		fmt.Printf("\nGeneration: %v\n", generation)
		grid.drawGrid()
		grid = grid.updateGrid(n)
		generation++
		time.Sleep(3 * time.Second)
	}

}

func (grid Grid) getNeighbors(row, col int) uint8 {
	var neighbors []uint8
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
				if i != row || j != col {
					neighbors = append(neighbors, grid[i][j])
				}
			}
		}
	}
	return sum(neighbors)
}

func sum(ar []uint8) (r uint8) {
	for _, v := range ar {
		r += v
	}
	return r
}

func (g Grid) updateGrid(n *int) (newGrid Grid) {
	newGrid = initGrid(n, false)
	for index := range g {
		for index2, value := range g[index] {
			neighbour := g.getNeighbors(index, index2)
			if neighbour == 3 || (neighbour == 2 && value == 1) {
				newGrid[index][index2] = 1
			}
		}
	}
	return newGrid
}

func (g Grid) drawGrid() {
	var stringGrid string
	for index := range g {
		for _, value2 := range g[index] {
			stringGrid += fmt.Sprintf("| %v |", value2)
		}
		stringGrid += "\n"
	}
	fmt.Print(stringGrid)
}

func initGrid(n *int, premier bool) (nascentGrid Grid) {
	nascentGrid = make(Grid, *n)
	for index := range nascentGrid {
		nascentGrid[index] = make([]uint8, *n)
		for index2 := range nascentGrid[index] {
			if premier {
				nascentGrid[index][index2] = uint8(rand.Intn(2))
			} else {
				nascentGrid[index][index2] = 0
			}
		}
	}
	return nascentGrid
}
