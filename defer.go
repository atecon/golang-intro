package main

import (
	"fmt"
	"time"
)

func cleanup() {
	fmt.Println("Cleaning up")
}

func work(i int) {
	fmt.Println("Working...")
	time.Sleep(time.Second * time.Duration(i))
	fmt.Println("Work done...", i)
}

func main() {
	defer cleanup() // will be the last method being caylled before exit
	work(3)
}
