package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/enkdress/go-templ/views"
	"github.com/enkdress/go-templ/views/components"
)

func deleteEncorian(id string) {
	var newEncorians []Encorian
	for _, encorian := range encorians {
		if encorian.Id != id {
			newEncorians = append(newEncorians, encorian)
		}
	}
	encorians = newEncorians
}

func findPizzaAvg() int {
	var pizzaSum int
	for _, encorian := range encorians {
		pizzaSum = pizzaSum + int(encorian.PizzaAmount)
	}

	avg := pizzaSum / len(encorians)

	return avg
}

func pizzaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		addEncorian(r.Form)
		w.Header().Add("HX-Trigger", "newEncorian")

		components.EncoriansList(encorians).Render(r.Context(), w)
	} else if r.Method == "DELETE" {
		encorianId := r.PathValue("id")
		deleteEncorian(encorianId)
		w.WriteHeader(200)
	}

}

func encoriansAverage(w http.ResponseWriter, r *http.Request) {
	avg := findPizzaAvg()
	components.EncoriansAvg(avg).Render(r.Context(), w)
}

func main() {
	component := views.Page()
	mux := http.NewServeMux()
	mux.Handle("/", templ.Handler(component))
	mux.HandleFunc("/pizza", pizzaHandler)
	mux.HandleFunc("/pizza/{id}", pizzaHandler)
	mux.HandleFunc("/encorians/avg", encoriansAverage)
	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", mux)
}
