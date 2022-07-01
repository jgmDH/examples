package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6} // len(s) es par

	for i := 0; i < len(s)-1; i++ {
		fmt.Println(s[i : i+2]) // Retorno de a pares.
	}
}
