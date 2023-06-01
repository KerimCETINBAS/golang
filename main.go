package main

import "fmt"

func changeName(n string) {
	n = "Lunafreya"
}

func changeName2(n string) string {
	n = "Lunafreya"
	return n
}

func changeName3(n *string) {
	*n = "Lucius"
}

func addNumber(nums map[int]int) {
	nums[3] = 3
}

func main() {

	// non-pointer values
	/*
		strings
		ints
		floats
		booleans
		arrays
		structs
	*/

	name := "Loctis"

	changeName(name)
	// name still noctis
	fmt.Println(name)

	name = changeName2(name)

	fmt.Println(name)

	// passing its pointer

	changeName3(&name)

	// name changed
	fmt.Println(name)

	// Pointer wrapper values
	/*
		    slices? tutorial says, but actually not work for slices
			maps
			functions */

	var numbers map[int]int = map[int]int{1: 1, 2: 2}

	addNumber(numbers)
	fmt.Println(numbers)
}
