package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, World!")
		})
	http.ListenAndServe(":3000", nil)
}