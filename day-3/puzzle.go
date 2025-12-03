package day_3

func puzzle(data []string, numBatteries int) int {
	totalJoltage := 0
	for _, line := range data {
		batteryBank := NewBatteryBank(line)
		totalJoltage += batteryBank.maxJoltage(numBatteries)
	}
	return totalJoltage
}

func Puzzle1(data []string) int {
	return puzzle(data, 2)
}

func Puzzle2(data []string) int {
	return puzzle(data, 12)
}
