package day_2

import (
	"strconv"
	"strings"
)

type Product struct {
	ID    int
	valid bool
}

func NewProduct(id string) Product {
	pID, _ := strconv.Atoi(id)

	// Id's should not contain duplicate sequences
	valid := false
	idMidPoint := len(id) / 2
	if id[:idMidPoint] != id[idMidPoint:] {
		valid = true
	}
	return Product{pID, valid}
}

type ProductRange struct {
	products []*Product
	First    int
	Last     int
}

func NewProductRange(first int, last int) ProductRange {
	var products []*Product
	for i := first; i <= last; i++ {
		product := NewProduct(strconv.Itoa(i))
		products = append(products, &product)
	}
	return ProductRange{products, first, last}
}

func ParseProductRange(rangeString string) (int, int) {
	rangeSlice := strings.Split(rangeString, "-")
	first, _ := strconv.Atoi(rangeSlice[0])
	last, _ := strconv.Atoi(rangeSlice[1])
	return first, last
}

func ParseProductRanges(rangesString string) [][]int {
	ranges := strings.Split(rangesString, ",")
	productRanges := make([][]int, len(ranges))

	for i, rangeString := range ranges {
		first, last := ParseProductRange(rangeString)
		productRanges[i] = []int{first, last}
	}
	return productRanges
}
