package main

import "fmt"

func incr(i int) int {
	return i + 1
}

func sq(i int) int {
	return i * i
}

func pickMath(i int) func(int) int {
	// 'func(int) int' is a return value
	if i == 1 {
		return incr
	}

	return sq
}

func main() {
	f1 := pickMath(1)
	f2 := pickMath(2)

	fmt.Println(f1(1336))
	fmt.Println(f2(4))
}
