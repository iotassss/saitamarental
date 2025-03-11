package domain

import (
	"fmt"
	"strconv"
)

type CityCode struct {
	prefectureCode   string
	municipalityCode string
	checkDigit       int
}

func NewCityCode(prefectureCode string, municipalityCode string) (CityCode, error) {
	// prefectureCode
	if len(prefectureCode) != 2 {
		return CityCode{}, fmt.Errorf("prefectureCode must be 2 digits")
	}
	intPrefectureCode, err := strconv.Atoi(prefectureCode)
	if err != nil {
		return CityCode{}, fmt.Errorf("prefectureCode must be number")
	}
	if intPrefectureCode < 1 || intPrefectureCode > 47 {
		return CityCode{}, fmt.Errorf("prefectureCode must be 1~47")
	}

	// municipalityCode
	if len(municipalityCode) != 3 {
		return CityCode{}, fmt.Errorf("municipalityCode must be 3 digits")
	}
	if _, err := strconv.Atoi(municipalityCode); err != nil {
		return CityCode{}, fmt.Errorf("municipalityCode must be number")
	}

	// Calculate check digit
	fullCode := prefectureCode + municipalityCode
	checkDigit, err := calculateCheckDigit(fullCode)
	if err != nil {
		return CityCode{}, err
	}

	return CityCode{
		prefectureCode:   prefectureCode,
		municipalityCode: municipalityCode,
		checkDigit:       checkDigit,
	}, nil
}

func calculateCheckDigit(code string) (int, error) {
	if len(code) != 5 {
		return 0, fmt.Errorf("code must be 5 digits")
	}

	weights := []int{6, 5, 4, 3, 2}
	sum := 0
	for i, weight := range weights {
		digit, err := strconv.Atoi(string(code[i]))
		if err != nil {
			return 0, fmt.Errorf("invalid digit in code")
		}
		sum += digit * weight
	}

	remainder := sum % 11
	checkDigit := 11 - remainder

	if checkDigit == 10 {
		checkDigit = 0
	} else if checkDigit == 11 {
		checkDigit = 5
	}

	return checkDigit, nil
}
