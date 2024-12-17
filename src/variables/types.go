package variables

import (
	"clean-go/src/models"
	"html/template"
	"net/http"
)

func VariableTypes(w http.ResponseWriter) {
	example := models.SimpleLayoutStructure{
		Title: "Variable Types",
		Body: `Go is a statically typed language, which means that the type of a variable is known at compile time.

The type of a variable determines the values that the variable can hold and the operations that can be performed on it.

Go has the following basic types:

1. bool
2. string
3. int, int8, int16, int32, int64
4. uint, uint8, uint16, uint32, uint64, uintptr
5. byte (alias for uint8)
6. rune (alias for int32)
7. float32, float64
8. complex64, complex128
`,
		Examples: []string{
			"var str string = \"Hello, Go!\"",
			"var num int = 42",
			"var flag bool = true",
			"var fnum float64 = 3.14",
		},
	}

	tmpl := template.Must(template.ParseFiles("templates/simple_layout.html"))
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, example)
}
