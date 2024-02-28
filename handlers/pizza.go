package handlers

import (
	"net/http"

	"github.com/enkdress/go-templ/constants"
	"github.com/enkdress/go-templ/controls"
	"github.com/enkdress/go-templ/views/components"
)

func PizzaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		controls.AddEncorian(r.Form)
		w.Header().Add("HX-Trigger", "newEncorian")

		components.EncoriansList(constants.Encorians).Render(r.Context(), w)
	} else if r.Method == "DELETE" {
		encorianId := r.PathValue("id")
		controls.DeleteEncorian(encorianId)
		w.WriteHeader(200)
	}

}
