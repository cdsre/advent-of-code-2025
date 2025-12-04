package day_4

var featFlags = make(map[string]any)

func setFeatFlag(flag string, value any) {
	featFlags[flag] = value
}

func getFeatFlag[T any](flag string, fallback T) T {
	if value, ok := featFlags[flag]; ok {
		return value.(T)
	}
	return fallback
}

func Puzzle1(data []string) int {
	return Puzzle(data, 4)
}

func Puzzle2(data []string) int {
	setFeatFlag("mutate", true)
	return Puzzle(data, 4)
}

func Puzzle(data []string, numNeighbours int) int {
	totalRollsRemoved := 0
	pg := NewPaperGrid(data)
	rollsRemoved := RemoveRolls(&pg, numNeighbours)
	totalRollsRemoved += rollsRemoved

	if getFeatFlag("mutate", false) {
		for rollsRemoved != 0 {
			rollsRemoved = RemoveRolls(&pg, numNeighbours)
			totalRollsRemoved += rollsRemoved
		}
	}
	return totalRollsRemoved
}
