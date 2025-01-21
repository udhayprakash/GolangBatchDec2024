package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	groceries := []string{"Apples", "Bananas", "Oranges"}

	server.GET("/groceries/:index", func(ctx *gin.Context) {
		index, _ := strconv.Atoi(ctx.Param("index"))
		if index < 0 || index >= len(groceries) {
			ctx.JSON(404, gin.H{"error": "Item not found"})
			return
		}

		ctx.JSON(200, gin.H{
			"message":        "Welcome to the Grocery Store!",
			"itemsAvailable": groceries,
			"selectedItem":   groceries[index],
		})
	})
	server.Run() // default port is 8080
}

/*

- curl http://localhost:8080/groceries/1
{"itemsAvailable":["Apples","Bananas","Oranges"],"message":"Welcome to the Grocery Store!","selectedItem":"Bananas"}
- curl http://localhost:8080/groceries/
404 page not found
- curl http://localhost:8080/groceries/0
{"itemsAvailable":["Apples","Bananas","Oranges"],"message":"Welcome to the Grocery Store!","selectedItem":"Apples"}
- curl http://localhost:8080/groceries/-1
{"error":"Item not found"}
- curl http://localhost:8080/groceries/3
{"error":"Item not found"}
-


*/
