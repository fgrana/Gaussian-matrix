package reader

import (
	"Gauss2/functions"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func FileToMatrix() functions.Matrix {

	dat, err := os.ReadFile("matrix.txt")
	check(err)
	s := strings.Split(string(dat), "\n")
	board := [][]float64{}
	for _, v := range s {
		s1 := strings.Split(v, ",")
		line := []float64{}
		for _, v1 := range s1 {
			variable, _ := strconv.ParseFloat(v1, 64)
			line = append(line, variable)
		}
		board = append(board, line)
	}
	var m = functions.Matrix{board}
	return m
}
