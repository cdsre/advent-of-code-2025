package day_4

func Puzzle1(data []string) int {
	return Puzzle(data, 4)
}

func Puzzle2(data []string) int {
	totalRollsRemoved := 0
	rollsRemoved := 1
	for rollsRemoved > 0 {
		rollsRemoved = Puzzle(data, 4, true)
		totalRollsRemoved += rollsRemoved
	}
	return totalRollsRemoved
}

func Puzzle(data []string, numNeighbours int, mutate ...bool) int {
	paperRollsRemoved := 0
	pg := NewPaperGrid(data)
	for i := range pg {
		for j := range pg[i] {
			if pg[i][j] != "@" {
				continue
			}

			neighbours := GetGridNeighbours(pg, i, j)
			neighbourRolls := 0
			for _, neighbour := range neighbours {
				if neighbour == "@" {
					neighbourRolls++
				}
			}

			if neighbourRolls < numNeighbours {
				paperRollsRemoved++
				if len(mutate) > 0 && mutate[0] {
					pg[i][j] = "."
				}
			}
		}
	}
	return paperRollsRemoved
}
