package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}


func main() {
	alex := person{ firstName: "Alex", lastName: "Smith" }
	alex.firstName = "Ales"

	fmt.Printf("First Name: %+v", alex)
}