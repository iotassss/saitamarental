package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
)

func TestNewCityCode(t *testing.T) {
	_, err := domain.NewCityCode("11", "001")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
