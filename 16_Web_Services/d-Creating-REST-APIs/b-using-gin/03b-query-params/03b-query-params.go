package main

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Grocery represents a grocery item
type Grocery2 struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

var groceries = []Grocery2{
	{ID: 1, Name: "Apples", Price: 1.99, Category: "fruits"},
	{ID: 2, Name: "Bananas", Price: 0.99, Category: "fruits"},
	{ID: 3, Name: "Oranges", Price: 2.49, Category: "fruits"},
	{ID: 4, Name: "Carrots", Price: 0.49, Category: "vegetables"},
	{ID: 5, Name: "Potatoes", Price: 0.99, Category: "vegetables"},
	{ID: 6, Name: "Rice", Price: 1.29, Category: "cereal"},
}

func main() {
	server := gin.Default()

	// Endpoint to filter groceries
	server.GET("/groceries/", getGroceryItems)

	server.Run() // Listen and serve on http://localhost:8080
}

func getGroceryItems(ctx *gin.Context) {
	// Step 1: Parse the Query Params
	minPriceStr := ctx.Query("minPrice")
	maxPriceStr := ctx.Query("maxPrice")
	category := strings.ToLower(ctx.Query("category")) // Convert category to lowercase

	// Basic request, with no query params
	// If no filters are provided, return all groceries
	if minPriceStr == "" && maxPriceStr == "" && category == "" {
		ctx.JSON(200, groceries)
		return
	}

	// Convert minPrice and maxPrice to float64
	// based on block scope concept
	var minPrice, maxPrice float64
	var err error

	if minPriceStr != "" {
		minPrice, err = strconv.ParseFloat(minPriceStr, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid minPrice"})
			return
		}
	}

	if maxPriceStr != "" {
		maxPrice, err = strconv.ParseFloat(maxPriceStr, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid maxPrice"})
			return
		}
	}

	var filteredGroceries []Grocery2
	// Filter groceries
	for _, item := range groceries {
		// Check price range
		priceInRange := true
		if minPriceStr != "" && item.Price < minPrice {
			priceInRange = false
		}
		if maxPriceStr != "" && item.Price > maxPrice {
			priceInRange = false
		}

		// Check category (case-insensitive)
		categoryMatch := true
		if category != "" && strings.ToLower(item.Category) != category {
			categoryMatch = false
		}

		// Add to filtered list if both conditions are met
		if priceInRange && categoryMatch {
			filteredGroceries = append(filteredGroceries, item)
		}

	}

	// Return filtered groceries
	ctx.JSON(200, filteredGroceries)
}

/*
~ curl http://localhost:8080/groceries/?maxPrice=1

Range
	maxPrice, and minPrice

Category


*/
