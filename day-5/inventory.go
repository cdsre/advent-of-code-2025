package day_5

import (
	"slices"
	"strconv"
	"strings"
)

type Inventory struct {
	freshRanges [][2]int
}

func NewInventory(freshRanges []string) *Inventory {
	ranges := make([][2]int, len(freshRanges))
	for i, freshRange := range freshRanges {
		rangeSplit := strings.Split(freshRange, "-")
		first, _ := strconv.Atoi(rangeSplit[0])
		last, _ := strconv.Atoi(rangeSplit[1])
		ranges[i] = [2]int{first, last}
	}
	return &Inventory{ranges}
}

func (inv *Inventory) ingredientIsFreshAny(ingredient int) bool {
	for _, freshRange := range inv.freshRanges {
		if ingredient >= freshRange[0] && ingredient <= freshRange[1] {
			return true
		}
	}
	return false
}

func (inv *Inventory) totalFreshIds() int {
	slices.SortFunc(inv.freshRanges, func(a, b [2]int) int {
		if a[0] != b[0] {
			// sort by range start
			return a[0] - b[0]
		}
		// If start is the same, sort by range end
		return a[1] - b[1]
	})

	totalIds := 0
	first := 0
	last := 0

	for _, freshRange := range inv.freshRanges {
		if freshRange[0] <= last {
			first = last + 1
		}

		if freshRange[0] > last {
			first = freshRange[0]
		}

		if freshRange[1] > last {
			last = freshRange[1]
		}

		totalIds += last - first + 1
	}
	return totalIds
}
