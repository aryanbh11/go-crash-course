package main

import "fmt"

func main() {
	a := 5
	b := &a // b is now a pointer to a (stores address of a)

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) // type of b 

	// Use * to read val from address
	fmt.Println(*b) // prints value at the address b stores (note that address b stores = value of b)
	fmt.Println(*&a) // prints the value at the address of a i.e value of a 

	// Changing val with pointer 
	*b = 10
	fmt.Println(a)
}
