package functions

func (m *Matrix) MGauss() (string, []float64) {
	for {
		verify, i, j, b := m.lookForCerosInLower()
		if verify {
			m.calculateTotalReduction(i, j, b)
		} else {
			break
		}

	}
	if !m.checkForUndefined() {
		return "Incompatible", []float64{}
	}
	unknowns := m.calcularXYZ()

	return "Compatible", unknowns
}
