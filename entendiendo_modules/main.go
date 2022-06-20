package main

import (
	"fmt"
	"log"

	"github.com/juanmachuca95/examples_go/paquetedos"
	"github.com/juanmachuca95/examples_go/paqueteuno"
)

func main() {
	log.Println("__ Soy el Main")

	fmt.Println(paqueteuno.Message())

	fmt.Println(paquetedos.Message())
}
