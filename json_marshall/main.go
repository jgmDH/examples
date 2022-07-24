package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Hola struct {
	Nombre string `json:"nombre"`
}

func main() {
	data := []byte(`{"Nombre":"juan", "Otro":"algo"}`)

	var h Hola
	err := json.Unmarshal(data, &h) // Ignorará el campo que no esta incluido en nuestro type Hola
	if err != nil {
		log.Print(err)
	}

	fmt.Printf("u: %+v\n", h)
}
