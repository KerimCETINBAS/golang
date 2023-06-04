package main

import (
	"fmt"
	"sync"
	"time"
)

// // nsync here
// func main() {

// 	counter("sheep")
// 	// sheep
// 	// sheep
// 	// sheep
// 	// sheep
// 	counter("cow")
// }

// by default main is a go routine
// func main() {
// 	// conter("sheel") runs concurrently
// 	go counter("sheep")

// 	// if make this a go routine too,
// 	// terminate (0) the program
// 	counter("cow")

// }

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go counter("sheep")

	go func() {
		counter("cow")
		wg.Done()
	}()

	wg.Wait()
}
func counter(str interface{}) {

	for {
		fmt.Println(str)
		time.Sleep(time.Second)

	}
}
