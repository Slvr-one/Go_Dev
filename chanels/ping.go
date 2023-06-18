package main

// https://www.golang-book.com/books/intro/10
import (
	"fmt"
	"time"
)

// c can only be sent to (<-)
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		// fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}

/* Using a channel like this synchronizes the two goroutines.
When pinger attempts to send a message on the channel
it will wait until printer is ready to receive the message.
(this is known as blocking) */

func main() {
	var c chan string = make(chan string)

	go pinger(c) //print ping
	go ponger(c) //print pong
	go printer(c)

	// goroutine return immediately to the next line, don't wait for the function to complete.
	// This is why the call to the Scanln function has been included;
	// without it the program would exit before being given the opportunity to to anything
	var input string
	fmt.Scanln(&input)
}
