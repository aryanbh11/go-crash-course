package main

import "fmt"

func main() {
	// Declare Array
	var fruitArr [2]string

	// Assign
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Print
	fmt.Println(fruitArr)

	// Shorthand
	fruitArrShort := [2]string{"Apple", "Orange"}
	fmt.Println(fruitArrShort)

	// Declare + Assign Slice
	fruitSlice := []string{"Apple", "Orange"}
	fruitSlice = append(fruitSlice, "Banana") // Only for slices

	fmt.Println(fruitSlice)

	// Array vs Slices
	// The only diff is that arrays are fixed length
}
