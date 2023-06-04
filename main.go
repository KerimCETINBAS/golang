package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	c2 := make(chan string)
	//go count("sheep", c)

	// for {

	// 	msg, open := <-c

	// 	if !open {
	// 		break
	// 	}

	// for msg := range c {

	// 	fmt.Println(msg)
	// }

	go func() {
		for {
			c <- "Every second"
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 second"
			time.Sleep(time.Second * 2)
		}
	}()

	for {

		// blocking as slowest one
		// fmt.Println(<-c)
		// fmt.Println(<-c2)

		select {
		case msg1 := <-c:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)

		}
	}
}

func count(str string, c chan string) {

	for {

		c <- str
		time.Sleep(time.Second)

	}
}
