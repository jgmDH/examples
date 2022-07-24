package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/example", MultiplesMiddlewares(MyHandler())...)
	r.Run()
}

func MyHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hemos pasado todos los middlewares"})
	}
}

func MultiplesMiddlewares(miHandler gin.HandlerFunc) []gin.HandlerFunc {
	middlewares := []gin.HandlerFunc{
		ValidateToken(),
		HolaBootcampers(),
	}

	middlewares = append(middlewares, miHandler)
	return middlewares
}

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token == "" {
			ctx.AbortWithStatusJSON(403, nil)
			return
		}

		ctx.Next()
	}
}

func HolaBootcampers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Algun middleware más.")

		ctx.Next()
	}
}
