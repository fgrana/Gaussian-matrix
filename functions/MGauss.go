package functions

func (m *Matriz) MGauss() (string, []float64) {
	for {
		verify, i, j, b := m.buscarAbajoCompleto()
		if verify {
			m.calcularReduccionTotal(i, j, b)
		} else {
			break
		}

	}
	if !m.checkForUndefined() {
		return "Incompatible", []float64{}
	}
	incognitas := m.calcularXYZ()

	return "Compatible", incognitas
}
