package main

import "fmt"

func main() {

	ages := [5]int8{15, 25, 32, 41, 54}

	for i, value := range ages {

		fmt.Printf("index =%v, ages =%v\n", i, value)
	}

	for _, value := range ages {
		// no index
		fmt.Println("age is", value)
	}

}
