package day_1

func Puzzle2(data []string) int {
	numOfZeros := 0
	safe := Safe{CurrentPosition: 50, minPosition: 0, maxPosition: 99}
	for _, rotation := range data {
		rotations, remainingSteps := safe.ReduceRotations(rotation)
		numOfZeros += rotations

		wouldPassZero := safe.WouldPassZero(remainingSteps)
		safe.Rotate(remainingSteps)
		if safe.CurrentPosition == 0 || wouldPassZero {
			numOfZeros++
		}
	}
	return numOfZeros
}
