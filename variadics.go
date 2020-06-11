package main

import "fmt"

func printer(sl ...string) {
	// Receive separate string values as a slice
	// for iterating over the slice using 'range'
	for _, s := range sl {
		fmt.Println(s)
	}
}

func main() {
	// Directly pass separate values
	printer("Hello", "World")

	sl := []string{"Hello", "World"}
	// 'sl...' will expand the string slice into separate values
	printer(sl...)
	// Same as:
	printer(sl[0], sl[1])
}
