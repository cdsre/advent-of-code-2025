package day_1

import "strconv"

type Safe struct {
	CurrentPosition int
	minPosition     int
	maxPosition     int
}

func (s *Safe) Left(steps int) int {
	remainingPositions := s.CurrentPosition - s.minPosition
	if steps > remainingPositions {
		remainingSteps := steps - remainingPositions - 1
		s.CurrentPosition = s.maxPosition
		return s.Left(remainingSteps)
	}
	newPosition := s.CurrentPosition - steps
	s.CurrentPosition = newPosition
	return s.CurrentPosition
}

func (s *Safe) Right(steps int) int {
	remainingPositions := s.maxPosition - s.CurrentPosition
	if steps > remainingPositions {
		remainingSteps := steps - remainingPositions - 1
		s.CurrentPosition = s.minPosition
		return s.Right(remainingSteps)
	}
	newPosition := s.CurrentPosition + steps
	s.CurrentPosition = newPosition
	return s.CurrentPosition
}

func (s *Safe) Rotate(rotation string) {
	direction, steps := ParseRotation(rotation)

	switch direction {
	case "L":
		s.Left(steps)
	case "R":
		s.Right(steps)
	}
}

func (s *Safe) WouldPassZero(rotation string) bool {
	direction, steps := ParseRotation(rotation)
	var result bool

	if s.CurrentPosition == 0 {
		return false
	}

	switch direction {
	case "L":
		result = s.CurrentPosition-steps < 0
	case "R":
		result = s.CurrentPosition+steps > s.maxPosition
	}
	return result
}

func ParseRotation(rotation string) (direction string, steps int) {
	direction = string(rotation[0])
	steps, _ = strconv.Atoi(rotation[1:])
	return direction, steps
}
