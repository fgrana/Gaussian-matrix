package functions

import "testing"

func TestCheckForUndefinedInLine(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{0, 1, 5, 8},
		{0, 0, 0, 0}}
	matriz := Matrix{board}
	value := matriz.chechForUndefinedInLine(0)
	if value != 12 {
		t.Error("Error, value should be equal to 12. Value =", value)
	}
	value = matriz.chechForUndefinedInLine(1)
	if value != 14 {
		t.Error("Error, value should be equal to 14. Value =", value)
	}
	value = matriz.chechForUndefinedInLine(2)
	if value != 0 {
		t.Error("Error, value should be equal to 0. Value =", value)
	}

}
func TestCheckForUndefined(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{0, 1, 5, 8},
		{0, 0, 0, 0}}
	matriz := Matrix{board}
	if matriz.checkForUndefined() {
		t.Error("should be true")
	}
	board = [][]float64{
		{1, 3, 4, 4},
		{0, 1, 5, 8},
		{0, 1, 2, 1}}
	matriz = Matrix{board}
	if !matriz.checkForUndefined() {
		t.Error("should be false")
	}
}
