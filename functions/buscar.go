package functions

func (m *Matriz) buscarAbajoCompleto() (bool, int, int, int) {
	verify, i, j := m.buscar0Abajo()
	if !verify {
		return false, 0, 0, 0
	}
	hayProximoMultiplicar, b := m.buscarProxMulti(i, j)

	if !hayProximoMultiplicar {
		return false, 0, 0, 0
	}
	return true, i, j, b
}
func (m *Matriz) buscar0Abajo() (bool, int, int) {
	for i := 0; i < len(m.Board); i++ {
		for j := 0; j < len(m.Board[0]); j++ {
			if m.Board[i][j] != 0 && j < i {

				return true, i, j
			}
		}
	}
	return false, 0, 0
}
func (m *Matriz) buscarProxMulti(i, j int) (bool, int) {
	for b := 0; b < len(m.Board); b++ {
		if m.Board[b][j] != 0 && b != i {
			var valorDeJsAnterioresCompatibles float64
			j2 := j
			for ; j2 > 0; j2-- {
				valorDeJsAnterioresCompatibles += m.Board[b][j2-1]
			}
			if valorDeJsAnterioresCompatibles == 0 {
				return true, b
			}
		}
	}
	return false, 0
}
