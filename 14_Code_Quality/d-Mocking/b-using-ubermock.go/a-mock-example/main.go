package main

import "fmt"

type RealService struct{}

func (r *RealService) Add(a, b int) int {
	return a + b
}

func (r *RealService) Multiply(a, b int) int {
	return a * b
}

func main() {
	service := &RealService{}
	fmt.Println("Addition:", service.Add(3, 7))
	fmt.Println("Multiplication:", service.Multiply(3, 7))
}
