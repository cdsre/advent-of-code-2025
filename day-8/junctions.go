package day_8

import (
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Junction struct {
	x, y, z int
}

type JunctGap struct {
	distance float64
	p1       Junction
	p2       Junction
}

type JunctCircuit struct {
	junctions []Junction
}

func (jc *JunctCircuit) Contains(j Junction) bool {
	return slices.Contains(jc.junctions, j)

}

func (jc *JunctCircuit) Add(j Junction) {
	jc.junctions = append(jc.junctions, j)
}

func getDistance(p1 Junction, p2 Junction) JunctGap {
	xdiff := float64(p1.x - p2.x)
	yDiff := float64(p1.y - p2.y)
	zDiff := float64(p1.z - p2.z)
	junctDist := math.Sqrt(xdiff*xdiff + yDiff*yDiff + zDiff*zDiff)
	return JunctGap{junctDist, p1, p2}
}

func getJuctGaps(junctions []Junction) []JunctGap {
	junctGaps := make([]JunctGap, 0, len(junctions))
	for i := 0; i < len(junctions)-1; i++ {
		for j := i + 1; j < len(junctions); j++ {
			junctGaps = append(junctGaps, getDistance(junctions[i], junctions[j]))
		}
	}
	return junctGaps
}

func sortJunctGaps(junctGaps []JunctGap) {
	sort.Slice(junctGaps, func(i, j int) bool {
		return junctGaps[i].distance < junctGaps[j].distance
	})
}

func getTotalJunctions(junctCircuits []*JunctCircuit) int {
	total := 0
	for _, jc := range junctCircuits {
		total += len(jc.junctions)
	}
	return total
}

func parseJunctions(data []string) []Junction {
	junctions := make([]Junction, len(data))
	for i, value := range data {
		stringCoord := strings.Split(value, ",")
		intCoord := make([]int, len(stringCoord))
		for j, coord := range stringCoord {
			strCoord := strings.TrimSpace(coord)
			intCoord[j], _ = strconv.Atoi(strCoord)
		}
		junctions[i] = Junction{intCoord[0], intCoord[1], intCoord[2]}
	}
	return junctions
}

func connectJunctions(junctGap JunctGap, junctCircuits []*JunctCircuit) []*JunctCircuit {
	var j1Circuit, j2Circuit *JunctCircuit
	j1, j2 := junctGap.p1, junctGap.p2

	for _, jc := range junctCircuits {
		if jc.Contains(j1) {
			j1Circuit = jc
		}
		if jc.Contains(j2) {
			j2Circuit = jc
		}
	}

	if j1Circuit == nil && j2Circuit == nil {
		junctCircuits = append(junctCircuits, &JunctCircuit{junctions: []Junction{j1, j2}})
	} else if j1Circuit != nil && j2Circuit == nil {
		j1Circuit.Add(j2)
	} else if j1Circuit == nil && j2Circuit != nil {
		j2Circuit.Add(j1)
	} else if j1Circuit != j2Circuit {
		for _, j := range j2Circuit.junctions {
			if !j1Circuit.Contains(j) {
				j1Circuit.Add(j)
			}
		}
		j2CircuitIndex := slices.Index(junctCircuits, j2Circuit)
		junctCircuits = append(junctCircuits[:j2CircuitIndex], junctCircuits[j2CircuitIndex+1:]...)
	}
	return junctCircuits
}
