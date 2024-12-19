package variables

import (
	"clean-go/src/models"
	"clean-go/src/routes"
	"html/template"
	"net/http"
)

func VariableDeclaration(w http.ResponseWriter) {
	example := models.Layout{
		Title: "Variable Declaration",
		Body: `Variables in Go are declared using the 'var' keyword.

The most common syntax for declaring a variable is: var name type = value. 
The type of the variable is optional, but it is recommended to specify it to avoid any confusion. 
Here are the list of types that can be used in Go:

- bool (true/false)
- string (text)
- int (integer)
- int8, int16, int32, int64 (signed integers)
- uint (unsigned integer)
- uint8, uint16, uint32, uint64 (unsigned integers)
- float32, float64 (floating point numbers)
- complex64, complex128 (complex numbers)
- byte (alias for uint8)
- rune (alias for int32)
- uintptr (unsigned integer to store the address of a pointer)
- array (fixed-size collection of items)
- slice (dynamic-size collection of items)
- map (key-value pairs)
- struct (collection of fields)
- interface (set of methods)
- function (block of code)
- pointer (memory address of a variable)
- channel (communication between goroutines)
- nil (zero value for pointers, interfaces, maps, slices, channels, and functions)
- const (constant value)
- iota (enumerator)
- var (variable declaration)
- type (type declaration)
- defer (deferred function call)
- go (goroutine)
- select (selects a send/receive operation on channels)
- range (iterates over elements in a variety of data structures)
`,
		Examples: []string{
			"var name bool = true",
			"var name string = 'Hello, World!'",
			"var name int = 42",
			"var name float64 = 3.14",
			"var name complex128 = 1 + 2i",
			"var name byte = 'a'",
			"var name rune = '☺'",
			"var name uintptr = 0x123456",
			"var name [3]int = [3]int{1, 2, 3}",
			"var name []int = []int{1, 2, 3}",
			"var name map[string]int = map[string]int{'one': 1, 'two': 2, 'three': 3}",
			"var name struct { x int; y int } = struct { x int; y int }{1, 2}",
			"var name interface{} = 42",
			"var name func() = func() { return }",
			"var name *int = &x",
			"var name chan int = make(chan int)",
			"var name nil",
			"const name = 42",
			"const name = iota",
			"var name",
			"type name int",
			"defer func() { return }()",
			"go func() { return }()",
			"select { case <-ch: }",
			"for i := range arr { }",
		},
		Navigation: routes.GetMainNavigation(),
	}

	tmpl := template.Must(template.ParseFiles("templates/simple_layout.html"))
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, example)
}
