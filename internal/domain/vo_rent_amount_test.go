package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
)

func TestNewRentAmount(t *testing.T) {
	_, err := domain.NewRentAmount("12.3")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
