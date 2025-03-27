package domain

import (
	"fmt"
)

type CityOrder int

// doesn't check if the order is unique
func NewCityOrder(order int) (CityOrder, error) {
	if order < 1 {
		return CityOrder(0), fmt.Errorf("city order must be greater than 0")
	}

	return CityOrder(order), nil
}

func (c CityOrder) Value() int {
	return int(c)
}

func (c CityOrder) ComesBefore(order CityOrder) bool {
	return c.Value() < order.Value()
}
