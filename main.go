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

func main() {

	fmt.Println("create new hero")
	heroName := ReadInput("Please enter hero name: ")
	heroRace := parseRace(ReadInput("Please enter your heros race;\nAvailable races are Miqote, Hyur, Lalafel: "))
	newhero := createHero(heroName, heroRace)

	fmt.Println(newhero)
}
