package variables

import (
	"html/template"
	"net/http"
)

type VariableExample struct {
	Title    string   `json:"title"`
	Body     string   `json:"body"`
	Examples []string `json:"examples"`
}

func VariableDeclaration(w http.ResponseWriter) {
	example := VariableExample{
		Title: "Variable Declaration",
		Body:  "Variables are declared using the var keyword. The type of the variable is optional, but it is recommended to specify the type to avoid any confusion.",
		Examples: []string{
			"var str string = \"Hello, Go!\"",
			"var num int = 42",
			"var flag bool = true",
			"var fnum float64 = 3.14",
		},
	}

	tmpl := template.Must(template.ParseFiles("templates/declarations.html"))
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, example)
}
