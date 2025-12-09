package day_9

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mikenye/geom2d/linesegment"
	"github.com/mikenye/geom2d/rectangle"
)

func Puzzle1(data []string) int {
	rectangles := getRectangles(data)
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].area > rectangles[j].area
	})

	return int(rectangles[0].area)
}

func convertMatrixToCartesian(data []string) []string {
	cartesian := make([]string, len(data))
	height := len(data)
	for i, matrixCoord := range data {
		pStrings := strings.Split(matrixCoord, ",")
		x, _ := strconv.Atoi(pStrings[0])
		y, _ := strconv.Atoi(pStrings[1])
		cartesian[i] = fmt.Sprintf("%d,%d", x, height-1-y)
	}
	return cartesian
}

func Puzzle2(data []string) int {
	//cartesian := convertMatrixToCartesian(data)
	points := getPoints(data)
	lines := getLines(points)
	rectangles := getRectangles2(points)

	allowedRectangles := make([]rectangle.Rectangle, 0, len(rectangles))
RectangleLoop:
	for _, r := range rectangles {
		//fmt.Printf("Checking rectangle %v\n", r)
		bottom, right, top, left := r.Edges()
		for _, l := range lines {
			//fmt.Printf("Checking line %v\n", l)
			for _, rl := range []linesegment.LineSegment{bottom, right, top, left} {
				//fmt.Printf("Checking rline %v\n", rl)
				intersection, intersect := rl.IntersectionPoints(l)

				if !intersect {
					//fmt.Printf("No intersection found\n")
					continue
				}

				if rl.Slope() != l.Slope() {
					//fmt.Printf("Slope mismatch %v\n", intersection)
					if intersection[0] == rl.Left() || intersection[0] == rl.Right() {
						//fmt.Printf("Skipping slope mismatch %v\n", intersection)
						continue
					}
				}

				if len(intersection) == 2 {
					continue
				}

				// if we get here, we must have an intersection
				//fmt.Printf("Intersection found %v\n", intersection)
				continue RectangleLoop
			}

			//r.EdgesIter(func(edge linesegment.LineSegment) bool {
			//	if edge.Intersects(l) && (edge.Slope() != l.Slope()) {
			//		intersected = true
			//	}
			//	return intersected
			//})
		}

		allowedRectangles = append(allowedRectangles, r)

	}

	largestArea := float64(0)
	for _, r := range allowedRectangles {
		realArea := (r.Height() + 1) * (r.Width() + 1)
		if realArea > largestArea {
			fmt.Printf("Found largest rectangle %v with area %v\n", r, realArea)
			largestArea = realArea
		}
	}
	return int(largestArea)
}
