package day_9

import (
	"strconv"
	"strings"

	"github.com/mikenye/geom2d/linesegment"
	"github.com/mikenye/geom2d/point"
	"github.com/mikenye/geom2d/rectangle"
)

func getRectangles2(points []point.Point) []rectangle.Rectangle {
	rectangles := make([]rectangle.Rectangle, 0, len(points))

	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			rectangles = append(rectangles, rectangle.New(points[i].X(), points[i].Y(), points[j].X(), points[j].Y()))
		}
	}
	return rectangles
}

func getPoints(data []string) []point.Point {
	points := make([]point.Point, 0, len(data))
	for _, p := range data {
		pStrings := strings.Split(p, ",")
		x, _ := strconv.Atoi(pStrings[0])
		y, _ := strconv.Atoi(pStrings[1])
		points = append(points, point.New(float64(x), float64(y)))
	}
	return points
}

func getLines(points []point.Point) []linesegment.LineSegment {
	lines := make([]linesegment.LineSegment, len(points))
	for i := 0; i < len(points)-1; i++ {
		lines[i] = linesegment.NewFromPoints(points[i], points[i+1])
	}
	lines = append(lines, linesegment.NewFromPoints(points[len(points)-1], points[0]))
	return lines
}
