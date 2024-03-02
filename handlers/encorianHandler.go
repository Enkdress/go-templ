package handlers

import (
	"net/http"

	"github.com/enkdress/go-templ/models"
	"github.com/enkdress/go-templ/views/components"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EncorianHandler struct {
	encorians []models.Encorian
	avg       int32
}

func (p *EncorianHandler) HandleAddEncorian(c echo.Context) error {
	var encorian models.Encorian
	err := c.Bind(&encorian)

	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	encorian.Id = uuid.NewString()
	p.encorians = append(p.encorians, encorian)
	p.findAvg()

	c.Response().Header().Add("HX-Trigger", "newEncorian")

	return components.EncoriansList(p.encorians, int(p.avg)).Render(c.Request().Context(), c.Response())
}

func (p *EncorianHandler) HandleGetEncorians(c echo.Context) error {
	return components.EncoriansList(p.encorians, int(p.avg)).Render(c.Request().Context(), c.Response())
}

func (p *EncorianHandler) HandleDeleteEncorian(c echo.Context) error {
	var newEncorians []models.Encorian
	encorianId := c.Param("id")

	for _, encorian := range p.encorians {
		if encorian.Id != encorianId {
			newEncorians = append(newEncorians, encorian)
		}
	}

	p.encorians = newEncorians

	return c.String(http.StatusNoContent, "")
}

func (p *EncorianHandler) findAvg() {
	var pizzaSum int

	for _, encorian := range p.encorians {
		pizzaSum = pizzaSum + int(encorian.PizzaAmount)
	}

	avg := pizzaSum / len(p.encorians)
	p.avg = int32(avg)
}
