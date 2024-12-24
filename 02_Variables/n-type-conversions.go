package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	d := 10

	// to get the data type
	fmt.Printf("%T\n", d)
	fmt.Println(reflect.TypeOf(d))

	// Explicit type conversion
	fmt.Println(int(d), int8(d), int16(d), int32(d), int64(d)) // 10 10 10 10 10

	d = 300
	fmt.Println(int(d), int8(d), int16(d), int32(d), int64(d)) // 300 44 300 300 300
	fmt.Println()

	f := 5.1
	i := int(f) // convert float to int

	fmt.Println(f, reflect.TypeOf(f)) // 5.1 float64
	fmt.Println(i, reflect.TypeOf(i)) // 5 int

	// Integers ===============================================================
	var a uint16 = 0x10fe //  bit pattern: 0001 0000 1111 1110

	// Truncation
	b := int8(a) // -2        bit pattern:           1111 1110

	// Sign extension
	c := uint16(b) // 0xfffe  bit pattern: 1111 1111 1111 1110
	fmt.Printf(`
		a = %v
		b = %v
		c = %v
	`, a, b, c)
	fmt.Println()

	// Floats ================================================================
	var f1 float64 = 1.9
	n := int64(f1) // 1
	n = int64(-f1) // -1

	n = 1234567890
	g := float32(n) // 1.234568e+09
	fmt.Println(g)

	// Integer to String ====================================================
	fmt.Println("string(97) =", string(97)) // "a"
	fmt.Println("string(-1) =", string(-1)) // "\ufffd" == "\xef\xbf\xbd"
	fmt.Println()

	// fmt.Println("string(97.5) =", string(97.5))
	// cannot convert 97.5 (untyped float constant) to type string

	// To get the decimal string representation of an integer:
	fmt.Println("strconv.Itoa(97)=", strconv.Itoa(97)) // "97"
	fmt.Println("strconv.Itoa(-1)=", strconv.Itoa(-1)) // "-1"

	
	// fmt.Println("strconv.Itoa(97.5)=", strconv.Itoa(97.5)) // "97"
	// cannot use 97.5 (untyped float constant) as int value in argument to strconv.Itoa (truncated)

	
	// string to int
	// fmt.Println(int("55")) //  cannot convert "55" (untyped string constant) to type int

	fmt.Println(int('5')) // 53, is ASCII code for '5' as character

	fmt.Println(strconv.Atoi("55"))

	val, _ := strconv.Atoi("55")
	fmt.Printf("%v - %T\n", val, val)

	
}
