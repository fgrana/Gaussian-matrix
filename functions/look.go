package functions

func (m *Matrix) lookForNextReduction() (bool, int, int, int) {
	verify, i, j := m.lookForCerosInLower()
	if !verify {
		return false, 0, 0, 0
	}
	isNextMultiplier, b := m.lookForNextMultiplier(i, j)

	if !isNextMultiplier {
		return false, 0, 0, 0
	}
	return true, i, j, b
}

func (m *Matrix) lookForCerosInLower() (bool, int, int) {
	for i := 0; i < len(m.Board); i++ {
		for j := 0; j < len(m.Board[0]); j++ {
			if m.Board[i][j] != 0 && j < i {

				return true, i, j
			}
		}
	}
	return false, 0, 0
}
func (m *Matrix) lookForNextMultiplier(i, j int) (bool, int) {
	for b := 0; b < len(m.Board); b++ {
		if m.Board[b][j] != 0 && b != i {
			var valueOfPreviusJ float64
			j2 := j
			for ; j2 > 0; j2-- {
				valueOfPreviusJ += m.Board[b][j2-1]
			}
			if valueOfPreviusJ == 0 {
				return true, b
			}
		}
	}
	return false, 0
}
