package main

import "fmt"

func main() {

	drinks := map[string]int8{
		"cola":      4,
		"fanta":     6,
		"fuse tea":  12,
		"olong tea": 4,
		"balck tea": 2,
		"water":     1,
	}

	fmt.Println(drinks)
	fmt.Println(drinks["water"])

	users := map[int]string{
		1: "Exc4l1bur16",
		2: "Exc3cuti6n3r",
		3: "King avatar",
		4: "IIProPlayerII",
	}

	fmt.Println(users)

	users[1] = "Hail nemo"

	fmt.Println(users)
}
