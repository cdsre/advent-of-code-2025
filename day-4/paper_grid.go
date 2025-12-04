package day_4

import (
	"strings"
)

func NewPaperGrid(data []string) [][]string {
	pg := make([][]string, len(data))
	for i, row := range data {
		pg[i] = strings.Split(row, "")
	}
	return pg
}

func addNeighbour(grid [][]string, neighbours *[]string, x int, y int) {
	if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
		*neighbours = append(*neighbours, grid[x][y])
	}
}

func GetGridNeighbours(grid [][]string, x int, y int) []string {
	var neighbours []string
	neighbourCoords := [][]int{
		{x - 1, y - 1}, {x - 1, y}, {x - 1, y + 1},
		{x, y - 1}, {x, y + 1},
		{x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1},
	}
	for _, neighbour := range neighbourCoords {
		addNeighbour(grid, &neighbours, neighbour[0], neighbour[1])
	}

	return neighbours
}
