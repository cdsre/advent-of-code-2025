package day_3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBatteryBank(t *testing.T) {
	testCases := []struct {
		data           string
		expectedLength int
	}{
		{"123", 3},
		{"123456789", 9},
		{"847263453123", 12},
	}
	for _, tc := range testCases {
		t.Run(tc.data, func(t *testing.T) {
			t.Parallel()
			bank := NewBatteryBank(tc.data)
			assert.Equal(t, tc.expectedLength, len(bank.batteries))
		})
	}
}

func TestGetMaxBatteries(t *testing.T) {
	testCases := []struct {
		data         string
		numBatteries int
		maxBatteries []int
	}{
		{"987654321111111", 2, []int{9, 8}},
		{"811111111111119", 2, []int{8, 9}},
		{"234234234234278", 2, []int{7, 8}},
		{"818181911112111", 2, []int{9, 2}},
		{"987654321111111", 12, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1}},
		{"811111111111119", 12, []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}},
		{"234234234234278", 12, []int{4, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}},
		{"818181911112111", 12, []int{8, 8, 8, 9, 1, 1, 1, 1, 2, 1, 1, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.data, func(t *testing.T) {
			t.Parallel()
			bank := NewBatteryBank(tc.data)
			assert.Equal(t, tc.maxBatteries, bank.getMaxBatteries(tc.numBatteries))
		})
	}
}

func TestMaxJoltage(t *testing.T) {
	testCases := []struct {
		data         string
		numBatteries int
		expected     int
	}{
		{"987654321111111", 2, 98},
		{"811111111111119", 2, 89},
		{"234234234234278", 2, 78},
		{"818181911112111", 2, 92},
		{"987654321111111", 12, 987654321111},
		{"811111111111119", 12, 811111111119},
		{"234234234234278", 12, 434234234278},
		{"818181911112111", 12, 888911112111},
	}

	for _, tc := range testCases {
		t.Run(tc.data, func(t *testing.T) {
			t.Parallel()
			bank := NewBatteryBank(tc.data)
			assert.Equal(t, tc.expected, bank.maxJoltage(tc.numBatteries))
		})
	}
}
