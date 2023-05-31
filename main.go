package main

import "fmt"

func main() {
	var stringVariable string = "Hello"

	fmt.Println(stringVariable)

	var stringVariable2 = "World"

	// declared but not assigned
	var stringVariable3 string
	fmt.Println(stringVariable, stringVariable2, stringVariable3)

	stringVariable3 = "Mars"

	fmt.Println(stringVariable, stringVariable2, stringVariable3)

	stringVariable4 := "Venus"

	fmt.Println(stringVariable4)

	// Ints

	var ageOne int = 20

	var ageTwo = 30

	ageTree := 40

	fmt.Println(ageOne, ageTwo, ageTree)

	// Bits & memory

	var numOne int8 = 26

	//var numOne int8 = 266 // error out of range
	var numTwo int16 = 266 // fine

	fmt.Println(numOne, numTwo)

	// uint can be only positive
	// var numFour uint = -10 // error
	var numFour uint = 40 // good

	var numFive float32 = 42145.25412

	fmt.Println(numFour, numFive)

}
