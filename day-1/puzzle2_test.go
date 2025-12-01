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
		expectedRemainingSteps string
	}{
		{"L10", 9, 1, "L0"},
		{"R15", 9, 1, "R5"},
		{"R75", 9, 7, "R5"},
	}

	for _, tc := range testCases {
		t.Run(tc.rotation, func(t *testing.T) {
			t.Parallel()
			safe := Safe{CurrentPosition: 50, minPosition: 0, maxPosition: tc.maxPosition}
			rotations, remainingSteps := safe.ReduceRotations(tc.rotation)
			assert.Equal(t, tc.expectedRotations, rotations)
			assert.Equal(t, tc.expectedRemainingSteps, remainingSteps)
		})
	}
}
