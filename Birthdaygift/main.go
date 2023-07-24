package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "style.css")
	})

	// http.HandleFunc("/script.js", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "script.js")
	// })

	http.ListenAndServe(":4000", nil)
}
