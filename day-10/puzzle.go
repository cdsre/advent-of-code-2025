package day_10

import "fmt"

func Puzzle1(data []string) int {
	return puzzle(data, "lights")
}

func Puzzle2(data []string) int {
	return puzzle(data, "joltages")
}

func puzzle(data []string, mode string) int {
	machines := parsePuzzleData(data)
	totalPresses := 0
	for i, machine := range machines {
		fmt.Printf("machine: %d\n", i)
		buttonsPressed := machine.EfficentStart(mode)
		totalPresses += buttonsPressed
	}
	return totalPresses
}
