package day_7

func Puzzle1(data []string) int {
	manifold := puzzle(data)
	return manifold.splitCounter
}

func Puzzle2(data []string) int {
	manifold := puzzle(data)
	return manifold.QuantumTimeLineCount(0, 0)
}

func puzzle(data []string) *Manifold {
	manifold := NewManifold(data)
	manifold.ActivateBeam()
	return manifold
}
