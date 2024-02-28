package models

type Encorian struct {
	Id          string
	Name        string `json:"name" form:"name"`
	PizzaAmount int8   `json:"pizzaAmount" form:"pizzaAmount"`
}
