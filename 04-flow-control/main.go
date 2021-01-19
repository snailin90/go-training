package main

import "fmt"

func main() {

	x := 15
	y := 10

	if x < y { // Parenthesis are not required
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "green"

	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is NOT blue or red")
	}

	//Switch

	switch color {
	case "red":
		fmt.Println("Color is red")
		break
	case "blue":
		fmt.Println("Color is red")
		break
	default:
		fmt.Println("Color is NOT blue or red")
	}


	// Long Method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//Short Method

	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// FizzBuzz

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// Range examples
	ids := []int{33, 76, 54, 23, 11, 2}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum ", sum)
	for i := 0; i < len(ids); i++ {
		fmt.Println(i, "- ID :", ids[i])
	}

	emails := map[string]string{"Bob": "bob@gmail.com", "mike": "mike@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name:", k)
	}
}
