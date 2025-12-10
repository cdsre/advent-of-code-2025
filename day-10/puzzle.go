package day_10

func Puzzle1(data []string) int {
	machines := parsePuzzleData(data)
	totalPresses := 0
	for _, machine := range machines {
		buttonsPressed := machine.EfficentStart()
		totalPresses += buttonsPressed
	}
	return totalPresses
}
