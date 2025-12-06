package day_6

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCalculator(t *testing.T) {
	c := NewCalculator()
	assert.Equal(t, 0, c.result)
	assert.Empty(t, c.operation)
	assert.False(t, c.operated)
}

func TestCalculatorResult(t *testing.T) {
	testCases := []struct {
		total int
	}{
		{5},
		{10},
		{35},
	}

	for _, tc := range testCases {
		t.Run(strconv.Itoa(tc.total), func(t *testing.T) {
			t.Parallel()
			c := Calculator{result: tc.total}
			assert.Equal(t, tc.total, c.Result())
		})
	}
}

func TestCalculatorAdd(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 3, 5}, 9},
		{[]int{3, 5}, 8},
		{[]int{10, -5}, 5},
	}

	for _, tc := range testCases {
		t.Run(strconv.Itoa(tc.expected), func(t *testing.T) {
			t.Parallel()
			c := NewCalculator()
			c.Add(tc.nums...)
			assert.Equal(t, tc.expected, c.Result())
		})
	}
}

func TestCalculatorMultiply(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 3, 5}, 15},
		{[]int{3, 5}, 15},
		{[]int{10, -5}, -50},
	}
	for _, tc := range testCases {
		t.Run(strconv.Itoa(tc.expected), func(t *testing.T) {
			t.Parallel()
			c := NewCalculator()
			c.Multiply(tc.nums...)
			assert.Equal(t, tc.expected, c.Result())
		})
	}
}

func TestCalculatorOperator(t *testing.T) {
	testCases := []struct {
		op string
	}{
		{"+"},
		{"*"},
	}
	for _, tc := range testCases {
		t.Run(tc.op, func(t *testing.T) {
			t.Parallel()
			c := NewCalculator()
			c.Operator(tc.op)
			assert.Equal(t, tc.op, c.operation)
		})
	}
}

func TestCalculatorApplyOperation(t *testing.T) {
	testCases := []struct {
		op          string
		nums        []int
		expected    int
		shouldError bool
	}{
		{"+", []int{1, 3, 5}, 9, false},
		{"*", []int{3, 5}, 15, false},
		{"/", []int{10, 5}, 2, true},
	}

	for _, tc := range testCases {
		t.Run(strconv.Itoa(tc.expected), func(t *testing.T) {
			t.Parallel()
			c := NewCalculator()
			c.Operator(tc.op)
			_, err := c.ApplyOperation(tc.nums...)

			if tc.shouldError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, c.result)
			}
		})
	}
}
