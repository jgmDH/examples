package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("Hola Bootcampers!")
	err := os.WriteFile("./archivo.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
