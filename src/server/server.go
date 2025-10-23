package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Start() {

	tmpl := template.Must(template.ParseFiles("../client/index.html"))

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			Title: "PUISSANCE 4",
			Todos: []Todo{
				{Title: "Learn Go", Done: true},
				{Title: "Learn Templates", Done: false},
				{Title: "Build an App", Done: false},
			},
		}

		tmpl.Execute(w, data)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Puissance 4")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
