package functions

import "fmt"

func VariadicFunction(nums ...int) {
	fmt.Println("Variadic function called with:", nums)
}

func NamedReturnValues(a int, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return
}
