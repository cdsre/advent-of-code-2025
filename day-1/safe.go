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
	direction := string(rotation[0])
	stepsStr := rotation[1:]
	steps, _ := strconv.Atoi(stepsStr)

	switch direction {
	case "L":
		s.Left(steps)
	case "R":
		s.Right(steps)
	}
}
