package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInput(promt string) string {
	r := bufio.NewReader(os.Stdin)
	fmt.Print(promt)
	v, _ := r.ReadString('\n')
	return strings.TrimSpace(v)
}

func parseRace(i string) Race {
	switch i {
	case "Lalafel":
		fmt.Println("You are selected", Race(Lalafel))
		return Race(Lalafel)
	case "Hyur":
		fmt.Println("You are selected", Race(Hyur))
		return Race(Hyur)
	case "Miqote":
		fmt.Println("You are selected", Race(Miqote))
		return Race(Miqote)
	default:
		fmt.Println("Please select valid race")
		return parseRace(strings.TrimSpace(ReadInput("Please select your race: ")))
	}
}

func main() {

	fmt.Println("create new hero")

	name := ReadInput("Please enter your name: ")

	fmt.Println("Available races are Hyur, Lalafel, Miqote")
	race := parseRace(strings.TrimSpace(ReadInput("Please select your race: ")))

	hero := createHero(name, race)
	fmt.Printf("You have created hero with name %v, and race %v", hero.name, hero.race)

}
