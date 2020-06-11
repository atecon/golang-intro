package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println("Iteration:", i)
	}

	// while loop style
	i := 0
	for i < 3 {
		fmt.Println("Iteration:", i)
		i++
	}

	// foreach loop
	S := []string{"Hello", "World", "!"}

	for i, s := range S {
		//fmt.Println("Index:", i, s)
		fmt.Printf("Index %d: %s", i, s)
	}
}
