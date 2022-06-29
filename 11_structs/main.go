package main

import "fmt"

// Define person struct 
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting function (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName
}

// hadBirthday method (pointer reciever)
func (p *Person) hadBirthday() {
	p.age++
}

func main() {
	// Init person usings struct 
	person1 := Person{firstName: "Aryan", lastName: "Bhatia", city: "Sydney", gender: "Male", age: 21}

	/**
		Alternative:
		person1 := Person{"Aryan", "Bhatia", "Sydney", "Male", 21}
	**/
	
	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1.age)

	person1.hadBirthday()
	fmt.Println(person1.age)

	fmt.Println(person1.greet())
}
