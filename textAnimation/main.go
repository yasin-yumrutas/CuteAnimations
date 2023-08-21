// package main

// import (
// 	"net/http"
// )

//	func main() {
//		http.Handle("/", http.FileServer(http.Dir(".")))
//		http.ListenAndServe(":4011", nil)
//	}
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/index.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.js")
	})

	http.HandleFunc("/index.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.css")
	})

	http.ListenAndServe(":4012", nil)
}
