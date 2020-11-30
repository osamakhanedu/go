package main

import "fmt"

func main() {

	// go has few primitive data types
	// int, float, bool, and complex
	// each has different version
	// which corresponde to memory size allocated to them
	// int, int8, int32, int 64
	// float32, float 64
	// bool
	// complex

	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Osama"
	fmt.Println(firstName)

	c := complex(3, 4)

	fmt.Println(c)

	r, im := real(c), imag(c)

	fmt.Println(r, im)
}
