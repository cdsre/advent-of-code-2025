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
