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
	case "D1P1":
		data := helpers.LoadData("day-1/puzzle_data.txt")
		result := day_1.Puzzle1(data)
		fmt.Printf("%d\n", result)
	case "D1P2":
		data := helpers.LoadData("day-1/puzzle_data.txt")
		result := day_1.Puzzle2(data)
		fmt.Printf("%d\n", result)
	case "D2P1":
		data := helpers.LoadData("day-2/puzzle_data.txt")
		result := day_2.Puzzle1(data[0])
		fmt.Printf("%d\n", result)
	case "D2P2":
		data := helpers.LoadData("day-2/puzzle_data.txt")
		result := day_2.Puzzle2(data[0])
		fmt.Printf("%d\n", result)
	case "D3P1":
		data := helpers.LoadData("day-3/puzzle_data.txt")
		result := day_3.Puzzle1(data)
		fmt.Printf("%d\n", result)
	case "D3P2":
		data := helpers.LoadData("day-3/puzzle_data.txt")
		result := day_3.Puzzle2(data)
		fmt.Printf("%d\n", result)
	case "D4P1":
		data := helpers.LoadData("day-4/puzzle_data.txt")
		result := day_4.Puzzle1(data)
		fmt.Printf("%d\n", result)
	case "D4P2":
		data := helpers.LoadData("day-4/puzzle_data.txt")
		result := day_4.Puzzle2(data)
		fmt.Printf("%d\n", result)
	case "D5P1":
		data := helpers.LoadData("day-5/puzzle_data.txt")
		result := day_5.Puzzle1(data)
		fmt.Printf("%d\n", result)
	case "D5P2":
		data := helpers.LoadData("day-5/puzzle_data.txt")
		result := day_5.Puzzle2(data)
		fmt.Printf("%d\n", result)
	case "D6P1":
		data := helpers.LoadData("day-6/puzzle_data.txt")
		result := day_6.Puzzle1(data)
		fmt.Printf("%d\n", result)
	case "D6P2":
		data := helpers.LoadData("day-6/puzzle_data.txt")
		result := day_6.Puzzle2(data)
		fmt.Printf("%d\n", result)
	case "D7P1":
		data := helpers.LoadData("day-7/puzzle_data.txt")
		result := day_7.Puzzle1(data)
		fmt.Printf("%d\n", result)
	case "D7P2":
		data := helpers.LoadData("day-7/puzzle_data.txt")
		result := day_7.Puzzle2(data)
		fmt.Printf("%d\n", result)
	case "D8P1":
		data := helpers.LoadData("day-8/puzzle_data.txt")
		result := day_8.Puzzle1(data, 1000)
		fmt.Printf("%d\n", result)
	case "D8P2":
		data := helpers.LoadData("day-8/puzzle_data.txt")
		result := day_8.Puzzle2(data, len(data))
		fmt.Printf("%d\n", result)
	}
}
