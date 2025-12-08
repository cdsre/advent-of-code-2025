package main

import (
	"fmt"
	"os"

	"github.com/cdsre/advent-of-code-2025/day-1"
	"github.com/cdsre/advent-of-code-2025/day-2"
	"github.com/cdsre/advent-of-code-2025/day-3"
	"github.com/cdsre/advent-of-code-2025/day-4"
	"github.com/cdsre/advent-of-code-2025/day-5"
	"github.com/cdsre/advent-of-code-2025/day-6"
	"github.com/cdsre/advent-of-code-2025/day-7"
	"github.com/cdsre/advent-of-code-2025/day-8"
	"github.com/cdsre/advent-of-code-2025/helpers"
)

func main() {
	switch os.Args[1] {
	case "day1puzzle1":
		data := helpers.LoadData("day-1/puzzle_data.txt")
		result := day_1.Puzzle1(data)
		fmt.Printf("%d\n", result)
	case "day1puzzle2":
		data := helpers.LoadData("day-1/puzzle_data.txt")
		result := day_1.Puzzle2(data)
		fmt.Printf("%d\n", result)
	case "day2puzzle1":
		data := helpers.LoadData("day-2/puzzle_data.txt")
		result := day_2.Puzzle1(data[0])
		fmt.Printf("%d\n", result)
	case "day2puzzle2":
		data := helpers.LoadData("day-2/puzzle_data.txt")
		result := day_2.Puzzle2(data[0])
		fmt.Printf("%d\n", result)
	case "day3puzzle1":
		data := helpers.LoadData("day-3/puzzle_data.txt")
		result := day_3.Puzzle1(data)
		fmt.Printf("%d\n", result)
	case "day3puzzle2":
		data := helpers.LoadData("day-3/puzzle_data.txt")
		result := day_3.Puzzle2(data)
		fmt.Printf("%d\n", result)
	case "day4puzzle1":
		data := helpers.LoadData("day-4/puzzle_data.txt")
		result := day_4.Puzzle1(data)
		fmt.Printf("%d\n", result)
	case "day4puzzle2":
		data := helpers.LoadData("day-4/puzzle_data.txt")
		result := day_4.Puzzle2(data)
		fmt.Printf("%d\n", result)
	case "day5puzzle1":
		data := helpers.LoadData("day-5/puzzle_data.txt")
		result := day_5.Puzzle1(data)
		fmt.Printf("%d\n", result)
	case "day5puzzle2":
		data := helpers.LoadData("day-5/puzzle_data.txt")
		result := day_5.Puzzle2(data)
		fmt.Printf("%d\n", result)
	case "day6puzzle1":
		data := helpers.LoadData("day-6/puzzle_data.txt")
		result := day_6.Puzzle1(data)
		fmt.Printf("%d\n", result)
	case "day6puzzle2":
		data := helpers.LoadData("day-6/puzzle_data.txt")
		result := day_6.Puzzle2(data)
		fmt.Printf("%d\n", result)
	case "day7puzzle1":
		data := helpers.LoadData("day-7/puzzle_data.txt")
		result := day_7.Puzzle1(data)
		fmt.Printf("%d\n", result)
	case "day7puzzle2":
		data := helpers.LoadData("day-7/puzzle_data.txt")
		result := day_7.Puzzle2(data)
		fmt.Printf("%d\n", result)
	case "day8puzzle1":
		data := helpers.LoadData("day-8/puzzle_data.txt")
		result := day_8.Puzzle1(data, 1000)
		fmt.Printf("%d\n", result)
	case "day8puzzle2":
		data := helpers.LoadData("day-8/puzzle_data.txt")
		result := day_8.Puzzle2(data)
		fmt.Printf("%d\n", result)
	}
}
