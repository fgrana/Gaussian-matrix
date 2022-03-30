package functions

func (m *Matrix) checkForUndefined() bool {
	for i := 0; i < len(m.Board); i++ {
		value := m.chechForUndefinedInLine(i)
		if value == 0 {
			return false
		}
	}
	return true
}

func (m *Matrix) chechForUndefinedInLine(i int) float64 {
	var value float64
	for j := 0; j < len(m.Board[0]); j++ {
		value += m.Board[i][j]
	}
	return value

}
