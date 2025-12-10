package day_10

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCombinationsWithRepeats(t *testing.T) {
	testCases := []struct {
		n      int
		k      int
		result [][]int
	}{
		{1, 1, [][]int{{0}}},
		{1, 2, [][]int{{0}, {1}}},
		{1, 3, [][]int{{0}, {1}, {2}}},
		{2, 1, [][]int{{0, 0}}},
		{2, 2, [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("case_%d_%d", tc.n, tc.k), func(t *testing.T) {
			t.Parallel()
			result := GetCombinationsWithRepeats(tc.n, tc.k)
			assert.Equal(t, tc.result, result)
		})
	}
}
