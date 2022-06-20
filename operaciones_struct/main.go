package main

import "fmt"

type Example struct {
	Descripcion int
}

func main() {
	examplesList := []Example{
		{Descripcion: 1},
		{Descripcion: 2},
		{Descripcion: 3},
	}

	// Lista de examples
	fmt.Println("Lista de examples ", examplesList)

	// Agregar
	examplesList = append(examplesList, Example{Descripcion: 23})
	fmt.Println("Lista de examples - Nuevo elemento ", examplesList)

	// Delete
	var exampleId int = 0 // Id del example a eliminar
	var positionExample int
	for indice := range examplesList {
		if indice == exampleId {
			positionExample = indice
		}
	}

	fmt.Println("Todos los elementos")
	for k, value := range examplesList {
		fmt.Printf("%d-%d ", k, value)
	}
	fmt.Println()

	fmt.Printf("Dame todo HASTA la posición (%d) - examplesList[:positionExample]\n", positionExample)
	for k, value := range examplesList[:positionExample] {
		fmt.Printf("%d-%d ", k, value)
	}
	fmt.Println()
	fmt.Printf("Dame todo DESDE la posición (%d) - examplesList[positionExample+1:]\n", positionExample+1)
	for k, value := range examplesList[positionExample+1:] {
		fmt.Printf("%d-%d ", k, value)
	}
	fmt.Println()

	examplesList = append(examplesList[:positionExample], examplesList[positionExample+1:]...)
	fmt.Println("Eliminado ", examplesList)
}
