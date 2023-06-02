package main

import "fmt"

func main() {

	defer fmt.Println("hello")

	defer fmt.Println("world")

	defFunc()

}

func defFunc() {
	for i := 0; i < 5; i++ {

		if i == 0 {
			defer fmt.Print(fmt.Sprintf("%v\n", i))
		} else {
			defer fmt.Print(i)
		}

	}
}
