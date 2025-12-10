package day_10

func Puzzle1(data []string) int {
	return puzzle(data, "lights")
}

func Puzzle2(data []string) int {
	return puzzle(data, "joltages")
}

func puzzle(data []string, mode string) int {
	machines := parsePuzzleData(data)
	totalPresses := 0
	for _, machine := range machines {
		buttonsPressed := machine.EfficentStart(mode)
		totalPresses += buttonsPressed
	}
	return totalPresses
}
