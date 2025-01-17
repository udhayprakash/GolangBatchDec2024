package main

import "fmt"

/*
fmt.Printf, like printf in C, produces formatted output
------------------------------------------------------------------------
verb        Description
------------------------------------------------------------------------
%d          decimal integer
%x, %o, %b  integer in hexadecimal, octal, binary
%f, %g, %e  floating-point number: 3.141593 3.141592653589793 3.141593e+00
%t          boolean: true or false
%c          rune (Unicode code point)
%s          string
%q          quoted string "abc" or rune 'c'
%v          any value in a natural for mat
%T          type of any value
%%          literal percent sign (no operand)


Other formatting functions like log.Printf, fmt.Errorf, ... use same rules

*/

func main() {
	fmt.Println("Hello")
	fmt.Println("world") // automatically, cursor moves to newline after displaying

	fmt.Print("Hello1\n")
	fmt.Print("World1\n")

	fmt.Printf("Hello1 - %d\n", 1)
	fmt.Printf("World1 - %f\n", 14.23)

	fmt.Printf("World1 - %t\n", 14.23) // %!t(float64=14.23)
	fmt.Printf("World1 - %t\n", true)  //World1 - true

	fmt.Println("=====================================")
	// H	e	l	l	o	 	w	o	r	l	d
	// 0   	1	2	3	4	5	6	7	8	9	10

	s := "Hello World" // "नमस्ते"
	fmt.Println("len(s) =", len(s))
	fmt.Printf("len(s) =%d\n", len(s))

	for i := 0; i < len(s); i++ {
		// fmt.Println(i, s[i], string(s[i]))

		fmt.Printf("\t data type  : %T \n", s[i]) // data type

		fmt.Printf("\t rune       : %c \n", s[i])         // character
		fmt.Printf("\t quoted     : %q \n", s[i])         // quoted character
		fmt.Printf("\t string     : %s \n", string(s[i])) // string
		fmt.Printf("\t UTF8 code  : %v \n", s[i])         // any data type

		fmt.Printf("\t Hexadecimal: %x \n", s[i])
		fmt.Printf("\t Octal      : %o \n", s[i])
		fmt.Printf("\t Binary     : %b \n\n", s[i])
	}
}


/*
Assignment:
	from a given string, create new string by filtering
	only alphabets(lower & upper cases)

	upper alphabets from 65 to 90
	lower alphabets from 97 to 122

*/