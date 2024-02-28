package handlers

import (
	"net/http"

	"github.com/enkdress/go-templ/controls"
	"github.com/enkdress/go-templ/views/components"
)

func EncoriansAverage(w http.ResponseWriter, r *http.Request) {
	avg := controls.FindPizzaAvg()
	components.EncoriansAvg(avg).Render(r.Context(), w)
}
