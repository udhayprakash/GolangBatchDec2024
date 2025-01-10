package main

type Calculator interface{
	Add(a, b int) int
}

type RealCalculator struct{}


func (r RealCalculator)Add(a, b int) int {
	return a+ b
}