package main

import "github.com/gin-gonic/gin"

type User struct {
	Username string `json:"username"`
	Gender string `json:"gender"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "World")
	})

	return router
}

func postRouter(g *gin.Engine) *gin.Engine {
	g.POST("/user/add", func(ctx *gin.Context) {
		var u User
		ctx.BindJSON(&u)
		ctx.JSON(200, u)
	})

	return g
}

func main() {
	router := setupRouter()
	router = postRouter(router)
	router.Run()
}
