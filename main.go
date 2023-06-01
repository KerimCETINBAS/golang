package main

import "fmt"

func main() {

	myHero := createHero("Yumia Lemon", Race(Miqote))

	fmt.Println(myHero)

	myHero.rename("Innocent Ahegao")

	fmt.Println(myHero.getRace())
}
