package day_8

import (
	"sort"
)

func Puzzle1(data []string, pairs int) int {
	junctGaps, junctCircuits := puzzle(data)

	for i := 0; i < pairs; i++ {
		junctCircuits = connectJunctions(junctGaps[i], junctCircuits)
	}

	total := 1
	sort.Slice(junctCircuits, func(i, j int) bool { return len(junctCircuits[i].junctions) > len(junctCircuits[j].junctions) })
	for _, jc := range junctCircuits[:3] {
		total *= len(jc.junctions)
	}

	return total
}

func Puzzle2(data []string, circuitSize int) int {
	junctGaps, junctCircuits := puzzle(data)

	for _, junctGap := range junctGaps {
		junctCircuits = connectJunctions(junctGap, junctCircuits)
		if getTotalJunctions(junctCircuits) == circuitSize {
			return junctGap.j1.x * junctGap.j2.x
		}
	}

	return -1
}

func puzzle(data []string) ([]JunctGap, []*JunctCircuit) {
	junctions := parseJunctions(data)
	junctGaps := getJuctGaps(junctions)
	sortJunctGaps(junctGaps)
	return junctGaps, make([]*JunctCircuit, 0, len(junctions)/2)
}
