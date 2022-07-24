package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/variosparams/:id/:name", VariosParams)

	r.Run()
}

func VariosParams(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")

	message := fmt.Sprintf("El id es %s - nombre %s", id, name)
	c.JSON(200, message)
}

//http://localhost:8080/variosparams/1/juan
