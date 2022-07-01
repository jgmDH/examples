package main

import (
	"fmt"
	"sync"
	"time"
)

type Carro struct {
	Piloto    string
	Recorrido float64
}

func main() {
	corredores := 3
	c1 := make(chan Carro, corredores)

	salida(c1)
	meta(c1)
}

func salida(c1 chan Carro) {
	nomCorredores := []Carro{
		{Piloto: "Nico", Recorrido: 0},
		{Piloto: "María", Recorrido: 0},
		{Piloto: "Sol", Recorrido: 0},
	}

	wg := sync.WaitGroup{}
	for _, corr := range nomCorredores {
		wg.Add(1)
		go run(c1, corr, &wg)
	}

	wg.Wait()
	close(c1)
}

func run(c1 chan Carro, carro Carro, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 3)
	carro.Recorrido = 100
	c1 <- carro
}

func meta(c1 chan Carro) {
	var win = false
	for val := range c1 {
		if val.Recorrido == 100 && !win {
			fmt.Println("🏆 ", val.Piloto, " - Recorrido: ", val.Recorrido)
			win = true
		} else {
			fmt.Println(val.Piloto, " - Recorrido: ", val.Recorrido)
		}

	}
}
