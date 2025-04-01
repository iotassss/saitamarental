package domain

import "errors"

var (
	ErrEntityNotFound = errors.New("entity not found")
	ErrIDAlreadySet   = errors.New("ID is already set and cannot be changed")
)
