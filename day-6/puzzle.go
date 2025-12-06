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
	// we will create a grid of all the chars to iterate. We will also create a colNums and colTotals list to track
	charGrid := make([][]string, len(data)-1)
	colNums := []int{}
	colTotals := []int{}

	// since we read from right to left we will reverse the operator string and then split it into fields
	reversedOps := reverseString(data[len(data)-1])
	operators := strings.Fields(reversedOps)
	operatorIndex := 0

	// Dont iterate over the operators in the last line just the number data
	for i, line := range data[:len(data)-1] {
		charGrid[i] = strings.Split(line, "")
	}

	//iterate throught he grid from top to bottom and right to left building columns strings
	for j := len(charGrid[0]) - 1; j >= 0; j-- {
		strb := strings.Builder{}
		for i := range charGrid {
			strb.WriteString(charGrid[i][j])
		}
		colString := strb.String()

		// If the colString is all spaces then its a column seperator and we should total the column numbers
		if strings.Count(colString, " ") == len(colString) {
			colTotals = append(colTotals, totalColumn(operators[operatorIndex], colNums))
			colNums = []int{}
			operatorIndex++
			continue
		}

		// If we got here then we have actual numbers, we need to strip out spaces, convert it to an int and add to our column list
		numS := strings.Replace(colString, " ", "", -1)
		numI, _ := strconv.Atoi(numS)
		colNums = append(colNums, numI)

	}

	// This is a classic end of loop need to process the last number
	colTotals = append(colTotals, totalColumn(operators[operatorIndex], colNums))

	// Sum all the column totals that are tracked in colTotals array
	finalTotal := 0
	for _, num := range colTotals {
		finalTotal += num
	}
	return finalTotal
}

func totalColumn(operator string, nums []int) int {
	total := 0
	switch operator {
	case "+":
		for _, num := range nums {
			total += num
		}
	case "*":
		// we cant multiply starting with 0 or the result will always be 0 so initialise total as 1.
		total = 1
		for _, num := range nums {
			total *= num
		}
	}

	return total
}
