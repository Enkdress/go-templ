package main

import (
	"github.com/enkdress/go-templ/handlers"
	"github.com/enkdress/go-templ/views"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	indexPage := views.Page()
	encorianHandler := handlers.EncorianHandler{}

	e.GET("/", func(c echo.Context) error {
		return indexPage.Render(c.Request().Context(), c.Response())
	})
	e.GET("/encorian/avg", encorianHandler.HandleFindAvg)
	e.POST("/encorian", encorianHandler.HandleAddEncorian)
	e.DELETE("/encorian/:id", encorianHandler.HandleDeleteEncorian)

	e.Logger.Fatal(e.Start(":3000"))
}
