package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Wooorld!",
		})
	})

	router.Run()
}
