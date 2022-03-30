package main

import (
	"Gauss2/reader"
	"fmt"
	"os"
)

func main() {
	m := reader.Reader()
	verify, incognitas := m.MGauss()
	f, err := os.Create("matrizRespuesta.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	err = m.EscribirSolucion(f, verify, incognitas)
	if err != nil {
		fmt.Println(err)
	}
}
