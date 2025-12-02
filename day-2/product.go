package day_2

import (
	"strconv"
	"strings"
)

type Product struct {
	ID    int
	valid bool
}

type ProductValidator interface {
	validate(p *Product) bool
}

func NewProduct(id int, validations ...ProductValidator) Product {
	product := Product{ID: id, valid: true}
	for _, validator := range validations {
		valid := validator.validate(&product)
		if !valid {
			product.valid = false
			break
		}
	}
	return product
}

type ProductRange struct {
	products []*Product
	First    int
	Last     int
}

func NewProductRange(first int, last int, validations ...ProductValidator) ProductRange {
	var products []*Product
	for i := first; i <= last; i++ {
		product := NewProduct(i, validations...)
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
