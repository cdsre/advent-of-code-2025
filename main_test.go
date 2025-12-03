package main

import (
	"testing"

	day_3 "github.com/cdsre/advent-of-code-2025/day-3"
	"github.com/cdsre/advent-of-code-2025/helpers"
)

func BenchmarkDay3Puzzle1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := helpers.LoadData("day-3/puzzle_data.txt")
		day_3.Puzzle1(data)
	}
}

func BenchmarkDay3Puzzle2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := helpers.LoadData("day-3/puzzle_data.txt")
		day_3.Puzzle2(data)
	}
}
