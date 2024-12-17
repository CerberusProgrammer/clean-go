package functions

import "fmt"

func BasicFunction() {
	fmt.Println("This is a basic function.")
}

func FunctionWithParameters(a int, b int) int {
	return a + b
}

func FunctionWithReturn() string {
	return "Hello from a function with return!"
}
