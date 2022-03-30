package functions

import (
	"fmt"
	"testing"
)

func TestLookForNextReductionAllOkey(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{0, 1, 5, 8},
		{0, 0, 1, 9}}
	matrix := Matrix{board}
	verify, i, j, b := matrix.lookForNextReduction()
	if verify && i != 0 && j != 0 && b != 0 {
		text, _ := fmt.Printf("verify should be true. i = %v, j = %v, b = %v", i, j, b)
		t.Error(text)
	}
}
func TestLookForNextReduction1th(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{1, 1, 5, 8},
		{0, 0, 1, 9}}
	matrix := Matrix{board}
	verify, i, j, b := matrix.lookForNextReduction()
	if !verify || i != 1 || j != 0 || b != 0 {
		text, _ := fmt.Printf("verify should be true. i = %v, j = %v, b = %v", i, j, b)
		t.Error(text)
	}
}
func TestLookForNextReduction2th(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{0, 1, 5, 8},
		{1, 0, 1, 9}}
	matrix := Matrix{board}
	verify, i, j, b := matrix.lookForNextReduction()
	if !verify || i != 2 || j != 0 || b != 0 {
		text, _ := fmt.Printf("verify should be true. i = %v, j = %v, b = %v", i, j, b)
		t.Error(text)
	}
}
func TestLookForNextReduction3th(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{0, 1, 5, 8},
		{0, 1, 1, 9}}
	matrix := Matrix{board}
	verify, i, j, b := matrix.lookForNextReduction()
	if !verify || i != 2 || j != 1 || b != 1 {
		text, _ := fmt.Printf("verify should be true. i = %v, j = %v, b = %v", i, j, b)
		t.Error(text)
	}
}
func TestLookForNextReduction4th(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{0, 1, 5, 8},
		{1, 1, 1, 9}}
	matrix := Matrix{board}
	verify, i, j, b := matrix.lookForNextReduction()
	if !verify || i != 2 || j != 0 || b != 0 {
		text, _ := fmt.Printf("verify should be true. i = %v, j = %v, b = %v", i, j, b)
		t.Error(text)
	}
}
func TestLookForNextReduction5th(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{1, 1, 5, 8},
		{1, 0, 1, 9}}
	matrix := Matrix{board}
	verify, i, j, b := matrix.lookForNextReduction()
	if !verify || i != 1 || j != 0 || b != 0 {
		text, _ := fmt.Printf("verify should be true. i = %v, j = %v, b = %v", i, j, b)
		t.Error(text)
	}
}
func TestLookForNextReduction6th(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{1, 1, 5, 8},
		{1, 1, 1, 9}}
	matrix := Matrix{board}
	verify, i, j, b := matrix.lookForNextReduction()
	if !verify || i != 1 || j != 0 || b != 0 {
		text, _ := fmt.Printf("verify should be true. i = %v, j = %v, b = %v", i, j, b)
		t.Error(text)
	}
}

func TestLookForCerosInLower(t *testing.T) {
	board := [][]float64{
		{0, 0, 0, 0},
		{1, 0, 0, 0},
		{0, 0, 0, 0}}
	matrix := Matrix{board}
	verify, i, j := matrix.lookForCerosInLower()
	if !verify {
		t.Error("should be true", i, j)
	}
	board = [][]float64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 0, 0, 0}}
	matrix = Matrix{board}
	verify, i, j = matrix.lookForCerosInLower()
	if !verify {
		t.Error("should be true", i, j)
	}
	board = [][]float64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 1, 0, 0}}
	matrix = Matrix{board}
	verify, i, j = matrix.lookForCerosInLower()
	if !verify {
		t.Error("should be true", i, j)
	}
	board = [][]float64{
		{1, 1, 1, 1},
		{0, 1, 1, 1},
		{0, 0, 1, 1}}
	matrix = Matrix{board}
	verify, i, j = matrix.lookForCerosInLower()
	if verify {
		t.Error("should be false", i, j)
	}
}

func TestLookForNextMultiplierThe1thPossibility(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{1, 1, 5, 8},
		{1, 1, 1, 9}}
	matrix := Matrix{board}
	verify, num := matrix.lookForNextMultiplier(1, 0)
	if verify && num != 0 {
		t.Error("should be true. Num should be 0. Num=", num)
	}
}
func TestLookForNextMultiplierThe2thPossibility(t *testing.T) {
	board := [][]float64{
		{0, 3, 4, 4},
		{1, 1, 5, 8},
		{1, 1, 1, 9}}
	matrix := Matrix{board}
	verify, num := matrix.lookForNextMultiplier(1, 0)
	if verify && num != 2 {
		t.Error("should be true. Num should be 1. Num=", num)
	}
}
func TestLookForNextMultiplierThe3thPossibility(t *testing.T) {
	board := [][]float64{
		{0, 3, 4, 4},
		{0, 1, 5, 8},
		{1, 1, 1, 9}}
	matrix := Matrix{board}
	verify, num := matrix.lookForNextMultiplier(1, 0)
	if verify && num != 2 {
		t.Error("should be true. Num should be 2. Num=", num)
	}
}
func TestLookForNextMultiplierThe4thPossibility(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{0, 1, 5, 8},
		{0, 1, 1, 9}}
	matrix := Matrix{board}
	verify, num := matrix.lookForNextMultiplier(0, 0)
	if !verify && num != 0 {
		t.Error("should be false. Num should be 0. Num=", num)
	}
}
func TestLookForNextMultiplierThe5thPossibility(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{1, 1, 5, 8},
		{0, 0, 1, 9}}
	matrix := Matrix{board}
	verify, num := matrix.lookForNextMultiplier(1, 0)
	if !verify && num != 1 {
		t.Error("should be false. Num should be 1. Num=", num)
	}
}
