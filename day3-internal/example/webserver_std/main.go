package main

import (
	"fmt"
	"net/http"
)

// START 1 OMIT
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", greet)

	http.ListenAndServe(":8080", mux)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World!")
}

// END 1 OMIT
