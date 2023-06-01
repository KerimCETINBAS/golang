package main

import "fmt"

func main() {

	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("Age less then 30")
	} else if age < 40 {
		fmt.Println("Age is less then 40")
	} else {
		fmt.Println("Age is equal or higher than 40")
	}

	names := []string{"jhon", "mary", "jane", "baby", "nana"}

	for index, name := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}

		fmt.Printf("the name at pos %v is %v\n", index, name)
	}

}
