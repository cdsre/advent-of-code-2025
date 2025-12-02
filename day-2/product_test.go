package day_2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	testCases := []struct {
		id    int
		valid bool
	}{
		{1188511885, true},
		{1188511886, true},
		{1010, true},
		{1011, true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("case_%d", tc.id), func(t *testing.T) {
			t.Parallel()
			product := NewProduct(tc.id)
			assert.Equal(t, tc.valid, product.valid)
		})
	}
}

func TestNewProductValidationsDuplicateSequence(t *testing.T) {
	testCases := []struct {
		id    int
		valid bool
	}{
		{1188511885, false},
		{1188511886, true},
		{1010, false},
		{1011, true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("case_%d", tc.id), func(t *testing.T) {
			t.Parallel()
			product := NewProduct(tc.id, duplicateSequenceValidator{})
			assert.Equal(t, tc.valid, product.valid)
		})
	}
}

func TestNewProductMultipleValidations(t *testing.T) {
	testCases := []struct {
		id    int
		valid bool
	}{
		{1, true},
		{3, false},
		{5, false},
		{7, true},
		{9, false},
		{10, false},
		{15, false},
		{22, true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("case_%d", tc.id), func(t *testing.T) {
			product := NewProduct(tc.id, notDivisableByThree{}, notDivisableByFive{})
			assert.Equal(t, tc.valid, product.valid)
		})
	}

}

func TestNewProductRange(t *testing.T) {
	testCases := []struct {
		first        int
		last         int
		expectedSize int
	}{
		{1, 10, 10},
		{10, 29, 20},
		{123, 456, 334},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("first_%d_last_%d", tc.first, tc.last), func(t *testing.T) {
			t.Parallel()
			productRange := NewProductRange(tc.first, tc.last)
			assert.Equal(t, tc.expectedSize, len(productRange.products))
		})
	}
}

func TestParseProductRange(t *testing.T) {
	testCases := []struct {
		rangeString string
		first       int
		last        int
	}{
		{"1-10", 1, 10},
		{"10-29", 10, 29},
		{"123-456", 123, 456},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("case_%s", tc.rangeString), func(t *testing.T) {
			t.Parallel()
			first, last := ParseProductRange(tc.rangeString)
			assert.Equal(t, tc.first, first)
			assert.Equal(t, tc.last, last)
		})
	}
}

func TestParseProductRanges(t *testing.T) {
	testCases := []struct {
		rangesString string
		expectedSize int
	}{
		{"1-10,12-29", 2},
		{"1-10,12-29,30-45", 3},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("case_%s", tc.rangesString), func(t *testing.T) {
			t.Parallel()
			ranges := ParseProductRanges(tc.rangesString)
			assert.Equal(t, tc.expectedSize, len(ranges))
		})
	}
}

type notDivisableByThree struct{}

func (d notDivisableByThree) validate(p *Product) bool {
	return p.ID%3 != 0
}

type notDivisableByFive struct{}

func (d notDivisableByFive) validate(p *Product) bool {
	return p.ID%5 != 0
}
