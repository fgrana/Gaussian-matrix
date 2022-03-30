package functions

import (
	"fmt"
	"os"
)

func (m *Matrix) AnswerToFile(f *os.File, verify string, Unknowns []float64) error {

	if verify == "Incompatible" {
		_, err := f.WriteString("incompatible\n")
		if err != nil {
			return err
		}
	} else {
		count := 0
		for i := len(Unknowns) - 1; i > -1; i-- {
			count++
			_, err1 := fmt.Fprintf(f, "Unknown number %v = %v \n", count, Unknowns[i])
			if err1 != nil {
				return err1
			}

		}
	}
	for i := 0; i < len(m.Board); i++ {
		_, err1 := fmt.Fprintf(f, " %v \n", m.Board[i])
		if err1 != nil {
			return err1
		}

	}
	return nil
}
