package main

import (
	"fmt"
	"time"
)

var c = make(chan string)

func pinger() {
	for {
		c <- "ping" // send to the channel
		time.Sleep(1 * time.Second)
	}
}

func printer() {
	for {
		msg := <-c // reads from channel -- blocks the chanel + waits for new stream
		fmt.Println("got:", msg)
	}
}

func main() {
	go pinger()
	go printer()

	time.Sleep(5 * time.Second)
}
