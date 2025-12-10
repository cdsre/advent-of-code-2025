package day_10

import (
	"testing"

	"github.com/cdsre/advent-of-code-2025/helpers"
	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	data := helpers.LoadData("puzzle_test_data.txt")
	result := Puzzle1(data)
	assert.Equal(t, 7, result)
}

func TestPuzzle1Real(t *testing.T) {
	data := helpers.LoadData("puzzle_data.txt")
	result := Puzzle1(data)
	assert.Equal(t, 447, result)
}
