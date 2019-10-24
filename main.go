// package Backend_Challenge
package main

import (
	"fmt"
	"league/handler"
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		handler.EchoHandler(w, r)
	})
	http.HandleFunc("/invert", func(w http.ResponseWriter, r *http.Request) {
		handler.InvertHandler(w, r)
	})
	http.HandleFunc("/flatten", func(w http.ResponseWriter, r *http.Request) {
		handler.FlattenHandler(w, r)
	})
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		handler.SumHandler(w, r)
	})
	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		handler.MultiplyHandler(w, r)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Run server error: ", err)
	}
}
