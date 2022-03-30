package functions

func (m *Matrix) calculateUnknowns() []float64 {
	var Unknowns []float64

	var h int = len(m.Board[0]) - 1
	i := h - 1

	for i != -1 {
		var knowns float64
		var count int
		for j := h - 1; j > i; j-- {
			knowns += m.Board[i][j] * Unknowns[count]
			count++
		}
		Unknown := (m.Board[i][h] - knowns) / m.Board[i][i]
		Unknowns = append(Unknowns, Unknown)
		i--
	}
	return Unknowns

}
func (m *Matrix) calculateTotalReduction(i, j, b int) {
	valuePrev := m.Board[b][j]
	valueActual := m.Board[i][j]
	for j2 := 0; j2 < len(m.Board[0]); j2++ {
		m.Board[i][j2] = m.Board[i][j2]*valuePrev - m.Board[b][j2]*valueActual
	}
}
