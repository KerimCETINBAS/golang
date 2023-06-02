package main

import "fmt"

type creature interface {
	getName() string
}

type monster struct {
	name string
	race string
}
type character struct {
	name string
	race string
}

func (c character) getName() string {
	return c.name
}

func (c monster) getName() string {
	return c.name
}

func printName(c creature) {
	fmt.Println(c.getName())
}

func main() {

	newMonster := monster{
		name: "goblin slayer",
		race: "hobgoblin",
	}

	printName(newMonster)
}
