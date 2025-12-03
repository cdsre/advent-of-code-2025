package day_3

import (
	"slices"
	"strconv"
)

type BatteryBank struct {
	batteries []int
}

func NewBatteryBank(data string) BatteryBank {
	batteries := make([]int, len(data))
	for i, battery := range data {
		joltage, _ := strconv.Atoi(string(battery))
		batteries[i] = joltage
	}
	return BatteryBank{batteries}
}

func (b BatteryBank) getMaxBatteries(numBatteries int) []int {
	slots := make([]int, numBatteries)
	batteriesIndex := 0
	batteriesLen := len(b.batteries)
	for i := range slots {
		maxBattRemaining := numBatteries - i
		eligibleBatteries := b.batteries[batteriesIndex : batteriesLen-maxBattRemaining+1]
		maxBatt := slices.Max(eligibleBatteries)
		maxBattIndex := slices.Index(eligibleBatteries, maxBatt)
		slots[i] = maxBatt
		batteriesIndex += maxBattIndex + 1
	}
	return slots
}

func (b BatteryBank) maxJoltage(numBatteries int) int {
	maxBatteries := b.getMaxBatteries(numBatteries)
	maxJoltage := ""
	for _, maxBattery := range maxBatteries {
		joltage := strconv.Itoa(maxBattery)
		maxJoltage += joltage
	}
	intJoltage, _ := strconv.Atoi(maxJoltage)
	return intJoltage
}
