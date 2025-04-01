package domain

import "fmt"

type ID int

func NewID(id int) (ID, error) {
	if id < 0 {
		return ID(0), fmt.Errorf("ID must be non-negative")
	}

	return ID(id), nil
}
