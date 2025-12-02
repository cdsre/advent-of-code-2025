package day_2

import (
	"testing"

	"github.com/cdsre/advent-of-code-2025/helpers"
	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	data := helpers.LoadData("puzzle_test_data.txt")
	result := Puzzle1(data[0])
	assert.Equal(t, 1227775554, result)
}
