package main

import (
	"fmt"
	"os"
)

func main() {
	file := "./archivo.txt"
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
