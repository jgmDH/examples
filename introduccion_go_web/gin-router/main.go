package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("PORT", "9000")
	r := gin.Default()
	r.GET("/bootcampers", Bootcampers)

	r.Run()
}

func Bootcampers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "👋 Hola bootcampers!"})
}
