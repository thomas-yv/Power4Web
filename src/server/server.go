package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/game", http.StatusSeeOther)
	})
	http.HandleFunc("/game", StartGame)

	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		colStr := r.URL.Query().Get("col")
		col, err := strconv.Atoi(colStr)
		if err != nil || col < 0 || col >= 7{
			http.Error(w, "Invalid column", http.StatusBadRequest)
			return
		}

		PlaceCoinLine(col)
		http.Redirect(w, r, "/game", http.StatusSeeOther)
	})

	fmt.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
