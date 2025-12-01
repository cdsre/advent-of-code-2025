package day_1

func Puzzle1(data []string) int {
	numOfZeros := 0
	safe := Safe{CurrentPosition: 50, minPosition: 0, maxPosition: 99}
	for _, rotation := range data {
		safe.Rotate(rotation)
		if safe.CurrentPosition == 0 {
			numOfZeros++
		}
	}
	return numOfZeros
}
