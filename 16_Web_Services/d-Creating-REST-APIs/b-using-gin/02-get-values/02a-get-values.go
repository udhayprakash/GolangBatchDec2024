package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	groceries := []string{"Apples", "Bananas", "Oranges"}

	server.GET("/groceries", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to the Grocery Store!",
			"items":   groceries,
		})
	})

	// server.Run("localhost:8000")
	server.Run() // default port is 8080
}

/*
- curl http://localhost:8080/groceries
{"items":["Apples","Bananas","Oranges"],"message":"Welcome to the Grocery Store!"}


*/