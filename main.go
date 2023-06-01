package main

import "fmt"

func main() {

	//var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}

	fmt.Println(ages, len(ages))

	names := [4]string{"jhon", "mary", "baby", "jane"}

	names[1] = "anna"
	fmt.Println(names, len(names))

	// slices (uses arrays under the hood)

	var scores = []int{100, 50, 60}

	scores[2] = 25

	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// slice ranges

	rangeOne := names[1:3]
	var rangeTwo = names[1:4]
	fmt.Println(rangeOne, rangeTwo)
}
