package main

import (
	"fmt"
	"math"
)

func SayHello(name string) {
	fmt.Println("Hello ", name)
}

func SayHelloToEveryone(names []string, function func(name string)) {
	for _, name := range names {
		function(name)
	}
}

func CircleArea(radius float32) float32 {
	return radius * math.Pi * math.Pi
}

func GetFavPlaces(p string) (string, error) {

	places := []string{"limsa lominsa", "gridania", "Ul'dah", "Ala Mhigo"}

	for _, place := range places {

		if place == p {
			return p, nil
		}
	}

	return "", fmt.Errorf("Place can not found")

}

func main() {
	names := [4]string{"Alphinaud", "Alisae", "Y'shtola", "Thancred"}

	SayHelloToEveryone(names[0:], SayHello)

	area1 := CircleArea(10.4)
	area2 := CircleArea(4)

	fmt.Printf("area1: %f\narea2: %f\n", area1, area2)

	favPlace1, err1 := GetFavPlaces("Ul'dah")
	favPlace2, err2 := GetFavPlaces("Mor dhona")

	fmt.Println(favPlace1, err1)
	fmt.Println(favPlace2, err2)
}
