package day_4

func Puzzle1(data []string) int {
	return Puzzle(data, 4)
}

func Puzzle2(data []string) int {
	return Puzzle(data, 4, true)
}

func Puzzle(data []string, numNeighbours int, mutate ...bool) int {
	totalRollsRemoved := 0
	pg := NewPaperGrid(data)
	rollsRemoved := RemoveRolls(&pg, numNeighbours, mutate...)
	totalRollsRemoved += rollsRemoved

	if len(mutate) > 0 && mutate[0] {
		for rollsRemoved != 0 {
			rollsRemoved = RemoveRolls(&pg, numNeighbours, mutate...)
			totalRollsRemoved += rollsRemoved
		}
	}
	return totalRollsRemoved
}
