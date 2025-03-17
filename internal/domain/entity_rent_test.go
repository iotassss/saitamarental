package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
)

func TestNewRent(t *testing.T) {
	id, err := domain.NewID(1)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	amount, err := domain.NewRentAmount("12.3")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	layout, err := domain.NewLayout("2LDK/3K")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	_, err = domain.NewRent(id, amount, layout)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
