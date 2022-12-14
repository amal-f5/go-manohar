package main

import (
	"fmt"
	"log"
	"net/http"
	"unit.nginx.org/go"
)
// To get this up working with Eve, add unit.nginx.org/go in the imports

func main() {
	log.Printf("Starting server at port 8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home")
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Eve!")
	})
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		return
	})
	
	// Here, change http.ListenAndServer -> unit.ListenAndServe
	if err := unit.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
