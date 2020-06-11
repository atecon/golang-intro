package main // package name 'main'

import "fmt"

func main() {
	var s string
	s = "Hello, world!"
	a := "Hello, another world!"
	i := 42
	b := s + a // concatenate strings

	fmt.Println(s, i)
	fmt.Println(a)
	fmt.Println(b)

	i += 8
	fmt.Println(i)
}
