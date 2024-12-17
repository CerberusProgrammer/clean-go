package data_structures

import "fmt"

func MapExample() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2

	fmt.Println("Map:", m)
}
