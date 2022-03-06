package main

import "fmt"

// Global variables can be declared here (NO SHORTHAND ALLOWED)

func main() {
	/**
		MAIN TYPES:

		string
		bool
		int (8-64)
		uint (8-64)
		uintptr
		byte	- alias for uint8
		rune 	- alias for int32
		float32 float64
		complex64 complex128
	**/

	var name string = "Brad"
	var age int = 32

	// While we can specify type, go can infer it as well
	var name_no_type = "Aryan"
	var age_no_type = 20

	// Shorthand (CAN ONLY BE USED INSIDE FUNCS)
	isCool := true

	fmt.Println(name, age, isCool)
	fmt.Println(name_no_type, age_no_type)
	fmt.Printf("%T\n", name_no_type)
	fmt.Printf("%T\n", age_no_type)

	// You can use "const" instead of "var" but then the variable won't be changeable
}
