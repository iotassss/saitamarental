package domain

import "fmt"

type ID int

func NewID(id int) (ID, error) {
	if id < 1 {
		return ID(0), fmt.Errorf("ID must be greater than 0")
	}

	return ID(id), nil
}
