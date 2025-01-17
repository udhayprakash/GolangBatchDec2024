package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func main() {
	order1 := order{
		ordId:      456,
		customerId: 56,
	}

	v := reflect.ValueOf(order1)
	// NumField() method returns the number of fields in a struct
	// Field(i int) method returns the reflect.Value of the ith field.
	fmt.Println("Number of fields", v.NumField())
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
	}
}