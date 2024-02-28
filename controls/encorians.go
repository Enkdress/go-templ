package controls

import (
	"net/url"
	"strconv"

	"github.com/enkdress/go-templ/constants"
	"github.com/enkdress/go-templ/models"
	"github.com/google/uuid"
)

func AddEncorian(form url.Values) {
	var encorian models.Encorian
	for key, value := range form {

		encorian.Id = uuid.NewString()
		if key == "name" {
			encorian.Name = value[0]
		}

		if key == "pizzaAmount" {
			pAmount, _ := strconv.Atoi(value[0])
			encorian.PizzaAmount = int8(pAmount)
		}
	}

	constants.Encorians = append(constants.Encorians, encorian)
}

func DeleteEncorian(id string) {
	var newEncorians []models.Encorian
	for _, encorian := range constants.Encorians {
		if encorian.Id != id {
			newEncorians = append(newEncorians, encorian)
		}
	}
	constants.Encorians = newEncorians
}

func FindPizzaAvg() int {
	var pizzaSum int
	for _, encorian := range constants.Encorians {
		pizzaSum = pizzaSum + int(encorian.PizzaAmount)
	}

	avg := pizzaSum / len(constants.Encorians)

	return avg
}
