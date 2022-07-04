package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/* type Pelicula struct {
	Titulo     string
	Disponible bool
}

type Bootcamper struct {
	Nombre string
}

var peliculas = []Pelicula{
	{Titulo: "Metaverso", Disponible: true},
	{Titulo: "Piratas del Caribe", Disponible: false},
}
*/
func main() {
	http.HandleFunc("/bootcampers", Bootcampers)
	/* 	http.HandleFunc("/peliculas", GetPeliculas) */

	fmt.Println("Iniciando el servidor . . . 🚀 en el puerto 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func Bootcampers(w http.ResponseWriter, r *http.Request) {
	type Message struct {
		Message string
	}

	m := Message{Message: "👋 Hola Bootcampers"}
	data, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(data)
}

/* func GetPeliculas(w http.ResponseWriter, r *http.Request) {
	if http.MethodGet != r.Method {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(peliculas)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(data)
	return
} */
