package day_1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafeLeft(t *testing.T) {
	testCases := []struct {
		startPos         int
		steps            int
		expectedPosition int
	}{
		{0, 3, 8},
		{5, 3, 2},
		{10, 7, 3},
		{0, 12, 10},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			t.Parallel()
			safe := Safe{CurrentPosition: tc.startPos, minPosition: 0, maxPosition: 10}
			steps := tc.steps % (safe.maxPosition - safe.minPosition + 1) // reduce un-needed rotations
			newPosition := safe.Left(steps)
			assert.Equal(t, tc.expectedPosition, newPosition)
		})
	}
}

func TestSafeRight(t *testing.T) {
	testCases := []struct {
		startPos         int
		steps            int
		expectedPosition int
	}{
		{10, 3, 2},
		{5, 3, 8},
		{2, 7, 9},
		{10, 12, 0},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			t.Parallel()
			safe := Safe{CurrentPosition: tc.startPos, minPosition: 0, maxPosition: 10}
			steps := tc.steps % (safe.maxPosition - safe.minPosition + 1) // reduce un-needed rotations
			newPosition := safe.Right(steps)
			assert.Equal(t, tc.expectedPosition, newPosition)
		})
	}
}

func TestSafeRotate(t *testing.T) {
	testCases := []struct {
		startPos         int
		rotation         string
		expectedPosition int
	}{
		{0, "L3", 8},
		{5, "L3", 2},
		{10, "L7", 3},
		{0, "R12", 1},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			t.Parallel()
			safe := Safe{CurrentPosition: tc.startPos, minPosition: 0, maxPosition: 10}
			safe.Rotate(tc.rotation)
			assert.Equal(t, tc.expectedPosition, safe.CurrentPosition)
		})
	}
}
