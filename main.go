package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type Grid [][]uint8

func main() {
	n := flag.Int("n", 5, "Value of \"n\"")

	flag.Parse()
	fmt.Println("Hello World", *n)
	grid := initGrid(n, true)
	var generation int
	for {
		fmt.Printf("\nGeneration: %v\n", generation)
		drawGrid(grid)
		grid = updateGrid(grid, n)
		generation++
		time.Sleep(3 * time.Second)
	}

}

func getNeighbors(grid Grid, row, col int) (neighbors []uint8) {
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
				if i != row || j != col {
					neighbors = append(neighbors, grid[i][j])
				}
			}
		}
	}
	return neighbors
}

func sum(ar []uint8) (r uint8) {
	for _, v := range ar {
		r += v
	}
	return r
}

func updateGrid(g Grid, n *int) (newGrid Grid) {
	rand.Seed(time.Now().UnixNano())
	newGrid = initGrid(n, false)
	for index := range g {
		for index2, value := range g[index] {
			neighbour := getNeighbors(g, index, index2)
			if (value == 1 && (sum(neighbour) == 2 || sum(neighbour) == 3)) || (value != 1 && sum(neighbour) == 3) {
				newGrid[index][index2] = 1
			}
		}
	}
	return newGrid
}

func drawGrid(g Grid) {
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
	nascentGrid = make([][]uint8, *n)
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
