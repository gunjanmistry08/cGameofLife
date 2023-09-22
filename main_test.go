package main

// "sync"

// func BenchmarkUpgradeSeq(b *testing.B) {
// 	var grid Grid
// 	i := 5
// 	for n := 0; n < b.N; n++ {
// 		grid = grid.updateGrid(&i)
// 	}

// }

// func BenchmarkUpgradeConc(b *testing.B) {
// 	var grid Grid
// 	i := 5
// 	for n := 0; n < b.N; n++ {
// 		grid = grid.updateGridConc(&i)
// 	}

// }

// func BenchmarkUpgradeWP(b *testing.B) {
// 	var grid Grid
// 	i := 5
// 	jobs := make(chan job)
// 	results := make(chan result)
// 	for i := 0; i < 5; i++ {
// 		go worker(i, jobs, results)
// 	}
// 	for n := 0; n < b.N; n++ {
// 		grid = grid.updateGridWP(&i, jobs, results)
// 	}

// }

// func BenchmarkUpgradePool(b *testing.B) {
// 	var grid Grid
// 	var p *sync.Pool
// 	for n := 0; n < b.N; n++ {
// 		grid = grid.updateGridPool(&n, p)
// 	}

// }

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkUpdate(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	grid := Grid{}
	grid.initialize()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		grid.updateConc()
	}
}

func BenchmarkUpdateNoConcurrency(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	grid := Grid{}
	grid.initialize()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		grid.update()
	}
}
