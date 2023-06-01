package main

import "fmt"

func main() {

	x := 6
	y := 7

	fmt.Printf("x=%v, y=%v\n-------------\n", x, y)
	fmt.Println("x + y =", x+y)
	fmt.Println("y - x =", y-x)
	fmt.Println("x * y =", x*y)
	fmt.Println("x / 2 =", x/2)
	fmt.Println("reminder of y / x =", y%x)

	x++
	fmt.Println(x)
	y--
	fmt.Println(y)
}
