package day_5

import (
	"slices"
	"strconv"
)

func Puzzle1(data []string) int {
	freshCount := 0
	ingredientIDs := data[slices.Index(data, "")+1:]
	inv := Puzzle(data)

	for _, ingredient := range ingredientIDs {
		ingredientId, _ := strconv.Atoi(ingredient)
		if inv.ingredientIsFreshAny(ingredientId) {
			freshCount++
		}
	}
	return freshCount
}

func Puzzle2(data []string) int {
	inv := Puzzle(data)
	return inv.totalFreshIds()
}

func Puzzle(data []string) *Inventory {
	dataSplitLine := slices.Index(data, "")
	freshIDRanges := data[:dataSplitLine]
	return NewInventory(freshIDRanges)
}
