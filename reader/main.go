package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hola bootcampers")
	b, err := io.ReadAll(r)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(b)
	fmt.Println(string(b))
}
