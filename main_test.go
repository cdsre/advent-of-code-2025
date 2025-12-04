package main

import (
	"testing"

	"github.com/cdsre/advent-of-code-2025/day-1"
	"github.com/cdsre/advent-of-code-2025/day-2"
	"github.com/cdsre/advent-of-code-2025/day-3"
	day_4 "github.com/cdsre/advent-of-code-2025/day-4"
	"github.com/cdsre/advent-of-code-2025/helpers"
)

func benchWrapper[T any](b *testing.B, name string, f func(data T) int, data T) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f(data)
		}
	})
}

func BenchmarkD1P1(b *testing.B) {
	data := helpers.LoadData("day-1/puzzle_data.txt")
	benchWrapper(b, "D1P1", day_1.Puzzle1, data)
}

func BenchmarkD1P2(b *testing.B) {
	data := helpers.LoadData("day-1/puzzle_data.txt")
	benchWrapper(b, "D1P2", day_1.Puzzle2, data)
}

func BenchmarkD2P1(b *testing.B) {
	data := helpers.LoadData("day-2/puzzle_data.txt")
	benchWrapper(b, "D2P1", day_2.Puzzle1, data[0])
}

func BenchmarkD2P2(b *testing.B) {
	data := helpers.LoadData("day-2/puzzle_data.txt")
	benchWrapper(b, "D2P2", day_2.Puzzle2, data[0])
}

func BenchmarkD3P1(b *testing.B) {
	data := helpers.LoadData("day-3/puzzle_data.txt")
	benchWrapper(b, "D3P1", day_3.Puzzle1, data)
}

func BenchmarkD3P2(b *testing.B) {
	data := helpers.LoadData("day-3/puzzle_data.txt")
	benchWrapper(b, "D3P2", day_3.Puzzle2, data)
}

func BenchmarkD4P1(b *testing.B) {
	data := helpers.LoadData("day-4/puzzle_data.txt")
	benchWrapper(b, "D4P1", day_4.Puzzle1, data)
}

func BenchmarkD4P2(b *testing.B) {
	data := helpers.LoadData("day-4/puzzle_data.txt")
	benchWrapper(b, "D4P2", day_4.Puzzle2, data)
}
