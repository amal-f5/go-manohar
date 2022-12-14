package main

import (
	"fmt"
	"log"
	"net/http"
	"unit.nginx.org/go"
)

func main() {
	log.Printf("Starting server at port 8080")
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Home")
// 	})
// 	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Eve!")
// 	})
// 	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
// 		return
// 	})
	if err := unit.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
		io.WriteString(w, "Hello, Go demo app!")
	}
}
