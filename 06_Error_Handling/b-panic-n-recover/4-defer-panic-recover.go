package main

import "fmt"

func main() {
	fmt.Println("in main() function")
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\tRecovered in f", r)
		}
	}()
	fmt.Println("\tCalling g.")
	g(0)
	fmt.Println("\tReturned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("\t\tPanicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("\t\tDefer in g", i)
	fmt.Println("\t\tPrinting in g", i)
	g(i + 1) // recursive call
}

/*

Explanation:
----------------
g(0)
                Defer in g 0
g(1)
                Defer in g 1
g(2)
                Defer in g 2

g(3)
                Defer in g 3

g(4)
				panic






*/