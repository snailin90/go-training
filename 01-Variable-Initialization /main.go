package main

import "fmt"

func main(){
	//MAIN Types
	//string
	//bool
	//int
	//int int8 int16 int32 int64
	//uint uint8 uiny16 uint32 uint64 uintptr
	//byte - alias for uint8
	//rune - alias for int32
	//float32 float64
	//complex64 complex128

	var name string = "Mario" // is not necessary to put the "string" type because go infer the type because of the value.

	var age int32 = 37 //  if you declare a variable in GO , you need to use , if you don't you will get a compilation error.
	// also no need to put the "int" type, because GO infer the type.

	var isCool = true

	const lastName = "INOA"
	// lastName = "test"    Const variable can not change the value once is set.

	fmt.Println(name, age)

	fmt.Printf("%T\n", name) // this allow you to get the type of specific variable
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)

	//ShortHand declaration
	carModel := "Toyota" // this type of assigment can be only inside function

	fmt.Println(carModel)
	fmt.Printf("%T\n", carModel)

	size := 1.3
	fmt.Printf("%T\n", size)

	player, position := "PlayerName", "Player Position"
	fmt.Println(player, position)
}
