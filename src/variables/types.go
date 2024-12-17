package variables

import (
	"fmt"
	"net/http"
)

func VariableTypes(w http.ResponseWriter) {
	var str string = "Hello, Go!"
	var num int = 42
	var flag bool = true
	var fnum float64 = 3.14

	fmt.Fprintf(w, "String: %s\n", str)
	fmt.Fprintf(w, "Integer: %d\n", num)
	fmt.Fprintf(w, "Boolean: %t\n", flag)
	fmt.Fprintf(w, "Float: %f\n", fnum)
}
