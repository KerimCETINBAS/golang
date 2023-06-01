package main

import "fmt"

func updateName(x *string) {
	*x = "lunafreya"
}
func main() {

	name := "tifa"

	fmt.Println("memory address of name :", &name)

	m := &name

	fmt.Println("memory address :", m)
	fmt.Println("value att memory address :", *m)

	updateName(m)

	// original name changed

	fmt.Println("name has changed to :", name)
}
