package main

import (
	"fmt"
	"os"
)

func main() {
	files, err := os.ReadDir("./")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(files) // []fs.DirEntry
	for _, value := range files {
		fmt.Println(value.Name())
	}

	filename := "./archivo.txt"
	f, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n", "Info file")
	fmt.Println(f.Name(), f.Size(), f.ModTime())

}
