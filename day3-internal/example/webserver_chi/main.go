package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mercadolibre/chi/middleware"
)

// START 1 OMIT
func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Get("/", greet)

	http.ListenAndServe(":8080", r)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World!")
}

// END 1 OMIT
