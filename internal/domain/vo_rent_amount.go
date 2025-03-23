package domain

import (
	"fmt"
	"math"
	"strconv"
)

type AverageRentAmount struct {
	amount      string
	floatAmount float64
}

func NewAverageRentAmount(amount string) (AverageRentAmount, error) {
	if amount == "" {
		return AverageRentAmount{}, fmt.Errorf("rent amount cannot be empty")
	}
	floatAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return AverageRentAmount{}, fmt.Errorf("rent amount must be a number")
	}
	if floatAmount*10 != float64(int(floatAmount*10)) {
		return AverageRentAmount{}, fmt.Errorf("rent amount must have one decimal place")
	}

	return AverageRentAmount{
		amount:      amount,
		floatAmount: floatAmount,
	}, nil
}

func (r AverageRentAmount) String() string  { return r.amount }
func (r AverageRentAmount) Number() float64 { return r.floatAmount }

func (r AverageRentAmount) Get10PercentUpper() (AverageRentAmount, error) {
	hundredTenPercent := math.Round(r.floatAmount*1.1*10) / 10
	return NewAverageRentAmount(fmt.Sprintf("%.1f", hundredTenPercent))
}

func (r AverageRentAmount) Get10PercentLower() (AverageRentAmount, error) {
	ninetyPercent := math.Round(r.floatAmount*0.9*10) / 10
	return NewAverageRentAmount(fmt.Sprintf("%.1f", ninetyPercent))
}
