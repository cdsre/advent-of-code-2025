package day_7

func Puzzle1(data []string) int {
	manifold := puzzle(data)
	return manifold.splitCounter
}

func Puzzle2(data []string) int {
	manifold := puzzle(data)
	manifold.Display()
	path := make([]string, 0)
	return manifold.QuantumTimeLineCount(0, 0, path)
}

func puzzle(data []string) *Manifold {
	manifold := NewManifold(data)
	manifold.ActivateBeam()
	return manifold
}
