package day_2

import (
	"strconv"
	"strings"
)

func Puzzle1(data string) int {
	duplicateSequenceValidator := func() ProductValidator {
		return func(p *Product) {
			idString := strconv.Itoa(p.ID)
			idMidPoint := len(idString) / 2
			if idString[:idMidPoint] == idString[idMidPoint:] {
				p.valid = false
			}
		}
	}
	return puzzle(data, duplicateSequenceValidator())
}

func Puzzle2(data string) int {
	repeatedSequenceValidator := func() ProductValidator {
		return func(p *Product) {
			idString := strconv.Itoa(p.ID)
			idLen := len(idString)
			for i := 1; i <= idLen/2; i++ {
				seq := idString[0:i]
				seqLen := len(seq)
				repeats := idLen / seqLen
				if idString == strings.Repeat(seq, repeats) {
					p.valid = false
					break
				}
			}
		}
	}
	return puzzle(data, repeatedSequenceValidator())
}

func puzzle(data string, validators ...ProductValidator) int {
	total := 0
	productRanges := ParseProductRanges(data)
	for _, productRange := range productRanges {
		products := NewProductRange(productRange[0], productRange[1], validators...).products
		for _, product := range products {
			if !product.valid {
				total += product.ID
			}
		}
	}

	return total
}
