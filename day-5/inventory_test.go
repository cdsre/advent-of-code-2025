package day_5

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInventory(t *testing.T) {
	testCases := []struct {
		data     []string
		expected Inventory
	}{
		{
			data: strings.Split("1-3,4-7", ","),
			expected: Inventory{[][2]int{
				{1, 3},
				{4, 7},
			}},
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			t.Parallel()
			inventory := NewInventory(tc.data)
			assert.Equal(t, &tc.expected, inventory)
		})
	}
}

func TestIngredientIsFreshAny(t *testing.T) {
	testCases := []struct {
		id    int
		fresh bool
	}{
		{1, true},
		{3, true},
		{5, false},
		{6, false},
		{7, true},
	}
	inv := Inventory{[][2]int{
		{1, 3},
		{7, 9},
	}}
	for _, tc := range testCases {
		t.Run(strconv.Itoa(tc.id), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.fresh, inv.ingredientIsFreshAny(tc.id))
		})
	}
}

func TestTotalFreshIds(t *testing.T) {
	testCases := []struct {
		data     []string
		expected int
	}{
		{
			data:     strings.Split("1-3,4-7", ","),
			expected: 7,
		},
		{
			data:     strings.Split("1-10,6-11,2-5,21-40", ","),
			expected: 31,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			t.Parallel()
			inv := NewInventory(tc.data)
			assert.Equal(t, tc.expected, inv.totalFreshIds())
		})
	}
}
