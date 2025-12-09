package day_9

import (
	"math"
	"strconv"
	"strings"
)

func getRectangles(data []string) []Rectangle {
	rectangles := make([]Rectangle, 0, len(data))

	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			p1 := NewPoint(data[i])
			p2 := NewPoint(data[j])
			rectangles = append(rectangles, NewRectangle(p1, p2))
		}
	}
	return rectangles
}

type Point struct {
	x, y int
}

func NewPoint(pointString string) Point {
	pStrings := strings.Split(pointString, ",")
	x, _ := strconv.Atoi(pStrings[0])
	y, _ := strconv.Atoi(pStrings[1])
	return Point{x, y}
}

type Rectangle struct {
	p1, p2              Point
	width, height, area float64
}

func NewRectangle(p1, p2 Point) Rectangle {
	width := math.Abs(float64(p2.x-p1.x)) + 1
	height := math.Abs(float64(p2.y-p1.y)) + 1
	return Rectangle{p1, p2, width, height, width * height}
}
