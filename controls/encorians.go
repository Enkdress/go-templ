package controls

import (
	"net/url"
	"strconv"

	"github.com/google/uuid"
)

func AddEncorian(form url.Values) {
	var encorian Encorian
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

	encorians = append(encorians, encorian)
}
