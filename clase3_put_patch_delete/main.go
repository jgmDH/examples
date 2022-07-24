package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Id       int
	Nombre   string
	Telefono string
}

type Contacto struct {
	Id       int
	Nombre   string
	Telefono string
}

var contactos = []Contacto{
	{Id: 1, Nombre: "Juan", Telefono: "+54 2341234233"},
	{Id: 2, Nombre: "Maria", Telefono: "+54 2343223"},
}

func main() {
	router := gin.Default()

	cr := router.Group("contactos")
	cr.PUT("/:id")

	router.Run()
}

func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id incorrecto"}) //400
		return
	}

	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if req.Id != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El campo id es requerido"})
		return
	}

}
