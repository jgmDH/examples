package main

import "fmt"

type MiError struct {
	mensaje string
	codigo  int
}

func (e *MiError) Error() string {
	return fmt.Sprintf("Codigo: %d - Mensaje: %s", e.codigo, e.mensaje)
}

func main() {
	err := ProbandoMiErrorPersonal(1, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ProbandoMiErrorPersonal(num1, num2 int) error {
	if (num1 - num2) == 0 {
		return &MiError{
			mensaje: "La resta dio cero",
			codigo:  500,
		}
	}

	return nil
}
