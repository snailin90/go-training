package main

import (
	"fmt"
	"strconv"
)

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Snailin"))
	fmt.Println(getSum(2, 3))


	// Init person using struct

	person1 := Person{firstName: "Snailin", lastName: "Inoa", city: "Santo Domingo", gender: "f", age: 30}
	fmt.Println(person1)

	//Alternative
	person2 := Person{"Snailin", "Inoa", "Santo Domingo", "m", 30}

	fmt.Println(person1.firstName)
	fmt.Println(person2)

	// Value Recievers , method that do calculation don't change the object
	// Pointer Recievers , method that change something
	person1.hasBirthday()
	person1.getMerried("Felix")
	fmt.Println(person1.greet())
}

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting Method (Value receiver)

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMerried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}

}

