package domain

import (
	"fmt"
	"strconv"
)

type CityCode struct {
	prefectureCode   PrefectureCode
	municipalityCode string
	checkDigit       int
}

func (c CityCode) PrefectureCode() PrefectureCode { return c.prefectureCode }

func ParseCityCode(code string) (CityCode, error) {
	if len(code) != 5 && len(code) != 6 {
		return CityCode{}, fmt.Errorf("code must be 5 or 6 digits")
	}

	prefectureCodeStr := code[:2] // 01-47
	municipalityCode := code[2:5] // 001-999

	prefectureCode, err := NewPrefectureCode(prefectureCodeStr)
	if err != nil {
		return CityCode{}, err
	}

	cityCode, err := NewCityCode(prefectureCode, municipalityCode)
	if err != nil {
		return CityCode{}, err
	}

	return cityCode, nil
}

func NewCityCode(prefectureCode PrefectureCode, municipalityCode string) (CityCode, error) {
	// municipalityCode
	if len(municipalityCode) != 3 {
		return CityCode{}, fmt.Errorf("municipalityCode must be 3 digits")
	}
	if _, err := strconv.Atoi(municipalityCode); err != nil {
		return CityCode{}, fmt.Errorf("municipalityCode must be number")
	}

	// Calculate check digit
	fullCode := prefectureCode.String() + municipalityCode
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

func (c CityCode) String() string {
	return c.prefectureCode.String() + c.municipalityCode + strconv.Itoa(c.checkDigit)
}

func (c CityCode) Equals(other CityCode) bool {
	return c.String() == other.String()
}
