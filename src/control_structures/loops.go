package control_structures

import "fmt"

func ForLoopExample() {
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}
}

func WhileLoopExample() {
	i := 0
	for i < 5 {
		fmt.Println("i:", i)
		i++
	}
}
