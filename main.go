package main

import (
	"Gauss2/reader"
	"fmt"
	"os"
)

func main() {
	m := reader.FileToMatrix()
	verify, unknowns := m.MGauss()
	f, err := os.Create("matrixAnswer.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	err = m.AnswerToFile(f, verify, unknowns)
	if err != nil {
		fmt.Println(err)
	}
}
