package day_10

import (
	"slices"
	"strconv"
	"strings"
)

func GetCombinationsWithRepeats(n, k int) [][]int {
	if n <= 0 || k <= 0 {
		return [][]int{}
	}

	result := make([][]int, 0)
	current := make([]int, n)

	var dfs func(pos int)
	dfs = func(pos int) {
		if pos == n {
			combo := make([]int, n)
			copy(combo, current)
			result = append(result, combo)
			return
		}
		for v := 0; v < k; v++ {
			current[pos] = v
			dfs(pos + 1)
		}
	}

	dfs(0)
	return result
}

type Machine struct {
	lightsCode []bool
	lights     []bool
	buttons    [][]int
	joltages   []int
}

func (m *Machine) Reset() {
	m.lights = make([]bool, len(m.lightsCode))
}

func (m *Machine) PressButton(i int) {
	button := m.buttons[i]
	for _, j := range button {
		m.lights[j] = !m.lights[j]
	}
}

func (m *Machine) LightsMatch() bool {
	return slices.Equal(m.lightsCode, m.lights)
}

func (m *Machine) EfficentStart() int {
	buttonPresses := 1
	for {
		buttonCombos := GetCombinationsWithRepeats(buttonPresses, len(m.buttons))
		for _, combo := range buttonCombos {
			m.Reset()
			for j := range combo {
				m.PressButton(combo[j])
			}
			if m.LightsMatch() {
				return buttonPresses
			}
		}
		buttonPresses++
	}
}

func newMachine(lightsCode []bool, buttons [][]int, joltages []int) Machine {
	return Machine{lightsCode, make([]bool, len(lightsCode)), buttons, joltages}
}

func parseLights(lightsCode []string) []bool {
	lights := make([]bool, len(lightsCode[0])-2)
	for i, v := range strings.Split(lightsCode[0], "")[1 : len(lightsCode[0])-1] {
		if v == "." {
			lights[i] = false
		}

		if v == "#" {
			lights[i] = true
		}
	}
	return lights
}

func parseJoltage(joltage []string) []int {
	joltages := make([]int, 0, len(joltage[0])-2)
	joltagesString := joltage[0][1 : len(joltage[0])-1]
	for _, v := range strings.Split(joltagesString, ",") {
		if v == "," {
			continue
		}
		res, _ := strconv.Atoi(v)
		joltages = append(joltages, res)
	}
	return joltages
}

func parseButtons(buttonsString []string) [][]int {
	buttons := make([][]int, len(buttonsString))
	for i, buttonString := range buttonsString {
		buttonPositions := strings.Split(buttonString[1:len(buttonString)-1], ",")
		buttons[i] = make([]int, len(buttonPositions))
		for j, buttonPosition := range buttonPositions {
			buttons[i][j], _ = strconv.Atoi(buttonPosition)
		}
	}
	return buttons
}

func parsePuzzleData(data []string) []Machine {
	machines := make([]Machine, len(data))
	for i, line := range data {
		parts := strings.Split(line, " ")
		lightsCode := parseLights(parts[:1])
		buttons := parseButtons(parts[1 : len(parts)-1])
		joltages := parseJoltage(parts[len(parts)-1:])
		machine := newMachine(lightsCode, buttons, joltages)
		machines[i] = machine
	}
	return machines
}
