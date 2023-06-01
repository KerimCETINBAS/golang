package main

import "fmt"

func main() {
	x := 0

	for {
		fmt.Println(x)
		x++

		if x > 5 {
			break
		}
	}

}
