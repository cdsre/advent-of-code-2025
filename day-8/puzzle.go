package day_8

import (
	"sort"
)

func Puzzle1(data []string, pairs int) int {
	junctGaps := puzzle(data)
	junctCircuits := make([]*JunctCircuit, 0, len(junctGaps)/2)

	for i := 0; i < pairs; i++ {
		junctCircuits = connectJunctions(junctGaps[i], junctCircuits)
	}

	sort.Slice(junctCircuits, func(i, j int) bool { return len(junctCircuits[i].junctions) > len(junctCircuits[j].junctions) })

	total := 1
	for _, jc := range junctCircuits[:3] {
		total *= len(jc.junctions)
	}
	return total
}

func Puzzle2(data []string) int {
	var junc1, junc2 Junction
	circuitSize := len(data)
	junctGaps := puzzle(data)
	junctCircuits := make([]*JunctCircuit, 0, len(junctGaps)/2)

CirrcuitLoop:
	for {
		for _, junctGap := range junctGaps {
			junctCircuits = connectJunctions(junctGap, junctCircuits)

			if len(junctCircuits) == 1 && getTotalJunctions(junctCircuits) == circuitSize {
				junc1 = junctGap.p1
				junc2 = junctGap.p2
				break CirrcuitLoop
			}
		}
	}

	return junc1.x * junc2.x
}

func puzzle(data []string) []JunctGap {
	junctions := parseJunctions(data)
	junctGaps := getJuctGaps(junctions)
	sortJunctGaps(junctGaps)
	return junctGaps
}
