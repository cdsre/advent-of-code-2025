package day_2

import (
	"strconv"
	"strings"
)

func Puzzle1(data string) int {
	return puzzle(data, duplicateSequenceValidator{})
}

func Puzzle2(data string) int {
	return puzzle(data, repeatedSequenceValidator{})
}

func puzzle(data string, validators ...ProductValidator) int {
	total := 0
	productRanges := ParseProductRanges(data)
	for _, productRange := range productRanges {
		first, last := productRange[0], productRange[1]
		productRange := NewProductRange(first, last, validators...)
		for _, product := range productRange.products {
			if !product.valid {
				total += product.ID
			}
		}
	}

	return total
}

type duplicateSequenceValidator struct{}

func (d duplicateSequenceValidator) validate(p *Product) bool {
	idString := strconv.Itoa(p.ID)
	idMidPoint := len(idString) / 2
	return idString[:idMidPoint] != idString[idMidPoint:]
}

type repeatedSequenceValidator struct{}

func (r repeatedSequenceValidator) validate(p *Product) bool {
	idString := strconv.Itoa(p.ID)
	idLen := len(idString)
	for i := 1; i <= idLen/2; i++ {
		seq := idString[0:i]
		seqLen := len(seq)
		repeats := idLen / seqLen
		if idString == strings.Repeat(seq, repeats) {
			return false
		}

	}
	return true
}
