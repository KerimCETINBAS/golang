package main

import "fmt"

func main() {
	// Print
	fmt.Print("Hello, ")
	fmt.Print("world! ")
	fmt.Print("\n")

	// Print line
	fmt.Println("Hello wolrld!")
	fmt.Println("Merhaba dunya!")

	age := 24
	name := "jdoe"
	fmt.Println("My age is", age, "and my name is", name)

	// Printf (formatted strings)

	// %_ = format specifiers
	fmt.Printf("My age is %v and my name is %v \n", age, name)

	fmt.Printf("My age is %q and my name is %q \n", age, name)

	fmt.Printf("May age is of type %T \n", age)

	// print float
	fmt.Printf("you scorec %0.1f points!\n", 255.55)
	// without point
	fmt.Printf("you scorec %0.f points!\n", 255.55)

	// SprintF (save formatted strings)

	str := fmt.Sprintf("My age is %v and my name is %v \n", age, name)

	println(str)
}
