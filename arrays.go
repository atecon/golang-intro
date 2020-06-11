package main // package name 'main'

import "fmt"

func main() {
	var S []string // string array
	S = []string{"Hello", "World"}
	fmt.Println(S)
	S = append(S, "!", "Go is cool")
	fmt.Println(S)
	Snew := append(S, "!")

	fmt.Println(Snew)
	// access via index
	fmt.Println(Snew[0])           // first element
	fmt.Println(Snew[len(Snew)-1]) // last element
	fmt.Println(Snew[0:3])
	fmt.Println(Snew[2:])
}
