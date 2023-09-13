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
			neighbour := make([]uint8, *n)
			// east
			if index2 < *n-1 {
				neighbour = append(neighbour, g[index][index2+1])
			}
			// west
			if index2 > 0 {
				neighbour = append(neighbour, g[index][index2-1])
			}
			// north
			if index > 0 {
				neighbour = append(neighbour, g[index-1][index2])
			}
			// south
			if index < *n-1 {
				neighbour = append(neighbour, g[index+1][index2])
			}
			// north west
			if index > 0 && index2 > 0 {
				neighbour = append(neighbour, g[index-1][index2-1])
			}
			// north east
			if index > 0 && index2 < *n-1 {
				neighbour = append(neighbour, g[index-1][index2+1])
			}
			// south east
			if index < *n-1 && index2 < *n-1 {
				neighbour = append(neighbour, g[index+1][index2+1])
			}
			// south west
			if index < *n-1 && index2 > 0 {
				neighbour = append(neighbour, g[index+1][index2-1])
			}
			if ((sum(neighbour) == 2 || sum(neighbour) == 3) && value == 1) || (sum(neighbour) == 3 && value == 0) {
				newGrid[index][index2] = 1
			} else {
				newGrid[index][index2] = 0
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
