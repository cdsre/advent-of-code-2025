package day_6

import "fmt"

type Calculator struct {
	result    int
	operation string
	operated  bool
}

func NewCalculator() Calculator {
	return Calculator{}
}

func (c *Calculator) Add(nums ...int) int {
	for _, num := range nums {
		c.result += num
	}
	return c.result
}

func (c *Calculator) Multiply(nums ...int) int {
	if !c.operated {
		c.operated = true
		c.result = 1
	}
	for _, num := range nums {
		c.result *= num
	}
	return c.result
}

func (c *Calculator) Operator(operation string) {
	c.operation = operation
}

func (c *Calculator) ApplyOperation(nums ...int) (int, error) {
	switch c.operation {
	case "+":
		return c.Add(nums...), nil
	case "*":
		return c.Multiply(nums...), nil

	default:
		return 0, fmt.Errorf("unknown operation: %s", c.operation)
	}
}

func (c *Calculator) Result() int {
	return c.result
}
