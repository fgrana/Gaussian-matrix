package functions

import "testing"

func TestCheckForUndefinedInLine(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{0, 1, 5, 8},
		{0, 0, 0, 0}}
	matriz := Matriz{board}
	value := matriz.chechForUndefinedInLine(0)
	if value != 8 {
		t.Error("Error, value should be equalto. Value =", value)
	}
	value = matriz.chechForUndefinedInLine(1)
	if value != 6 {
		t.Error("Error, value should be equalto. Value =", value)
	}
	value = matriz.chechForUndefinedInLine(2)
	if value != 0 {
		t.Error("Error, value should be equalto. Value =", value)
	}

}
func TestCheckForUndefined(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{0, 1, 5, 8},
		{0, 0, 0, 0}}
	matriz := Matriz{board}
	if matriz.checkForUndefined() {
		t.Error("should be true")
	}
	board = [][]float64{
		{1, 3, 4, 4},
		{0, 1, 5, 8},
		{0, 1, 2, 1}}
	matriz = Matriz{board}
	if !matriz.checkForUndefined() {
		t.Error("should be false")
	}
}
