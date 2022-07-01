package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hola mundo!")
	}

	fmt.Println(handler)
}
