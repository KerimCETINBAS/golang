package main

import "fmt"

func main() {

	for x := 0; x < 5; x++ {
		fmt.Println(x)

	}

	ages := [5]int8{15, 25, 32, 41, 54}

	for i := 0; i < len(ages); i++ {

		fmt.Printf("index =%v, ages =%v\n", i, ages[i])
	}

}
