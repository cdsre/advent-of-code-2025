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
	maxBatteries := make([]int, numBatteries)
	for i, battery := range b.batteries {
		for j, maxBattery := range maxBatteries {
			remainingBatteries := len(b.batteries[i:])
			remainingMaxBatteries := len(maxBatteries[j:])

			if remainingBatteries < remainingMaxBatteries {
				continue
			}

			if battery > maxBattery {
				maxBatteries[j] = battery
				slices.Replace(maxBatteries, j+1, len(maxBatteries), 0)
				break
			}
		}
	}
	return maxBatteries
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
