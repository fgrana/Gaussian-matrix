package functions

import (
	"fmt"
	"testing"
)

func TestCalcularXYZ(t *testing.T) {
	board := [][]float64{
		{1, 1, -1, 1},
		{0, -1, 4, -2},
		{0, 0, 1, 1}}
	matriz := Matriz{board}
	x, y, z := matriz.calcularXYZ()
	if x != -4 && y != 6 && z != 1 {
		fmt.Printf("x should be %v,is %v insted.y should be %v,is %v insted.z should be %v,is %v insted.", -4, x, 6, y, 1, z)
		t.Error("should be")
	}
}
func TestCalcularReduccionTotal1(t *testing.T) {
	board := [][]float64{
		{1, 1, -1, 1},
		{0, -1, 4, -2},
		{0, -2, 9, -3}}
	matriz := Matriz{board}
	matriz.calcularReduccionTotal(2, 1, 0)
	board = [][]float64{
		{1, 1, -1, 1},
		{0, -1, 4, -2},
		{0, 0, 1, 1}}
	matriz1 := Matriz{board}
	for j := 0; j < len(matriz.Board); j++ {
		if matriz.Board[1][j] != matriz1.Board[1][j] {
			t.Error("las matrices deberian ser iguales", matriz.Board, matriz1.Board)
		}
	}
}
func TestCalcularReduccionTotal2(t *testing.T) {
	board := [][]float64{
		{1, 1, -1, 1},
		{1, -8, 4, -2},
		{1, -1, 9, -3}}
	matriz := Matriz{board}
	matriz.calcularReduccionTotal(2, 0, 1)
	board = [][]float64{
		{1, 1, -1, 1},
		{1, -8, 4, -2},
		{0, 7, 5, -1}}
	matriz1 := Matriz{board}
	for j := 0; j < len(matriz.Board); j++ {
		if matriz.Board[2][j] != matriz1.Board[2][j] {
			t.Error("las matrices deberian ser iguales", matriz.Board, matriz1.Board, matriz.Board[1][j], matriz1.Board[1][j], j)
		}
	}
}
