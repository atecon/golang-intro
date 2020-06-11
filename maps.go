package main

import "fmt"

func getkeys(m map[string]int) []string {
	/* */
	keys := make([]string, len(m))
	i := 0
	for v := range m {
		keys[i] = v
		i++
	}
	return keys
}

func main() {
	m := make(map[string]int) // key (str), value(int)
	m["a"] = 122
	m["b"] = -1323
	fmt.Println(m)

	for i, v := range m {
		fmt.Printf("key = %s, value = %d\n", i, v)
	}

	// Accessing
	value, ok := m["a"]
	if ok {
		fmt.Println("Got", value)
	}

	// Retrieve all keys -- unsorted order
	keys := make([]string, len(m))
	i := 0
	for v := range m {
		keys[i] = v
		i++
	}
	fmt.Println(keys)

	fmt.Println(getkeys(m))
}
