package day_11

import (
	"testing"

	"github.com/cdsre/advent-of-code-2025/helpers"
	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	data := helpers.LoadData("puzzle_test_data.txt")
	result := Puzzle1(data)
	assert.Equal(t, 5, result)
}

func TestPuzzle2(t *testing.T) {
	data := helpers.LoadData("puzzle_test_data2.txt")
	result := Puzzle2(data)
	assert.Equal(t, 2, result)
}
