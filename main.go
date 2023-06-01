package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends!"

	// strings
	fmt.Println(strings.Contains(greeting, "hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "re"))

	fmt.Println(strings.Split(greeting, "\r"))
	fmt.Println(strings.Split(greeting, " "))

	fmt.Println("original string value =", greeting)

	// sort
	ages := []int{45, 15, 20, 30, 79, 43, 18}

	sort.Ints(ages)
	//sort.IntSlice(ages).Sort()
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)

	fmt.Println(index)

	names := []string{"jhon", "jane", "baby", "doe", "anna"}

	sort.Strings(names)

	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "jane"))
}
