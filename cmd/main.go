package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/enkdress/go-templ/handlers"
	"github.com/enkdress/go-templ/views"
)

func main() {
	component := views.Page()
	mux := http.NewServeMux()

	mux.Handle("/", templ.Handler(component))
	mux.HandleFunc("/pizza", handlers.PizzaHandler)
	mux.HandleFunc("/pizza/{id}", handlers.PizzaHandler)
	mux.HandleFunc("/encorians/avg", handlers.EncoriansAverage)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", mux)
}
