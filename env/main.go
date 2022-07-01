package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("DEFINIENDO_VARIABLE", "Hola Bootcampers")

	fmt.Println(os.Getenv("DEFINIENDO_VARIABLE"))

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Printf("%s: %s\n", pair[0], pair[1])
	}
}
