package main 

import "fmt"

// Block Level Scoping 

func main(){
	y:= 20 
	fmt.Println(y) // 20 

	{  // block scope
		y := 10 
		fmt.Println(y) // 10
	}
	
	fmt.Println(y) // 20
}