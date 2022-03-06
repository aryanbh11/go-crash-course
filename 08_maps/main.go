package main

import "fmt"

func main() {
	// Define Map [key_type]value_type
	emails := make(map[string]string)

	// Assign
	emails["Bob"] = "bob@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))

	delete(emails, "Bob")

	fmt.Println(emails)

	// Declare map with key and values
	emails_with_kv := map[string]string{"Bob": "bob@gmail.com"}
	fmt.Println(emails_with_kv)
}
