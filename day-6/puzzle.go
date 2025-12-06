package day_6

import (
	"slices"
	"strconv"
	"strings"
)

func Puzzle1(data []string) int {
	// clone the data before reversing as reversing the slice is done in place and will cause issues on benchmarking
	data = slices.Clone(data)
	slices.Reverse(data)
	return Puzzle(data)
}

func Puzzle2(data []string) int {
	return processCephFormatData(data)
}

func Puzzle(data []string) int {
	calcs := make([]Calculator, len(strings.Fields(data[0])))
	for _, line := range data {
		for i, value := range strings.Fields(line) {
			switch value {
			case "+", "*":
				calcs[i].Operator(value)
			default:
				num, _ := strconv.Atoi(value)
				calcs[i].ApplyOperation(num)
			}
		}
	}

	total := 0
	for _, calc := range calcs {
		total += calc.Result()
	}
	return total
}

func reverseString(data string) string {
	strb := strings.Builder{}
	for i := len(data) - 1; i >= 0; i-- {
		strb.WriteString(string(data[i]))
	}
	return strb.String()
}

func processCephFormatData(data []string) int {
	charGrid := make([][]string, len(data)-1)
	nums := []string{}
	totalNums := []int{}

	operators := strings.Fields(reverseString(data[len(data)-1]))
	operatorIndex := 0

	// Dont iterate over the operators in the last line
	for i, line := range data[:len(data)-1] {
		charGrid[i] = strings.Split(line, "")
	}

	for j := len(charGrid[0]) - 1; j >= 0; j-- {
		strb := strings.Builder{}
		for i := range charGrid {
			strb.WriteString(charGrid[i][j])
		}

		numString := strb.String()
		if strings.Count(numString, " ") == len(numString) {
			total := 0
			switch operators[operatorIndex] {
			case "+":
				for _, num := range nums {
					numS := strings.Replace(num, " ", "", -1)
					numI, _ := strconv.Atoi(numS)
					total += numI
				}
			case "*":
				// we cant multiply starting with 0 so initialise total as 1.
				total = 1
				for _, num := range nums {
					numS := strings.Replace(num, " ", "", -1)
					numI, _ := strconv.Atoi(numS)
					total *= numI
				}
			}

			totalNums = append(totalNums, total)
			nums = []string{}
			operatorIndex++
			continue
		}

		nums = append(nums, numString)

	}

	total := 0
	switch operators[operatorIndex] {
	case "+":
		for _, num := range nums {
			numS := strings.Replace(num, " ", "", -1)
			numI, _ := strconv.Atoi(numS)
			total += numI
		}
	case "*":
		// we cant multiply starting with 0 so initialise total as 1.
		total = 1
		for _, num := range nums {
			numS := strings.Replace(num, " ", "", -1)
			numI, _ := strconv.Atoi(numS)
			total *= numI
		}
	}

	totalNums = append(totalNums, total)
	nums = []string{}
	operatorIndex++

	finalTotal := 0
	for _, num := range totalNums {
		finalTotal += num
	}
	return finalTotal
}
