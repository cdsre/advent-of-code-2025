package day_1

import (
	"fmt"
)

func reduceRotations(rotation string, maxPositions int) (int, string) {
	direction, steps := ParseRotation(rotation)

	wheelSize := maxPositions + 1 // Wheel starts at 0
	rotations := steps / wheelSize
	remainingSteps := steps % wheelSize
	return rotations, fmt.Sprintf("%s%d", direction, remainingSteps)
}

func Puzzle2(data []string) int {
	numOfZeros := 0
	safe := Safe{CurrentPosition: 50, minPosition: 0, maxPosition: 99}
	for _, rotation := range data {
		rotations, remainingSteps := reduceRotations(rotation, safe.maxPosition)
		numOfZeros += rotations

		wouldPassZero := safe.WouldPassZero(remainingSteps)
		safe.Rotate(remainingSteps)
		if safe.CurrentPosition == 0 || wouldPassZero {
			numOfZeros++
		}
	}
	return numOfZeros
}
