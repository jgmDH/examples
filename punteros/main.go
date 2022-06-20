package main

import (
	"fmt"
)

func main() {
	var nombre string = "juan"
	fmt.Println(nombre, &nombre)

	var numero *int
	fmt.Println(numero, &numero)
	asignar := 2
	numero = &asignar
	fmt.Println(*numero, &numero)

	*numero = 3
	fmt.Println(**&numero, &numero)

	/* **&numero = 4
	fmt.Println(*numero) */
}
