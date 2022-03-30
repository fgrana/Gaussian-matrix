package functions

import (
	"testing"
)

func TestCalculateUnknowns(t *testing.T) {
	board := [][]float64{
		{1, 1, -1, 1},
		{0, -1, 4, -2},
		{0, 0, 1, 1}}
	matrix := Matrix{board}
	x := matrix.calculateUnknowns()
	if x[2] != -4 || x[1] != 6 || x[0] != 1 {
		t.Error("x should be -4, y should be 6.z should be 1.Values are:", x[2], x[1], x[0])
	}
}
func TestCalculateTotalReduction1(t *testing.T) {
	board := [][]float64{
		{1, 1, -1, 1},
		{0, -1, 4, -2},
		{0, -2, 9, -3}}
	matrix := Matrix{board}
	matrix.calculateTotalReduction(2, 1, 0)
	board = [][]float64{
		{1, 1, -1, 1},
		{0, -1, 4, -2},
		{0, 0, 1, 1}}
	matrix1 := Matrix{board}
	for j := 0; j < len(matrix.Board); j++ {
		if matrix.Board[1][j] != matrix1.Board[1][j] {
			t.Error("Matrix should be equal", matrix.Board, matrix1.Board)
		}
	}
}
func TestCalculateTotalReduction2(t *testing.T) {
	board := [][]float64{
		{1, 1, -1, 1},
		{1, -8, 4, -2},
		{1, -1, 9, -3}}
	matrix := Matrix{board}
	matrix.calculateTotalReduction(2, 0, 1)
	board = [][]float64{
		{1, 1, -1, 1},
		{1, -8, 4, -2},
		{0, 7, 5, -1}}
	matrix1 := Matrix{board}
	for j := 0; j < len(matrix.Board); j++ {
		if matrix.Board[2][j] != matrix1.Board[2][j] {
			t.Error("Matrix should be equal", matrix.Board, matrix1.Board, matrix.Board[1][j], matrix1.Board[1][j], j)
		}
	}
}
