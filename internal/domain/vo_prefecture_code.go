package domain

import (
	"fmt"
	"strconv"
)

const (
	SaitamaPrefectureCode = "11"
	TokyoPrefectureCode   = "13"
)

type PrefectureCode string

func NewPrefectureCode(code string) (PrefectureCode, error) {
	// at the moment, only Saitama and Tokyo are supported
	if code != SaitamaPrefectureCode && code != TokyoPrefectureCode {
		return "", fmt.Errorf("invalid prefecture code")
	}

	if len(code) != 2 {
		return PrefectureCode(""), fmt.Errorf("prefectureCode must be 2 digits")
	}
	intPrefectureCode, err := strconv.Atoi(code)
	if err != nil {
		return PrefectureCode(""), fmt.Errorf("prefectureCode must be number")
	}
	if intPrefectureCode < 1 || intPrefectureCode > 47 {
		return PrefectureCode(""), fmt.Errorf("prefectureCode must be 1~47")
	}

	return PrefectureCode(code), nil
}

func (p PrefectureCode) String() string {
	return string(p)
}

func (p PrefectureCode) Equals(other PrefectureCode) bool {
	return p.String() == other.String()
}
