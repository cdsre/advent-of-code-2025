package day_4

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPaperGrid(t *testing.T) {
	testCases := []struct {
		input    []string
		expected [][]string
	}{
		{
			input:    []string{".@", "@."},
			expected: [][]string{{".", "@"}, {"@", "."}},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			pg := NewPaperGrid(tc.input)
			assert.Equal(t, tc.expected, pg)
		})
	}
}

func TestGetGridNeighbours(t *testing.T) {
	testCases := []struct {
		x        int
		y        int
		expected []string
	}{
		{x: 0, y: 0, expected: []string{"@", ".", "@"}},
		{x: 0, y: 1, expected: []string{".", ".", ".", "@", "."}},
		{x: 0, y: 2, expected: []string{"@", "@", "."}},
		{x: 1, y: 0, expected: []string{".", "@", "@", ".", "@"}},
		{x: 1, y: 1, expected: []string{".", "@", ".", ".", ".", ".", "@", "."}},
		{x: 1, y: 2, expected: []string{"@", ".", "@", "@", "."}},
		{x: 2, y: 0, expected: []string{".", "@", "@"}},
		{x: 2, y: 1, expected: []string{".", "@", ".", ".", "."}},
		{x: 2, y: 2, expected: []string{"@", ".", "@"}},
	}

	grid := [][]string{
		{".", "@", "."},
		{".", "@", "."},
		{".", "@", "."},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d-%d", tc.x, tc.y), func(t *testing.T) {
			result := GetGridNeighbours(grid, tc.x, tc.y)
			assert.Equal(t, tc.expected, result)
		})
	}
}
