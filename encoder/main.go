package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Bootcamper struct {
	Nombre string
}

func main() {
	nBootcamper := &Bootcamper{Nombre: "Nico"}

	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(nBootcamper)
	if err != nil {
		fmt.Println(err)
		return
	}

}
