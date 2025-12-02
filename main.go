package main

import (
	"fmt"
	"os"

	"github.com/cdsre/advent-of-code-2025/day-1"
	"github.com/cdsre/advent-of-code-2025/day-2"
	"github.com/cdsre/advent-of-code-2025/helpers"
)

func main() {
	switch os.Args[1] {
	case "day1puzzle1":
		day1Puzzle1()
	case "day1puzzle2":
		day1Puzzle2()
	case "day2puzzle1":
		day2Puzzle1()
	}
}

func day1Puzzle1() {
	data := helpers.LoadData("day-1/puzzle_data.txt")
	result := day_1.Puzzle1(data)
	fmt.Printf("%d\n", result)
}

func day1Puzzle2() {
	data := helpers.LoadData("day-1/puzzle_data.txt")
	result := day_1.Puzzle2(data)
	fmt.Printf("%d\n", result)
}

func day2Puzzle1() {
	data := helpers.LoadData("day-2/puzzle_data.txt")
	result := day_2.Puzzle1(data[0])
	fmt.Printf("%d\n", result)
}
