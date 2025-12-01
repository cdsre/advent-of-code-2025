package main

import (
	"fmt"

	"github.com/cdsre/advent-of-code-2025/day-1"
	"github.com/cdsre/advent-of-code-2025/helpers"
)

func main() {
	print(day1Puzzle2())
}

func day1Puzzle1() string {
	data := helpers.LoadData("day-1/puzzle_data.txt")
	result := day_1.Puzzle1(data)
	return fmt.Sprintf("%d", result)
}

func day1Puzzle2() string {
	data := helpers.LoadData("day-1/puzzle_data.txt")
	result := day_1.Puzzle2(data)
	return fmt.Sprintf("%d", result)
}
