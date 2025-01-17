package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	var cNum complex128 = 3 + 4i
	fmt.Printf("value = %v  datatype  = %T\n", cNum, cNum)
	fmt.Println()

	fmt.Println("cmplx.Abs(cNum) =", cmplx.Abs(cNum))

	fmt.Println("cmplx.Abs(3+4i)        =", cmplx.Abs(3+4i))
	// sqrt of (sum of real value squre and imag value square)

	fmt.Println("cmplx.Conj(3+4i)       =", cmplx.Conj(3+4i))

	fmt.Println("cmplx.Exp(1i*math.Pi)+1=", cmplx.Exp(1i*math.Pi)+1)
	fmt.Println()

	r, theta := cmplx.Polar(2i)
	fmt.Printf("r: %.1f, θ: %.1f*π", r, theta/math.Pi)
	fmt.Println()

	var x complex128 = complex(1, 2)           			// 1+2i
	var y complex128 = complex(3, 4)           			// 3+4i
	fmt.Println("x * y       			=", x*y)       // "(-5+10i)"
	fmt.Println("real(x * y) 			=", real(x*y)) // "-5"
	fmt.Println("imag(x * y) 			=", imag(x*y)) // "10"
	fmt.Println()

	fmt.Println("1i * 1i     			=", 1i*1i)            // "(-1+0i)", i^2 = -1
	fmt.Println("cmplx.Sqrt(-1)			=", cmplx.Sqrt(-1)) // "(0+1i)"

}
