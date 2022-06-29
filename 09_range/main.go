package main

import "fmt"

func main() {
	// Iterating through arrays using range  
	ids := []int{33, 2, 3, 4}
	sum := 0
	for idx, id := range ids {
		fmt.Printf("%d\n", idx)
		sum += id
	}		
	fmt.Printf("Sum: %d\n", sum)

	// Iterating through maps using range 
	emails := map[string]string{"Bob": "bob@gmail.com", "Aryan": "aryan@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
}
