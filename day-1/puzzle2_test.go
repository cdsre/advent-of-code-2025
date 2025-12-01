package day_1

import (
	"testing"

	"github.com/cdsre/advent-of-code-2025/helpers"
	"github.com/stretchr/testify/assert"
)

func TestPuzzle2(t *testing.T) {
	data := helpers.LoadData("puzzle_test_data.txt")
	result := Puzzle2(data)
	assert.Equal(t, 6, result)
}

func TestCountRotationsPassedZero(t *testing.T) {
	testCases := []struct {
		rotation               string
		maxPosition            int
		expectedRotations      int
		expectedRemainingSteps int
	}{
		{"L10", 9, 1, 0},
		{"R15", 9, 1, 5},
		{"R75", 9, 7, 5},
	}

	for _, tc := range testCases {
		t.Run(tc.rotation, func(t *testing.T) {
			t.Parallel()
			rotations, remainingSteps := reduceRotations(tc.rotation, tc.maxPosition)
			assert.Equal(t, tc.expectedRotations, rotations)
			assert.Equal(t, tc.expectedRemainingSteps, remainingSteps)
		})
	}
}
