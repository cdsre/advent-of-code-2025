package day_11

import (
	"slices"
	"strings"

	"github.com/dominikbraun/graph"
)

func Puzzle1(data []string) int {
	g := parsePuzzleData(data)
	paths, _ := graph.AllPathsBetween(g, "you", "out")
	return len(paths)
}

func Puzzle2(data []string) int {
	g := parsePuzzleData(data)
	paths, _ := graph.AllPathsBetween(g, "svr", "out")
	pathCount := 0
	for _, path := range paths {
		if slices.Contains(path, "dac") && slices.Contains(path, "fft") {
			pathCount++
		}
	}
	return pathCount
}

func parsePuzzleData(data []string) graph.Graph[string, string] {
	devices := make(map[string][]string, len(data))
	for _, line := range data {
		colonIndex := strings.Index(line, ":")
		inputNode := line[:colonIndex]
		outputNodes := strings.Split(line[colonIndex+1:], " ")
		devices[inputNode] = outputNodes
	}

	g := graph.New(graph.StringHash, graph.Directed())

	for inputNode, outputNodes := range devices {
		g.AddVertex(inputNode)
		for _, outputNode := range outputNodes {
			_, err := g.Vertex(outputNode)
			if err != nil {
				g.AddVertex(outputNode)
			}
			g.AddEdge(inputNode, outputNode)
		}
	}
	return g
}
