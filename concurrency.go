package main

import (
	"fmt"
	"time"
)

func work(i int) {
	fmt.Println("Working...")
	time.Sleep(time.Second * time.Duration(i))
	fmt.Println("Work done...", i)
}

func main() {
	for i := 1; i < 5; i++ {
		go work(i) // initialize a thread in the background
	}
	fmt.Println("Hello world")
	time.Sleep(time.Second * time.Duration(7))
}
