package main

import (
	"clean-go/src/variables"
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/css", fs))

	http.HandleFunc("/variables/declarations", declarationsHandler)
	http.HandleFunc("/variables/types", typesHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func declarationsHandler(w http.ResponseWriter, r *http.Request) {
	variables.VariableDeclaration(w)
}

func typesHandler(w http.ResponseWriter, r *http.Request) {
	variables.VariableTypes(w)
}
