package day_7

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cdsre/advent-of-code-2025/helpers"
	"github.com/stretchr/testify/assert"
)

func TestNewManifold(t *testing.T) {
	testCases := []struct {
		data         []string
		expectedRows int
		expectedCols int
	}{
		{
			strings.Split("..S..\n..^..", "\n"),
			2,
			5,
		},
		{
			strings.Split("..S\n..^\n...\n...", "\n"),
			4,
			3,
		},
		{
			helpers.LoadData("puzzle_test_data.txt"),
			16,
			15,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			manifold := NewManifold(tc.data)
			rowLen := len(manifold.manifoldMatrix)
			colLen := len(manifold.manifoldMatrix[0])
			assert.Equal(t, tc.expectedRows, rowLen)
			assert.Equal(t, tc.expectedCols, colLen)
		})
	}
}
