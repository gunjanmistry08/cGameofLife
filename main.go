package main

// import (
// 	"flag"
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// type Grid [][]uint8
// type Cell bool

// type job struct {
// 	old Grid
// 	row uint8
// 	col uint8
// }
// type result struct {
// 	newState uint8
// 	row      uint8
// 	col      uint8
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	n := flag.Int("n", 5, "Value of \"n\"")

// 	flag.Parse()

// 	fmt.Println("Hello World", *n)

// 	// jobs := make(chan job)
// 	// results := make(chan result)

// 	// for i := 0; i < 5; i++ {
// 	// 	go worker(i, jobs, results)
// 	// }
// 	grid := initGrid(n, true)
// 	for i := 0; ; i++ {
// 		fmt.Printf("\nGeneration: %v\n", i)
// 		grid.drawGrid()
// 		grid = grid.updateGrid(n)
// 		time.Sleep(3 * time.Second)
// 	}
// }

// func (grid Grid) getNeighbors(row, col int) uint8 {
// 	var neighbors []uint8
// 	for i := row - 1; i <= row+1; i++ {
// 		for j := col - 1; j <= col+1; j++ {
// 			if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
// 				if i != row || j != col {
// 					neighbors = append(neighbors, grid[i][j])
// 				}
// 			}
// 		}
// 	}
// 	return sum(neighbors)
// }

// func sum(ar []uint8) (r uint8) {
// 	for _, v := range ar {
// 		r += v
// 	}
// 	return r
// }

// func (g Grid) updateGrid(n *int) (newGrid Grid) {
// 	newGrid = initGrid(n, false)
// 	for index := range g {
// 		for index2, value := range g[index] {
// 			neighbour := g.getNeighbors(index, index2)
// 			if neighbour == 3 || (neighbour == 2 && value == 1) {
// 				newGrid[index][index2] = 1
// 			}
// 		}
// 	}
// 	return newGrid
// }
// func (g Grid) updateGridConc(n *int) (newGrid Grid) {
// 	var wg sync.WaitGroup
// 	newGrid = initGrid(n, false)
// 	for index := range g {
// 		for index2 := range g[index] {
// 			wg.Add(1)
// 			go func(row, col int) {
// 				neighbour := g.getNeighbors(row, col)
// 				if neighbour == 3 || (neighbour == 2 && g[row][col] == 1) {
// 					newGrid[row][col] = 1
// 				}
// 				wg.Done()
// 			}(index, index2)
// 		}
// 	}
// 	wg.Wait()
// 	return newGrid
// }

// // func (g Grid) updateGridWP(n *int, jobs chan job, results chan result) (newGrid Grid) {
// // 	newGrid = initGrid(n, false)
// // 	for index := range g {
// // 		for index2 := range g[index] {
// // 			js := job{
// // 				row: uint8(index),
// // 				col: uint8(index2),
// // 				old: g,
// // 			}
// // 			jobs <- js
// // 			rs := <-results
// // 			newGrid[rs.row][rs.col] = rs.newState
// // 		}
// // 	}
// // 	return newGrid
// // }
// // func (g Grid) updateGridPool(n *int, pool *sync.Pool) (newGrid Grid) {
// // 	newGrid = initGrid(n, false)
// // 	var wg sync.WaitGroup
// // 	for index := range g {
// // 		for index2 := range g[index] {
// // 			wg.Add(1)
// // 			go func(row, col int) {
// // 				neighbours := g.getNeighbors(row, col)
// // 				if neighbours == 3 || (neighbours == 2 && g[row][col] == 1) {
// // 					newGrid[row][col] = 1
// // 				}
// // 				pool.Put(&row)
// // 				wg.Done()
// // 			}(index, index2)

// // 		}
// // 	}
// // 	wg.Wait()
// // 	return newGrid
// // }

// func (g Grid) drawGrid() {
// 	var stringGrid string
// 	for index := range g {
// 		for _, value2 := range g[index] {
// 			stringGrid += fmt.Sprintf("| %v |", value2)
// 		}
// 		stringGrid += "\n"
// 	}
// 	fmt.Print(stringGrid)
// }

// func initGrid(n *int, premier bool) (nascentGrid Grid) {
// 	nascentGrid = make(Grid, *n)
// 	for index := range nascentGrid {
// 		nascentGrid[index] = make([]uint8, *n)
// 		for index2 := range nascentGrid[index] {
// 			if premier {
// 				nascentGrid[index][index2] = uint8(rand.Intn(2))
// 			} else {
// 				nascentGrid[index][index2] = 0
// 			}
// 		}
// 	}
// 	return nascentGrid
// }

// // func worker(id int, jobs <-chan job, results chan<- result) {
// // 	for j := range jobs {
// // 		rs := result{}
// // 		rs.row = j.row
// // 		rs.col = j.col
// // 		nn := j.old.getNeighbors(int(j.row), int(j.col))
// // 		if nn == 3 || (nn == 2 && j.old[j.row][j.col] == 1) {
// // 			rs.newState = 1
// // 		}
// // 		results <- rs
// // 	}
// // }
