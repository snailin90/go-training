package main

import "fmt"

func main() {
	a := 5
	b := &a // this kind of assigment is getting the memory address where the values a is located. ex: 0xc00012078

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) // it will return *int  , it means is a pointer type

	// Use to read val from address

	fmt.Println(*b) // if you want to print the value of the pointer , you put asterisk in front of the variable

	// Change value of pointer

	*b = 10
	fmt.Println(a) // [a] variable is updated because the pointer was updated

}
