package functions

import (
	"fmt"
	"os"
)

func (m *Matriz) EscribirSolucion(f *os.File,verify string, incognitas []float64) error {

	if verify == "Incompatible" {
		_, err := f.WriteString("incompatible\n")
		if err != nil {
			return err
		}
	} else {
		count := 0
		for i := len(incognitas) - 1; i > -1; i-- {
			count++
			_,err1:=fmt.Fprintf(f,"incognita numero %v = %v \n", count, incognitas[i])
			if err1 != nil {
				return err1
			}

		}
	}
	for i := 0; i < len(m.Board); i++ {
		_, err1 := fmt.Fprintf(f," %v \n", m.Board[i])
		if err1 != nil {
			return err1
		}

	}
	return nil
}