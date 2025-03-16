package domain

import (
	"fmt"
	"strconv"
)

type RentAmount struct {
	amount      string
	floatAmount float64
}

func NewRentAmount(amount string) (RentAmount, error) {
	if amount == "" {
		return RentAmount{}, fmt.Errorf("rent amount cannot be empty")
	}
	floatAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return RentAmount{}, fmt.Errorf("rent amount must be a number")
	}
	if floatAmount*10 != float64(int(floatAmount*10)) {
		return RentAmount{}, fmt.Errorf("rent amount must have one decimal place")
	}

	return RentAmount{
		amount:      amount,
		floatAmount: floatAmount,
	}, nil
}

func (r RentAmount) String() string {
	return r.amount
}

func (r RentAmount) Number() float64 {
	return r.floatAmount
}
