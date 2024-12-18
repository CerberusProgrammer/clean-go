package main

import (
	"clean-go/src/config"
	"clean-go/src/pages/variables"
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir(config.StaticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, config.TemplatesDir+"/home_layout.html")
	})

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
