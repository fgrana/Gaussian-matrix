package functions

import (
	"fmt"
	"testing"
)

func TestCheckearAbajoCompletoAllOkey(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{0, 1, 5, 8},
		{0, 0, 1, 9}}
	matriz := Matrix{board}
	verify, i, j, b := matriz.lookForCerosInLower()
	if verify && i != 0 && j != 0 && b != 0 {
		texto, _ := fmt.Printf("verify should be true. i = %v, j = %v, b = %v", i, j, b)
		t.Error(texto)
	}
}
func TestCheckearAbajoCompleto1th(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{1, 1, 5, 8},
		{0, 0, 1, 9}}
	matriz := Matrix{board}
	verify, i, j, b := matriz.lookForCerosInLower()
	if !verify || i != 1 || j != 0 || b != 0 {
		texto, _ := fmt.Printf("verify should be true. i = %v, j = %v, b = %v", i, j, b)
		t.Error(texto)
	}
}
func TestCheckearAbajoCompleto2th(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{0, 1, 5, 8},
		{1, 0, 1, 9}}
	matriz := Matrix{board}
	verify, i, j, b := matriz.lookForCerosInLower()
	if !verify || i != 2 || j != 0 || b != 0 {
		texto, _ := fmt.Printf("verify should be true. i = %v, j = %v, b = %v", i, j, b)
		t.Error(texto)
	}
}
func TestCheckearAbajoCompleto3th(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{0, 1, 5, 8},
		{0, 1, 1, 9}}
	matriz := Matrix{board}
	verify, i, j, b := matriz.lookForCerosInLower()
	if !verify || i != 2 || j != 1 || b != 1 {
		texto, _ := fmt.Printf("verify should be true. i = %v, j = %v, b = %v", i, j, b)
		t.Error(texto)
	}
}
func TestCheckearAbajoCompleto4th(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{0, 1, 5, 8},
		{1, 1, 1, 9}}
	matriz := Matrix{board}
	verify, i, j, b := matriz.lookForCerosInLower()
	if !verify || i != 2 || j != 0 || b != 0 {
		texto, _ := fmt.Printf("verify should be true. i = %v, j = %v, b = %v", i, j, b)
		t.Error(texto)
	}
}
func TestCheckearAbajoCompleto5th(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{1, 1, 5, 8},
		{1, 0, 1, 9}}
	matriz := Matrix{board}
	verify, i, j, b := matriz.lookForCerosInLower()
	if !verify || i != 1 || j != 0 || b != 0 {
		texto, _ := fmt.Printf("verify should be true. i = %v, j = %v, b = %v", i, j, b)
		t.Error(texto)
	}
}
func TestCheckearAbajoCompleto6th(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{1, 1, 5, 8},
		{1, 1, 1, 9}}
	matriz := Matrix{board}
	verify, i, j, b := matriz.lookForCerosInLower()
	if !verify || i != 1 || j != 0 || b != 0 {
		texto, _ := fmt.Printf("verify should be true. i = %v, j = %v, b = %v", i, j, b)
		t.Error(texto)
	}
}

func TestBuscar0Abajo(t *testing.T) {
	board := [][]float64{
		{0, 0, 0, 0},
		{1, 0, 0, 0},
		{0, 0, 0, 0}}
	matriz := Matrix{board}
	verify, i, j := matriz.buscar0Abajo()
	if !verify {
		t.Error("should be true", i, j)
	}
	board = [][]float64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 0, 0, 0}}
	matriz = Matrix{board}
	verify, i, j = matriz.buscar0Abajo()
	if !verify {
		t.Error("should be true", i, j)
	}
	board = [][]float64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 1, 0, 0}}
	matriz = Matrix{board}
	verify, i, j = matriz.buscar0Abajo()
	if !verify {
		t.Error("should be true", i, j)
	}
	board = [][]float64{
		{1, 1, 1, 1},
		{0, 1, 1, 1},
		{0, 0, 1, 1}}
	matriz = Matrix{board}
	verify, i, j = matriz.buscar0Abajo()
	if verify {
		t.Error("should be false", i, j)
	}
}

func TestCheckearCompatibleTakeThe1thPossibility(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{1, 1, 5, 8},
		{1, 1, 1, 9}}
	matriz := Matrix{board}
	verify, num := matriz.buscarProxMulti(1, 0)
	if verify && num != 0 {
		t.Error("should be true. Num should be 0. Num=", num)
	}
}
func TestCheckearCompatibleTakeThe2thPossibility(t *testing.T) {
	board := [][]float64{
		{0, 3, 4, 4},
		{1, 1, 5, 8},
		{1, 1, 1, 9}}
	matriz := Matrix{board}
	verify, num := matriz.buscarProxMulti(1, 0)
	if verify && num != 2 {
		t.Error("should be true. Num should be 1. Num=", num)
	}
}
func TestCheckearCompatibleTakeThe3thPossibility(t *testing.T) {
	board := [][]float64{
		{0, 3, 4, 4},
		{0, 1, 5, 8},
		{1, 1, 1, 9}}
	matriz := Matrix{board}
	verify, num := matriz.buscarProxMulti(1, 0)
	if verify && num != 2 {
		t.Error("should be true. Num should be 2. Num=", num)
	}
}
func TestCheckearCompatibleTakeThe4thPossibility(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{0, 1, 5, 8},
		{0, 1, 1, 9}}
	matriz := Matrix{board}
	verify, num := matriz.buscarProxMulti(0, 0)
	if !verify && num != 0 {
		t.Error("should be false. Num should be 0. Num=", num)
	}
}
func TestCheckearCompatibleTakeThe5thPossibility(t *testing.T) {
	board := [][]float64{
		{1, 3, 4, 4},
		{1, 1, 5, 8},
		{0, 0, 1, 9}}
	matriz := Matrix{board}
	verify, num := matriz.buscarProxMulti(1, 0)
	if !verify && num != 1 {
		t.Error("should be false. Num should be 1. Num=", num)
	}
}
