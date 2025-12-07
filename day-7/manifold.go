package day_7

import (
	"fmt"
	"slices"
	"strings"
)

type Manifold struct {
	beamStatus     map[int]bool
	splitCounter   int
	manifoldMatrix [][]string
}

func NewManifold(data []string) *Manifold {
	gridWidth := len(strings.Split(data[0], ""))
	beamStatus := make(map[int]bool, gridWidth)
	manifoldMatrix := make([][]string, len(data))

	for i, row := range data {
		manifoldMatrix[i] = strings.Split(row, "")
	}

	for i := range manifoldMatrix[0] {
		beamStatus[i] = false
	}
	return &Manifold{beamStatus, 0, manifoldMatrix}
}

func (m *Manifold) ActivateBeam() {
	start := slices.Index(m.manifoldMatrix[0], "S")
	m.beamStatus[start] = true
	m.manifoldMatrix[1][start] = "|"
	for i := 1; i < len(m.manifoldMatrix); i++ {
		for j := range m.manifoldMatrix[i] {
			if m.beamStatus[j] {
				if m.manifoldMatrix[i][j] == "^" {
					m.splitBeam(i, j)
					continue
				}
				if m.manifoldMatrix[i-1][j] == "|" {
					m.manifoldMatrix[i][j] = "|"
				}
			}
		}
	}
}

func (m *Manifold) splitBeam(row int, beamIndex int) {
	m.beamStatus[beamIndex] = false
	left := beamIndex - 1
	right := beamIndex + 1

	if left >= 0 {
		m.beamStatus[left] = true
		m.manifoldMatrix[row][left] = "|"
	}

	if right < len(m.beamStatus) {
		m.beamStatus[right] = true
		m.manifoldMatrix[row][right] = "|"
	}

	m.splitCounter++
}

func (m *Manifold) Display() {
	strb := strings.Builder{}
	for _, row := range m.manifoldMatrix {
		strb.WriteString(strings.Join(row, ""))
		strb.WriteString("\n")
	}
	println(strb.String())
}

func (m *Manifold) QuantumTimeLineCount(i, j int, path []string) int {
	if i >= len(m.manifoldMatrix) {
		fmt.Printf("%v\n", strings.Join(path, "#"))
		path = path[:len(path)-1]
		return 1
	}

	if j >= len(m.manifoldMatrix[i]) {
		return 0
	}

	switch m.manifoldMatrix[i][j] {
	case "S", "|":

		path = append(path, fmt.Sprintf("%d,%d", i, j))
		return m.QuantumTimeLineCount(i+1, j, path)
	case "^":

		path = append(path, fmt.Sprintf("%d,%d", i, j))
		left := m.QuantumTimeLineCount(i+1, j-1, path)
		right := m.QuantumTimeLineCount(i+1, j+1, path)
		return left + right
	default:
		return m.QuantumTimeLineCount(i, j+1, path)
	}
}
