package domain

import "fmt"

type TravelTime int

func NewTravelTime(time int) (TravelTime, error) {
	if time < 0 {
		return TravelTime(0), fmt.Errorf("travel time must be greater than or equal to 0")
	}

	return TravelTime(time), nil
}

func (t TravelTime) Value() int {
	return int(t)
}

func (t TravelTime) GetFormattedValue() string {
	return fmt.Sprintf("%d分", t)
}
