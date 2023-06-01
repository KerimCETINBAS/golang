package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//stringFloat := "124.432ff"
	stringFloat := "124.432"

	numberFloat, err := strconv.ParseFloat(stringFloat, 32)

	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println(numberFloat)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(int32(numberFloat), math.Floor(numberFloat))

	}

}
