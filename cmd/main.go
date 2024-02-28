package main

import (
	"github.com/enkdress/go-templ/handlers"
	"github.com/enkdress/go-templ/views"
	"github.com/labstack/echo/v4"
)

func main() {

	// Create an instance for a new server
	e := echo.New()

	// Get the Index page
	indexPage := views.Page()

	encorianHandler := handlers.EncorianHandler{}

	// Routes declarations
	e.GET("/", func(c echo.Context) error {
		return indexPage.Render(c.Request().Context(), c.Response())
	})
	e.GET("/encorian/avg", encorianHandler.HandleFindAvg)
	e.POST("/encorian", encorianHandler.HandleAddEncorian)
	e.DELETE("/encorian/:id", encorianHandler.HandleDeleteEncorian)

	e.Logger.Fatal(e.Start(":3000"))
}
