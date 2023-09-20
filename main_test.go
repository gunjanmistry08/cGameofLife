package main

import "testing"

func BenchmarkUpgradeSeq(b *testing.B) {
	var grid Grid
	i := 5
	for n := 0; n < b.N; n++ {
		grid = grid.updateGrid(&i)
	}

}
func BenchmarkUpgradeConc(b *testing.B) {
	var grid Grid
	i := 5
	for n := 0; n < b.N; n++ {
		grid = grid.updateGridConc(&i)
	}

}
