package main

import (
	"context"
	"fmt"
)

/*
costly resource
	memroy
	cpu

limited reosurces
	db connections pool
	socket connections
*/

type keyOne string
type keyTwo string

func main() {

	type keyval string

	// Generating Parent-Context
	ctx := context.Background()

	ctx = context.WithValue(ctx, keyOne("one"), "valueOne")
	ctx = context.WithValue(ctx, keyTwo("one"), "valueTwo")

	fmt.Println(ctx.Value(keyOne("one")).(string))
	fmt.Println(ctx.Value(keyTwo("one")).(string))

	// --------------------------------------
	// Here, context is created with values
	vctx := context.WithValue(ctx, keyval("request-id"), keyval("123"))

	// fmt.Println(vctx)
	fmt.Println("request-id", vctx.Value(keyval("request-id")))
}
