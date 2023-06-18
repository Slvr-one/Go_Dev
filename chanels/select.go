package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// setting a value for c1, c2 (chanels of type str - async, )
	go func() {
		for {
			time.Sleep(time.Second * 2)
			c1 <- "from 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 3)
			c2 <- "from 2"
		}
	}()

	go func() {
		for {
			// select picks first channel thats ready and receives from it (or sends to it) if both ready than random.
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			// implement a timeout:
			case <-time.After(time.Second * 4):
				// creates a channel and after the given duration will send the current time on it.
				fmt.Println("timeout")
			default: //happens immediately if none of the channels are ready.
				fmt.Println("no chan is ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
