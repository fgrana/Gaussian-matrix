package functions

func (m *Matrix) calcularXYZ() []float64 {
	var incognitas []float64

	var h int = len(m.Board[0]) - 1
	i := h - 1

	for i != -1 {
		var inDescubiertas float64
		var count int
		for j := h - 1; j > i; j-- {
			inDescubiertas += m.Board[i][j] * incognitas[count]
			count++
		}
		incognita := (m.Board[i][h] - inDescubiertas) / m.Board[i][i]
		incognitas = append(incognitas, incognita)
		i--
	}
	return incognitas

}
func (m *Matrix) calculateTotalReduction(i, j, b int) {
	valorPrev := m.Board[b][j]
	valorActual := m.Board[i][j]
	for j2 := 0; j2 < len(m.Board[0]); j2++ {
		m.Board[i][j2] = m.Board[i][j2]*valorPrev - m.Board[b][j2]*valorActual
	}
}
