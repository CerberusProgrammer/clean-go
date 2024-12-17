package variables

import (
	"clean-go/src/models"
	"html/template"
	"net/http"
)

func VariableDeclaration(w http.ResponseWriter) {
	example := models.SimpleLayoutStructure{
		Title: "Variable Declaration",
		Body: `Variables are declared using the var keyword. 
The type of the variable is optional, but it is recommended to specify the type to avoid any confusion.

Syntax:
var variable_name type = value
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
