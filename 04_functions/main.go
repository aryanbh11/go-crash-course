package main

import "fmt"

/**
	- argument should always be followed by its type
	- u also have to specify a return type (unless its void then u leave it empty)
**/
func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(greeting("Aryan"))
	fmt.Println(getSum(3, 4))
}
