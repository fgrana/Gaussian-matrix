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
func Reader() functions.Matriz {

	dat, err := os.ReadFile("matriz.txt")
	check(err)
	s := strings.Split(string(dat), "\n")
	board := [][]float64{}
	for _, v := range s {
		s1 := strings.Split(v, ",")
		miniBoard := []float64{}
		for _, v1 := range s1 {
			f, _ := strconv.ParseFloat(v1, 64)
			miniBoard = append(miniBoard, f)
		}
		board = append(board, miniBoard)
	}
	var m = functions.Matriz{board}
	return m
}
