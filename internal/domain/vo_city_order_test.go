package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarent/internal/domain"
)

func TestNewCityOrder(t *testing.T) {
	_, err := domain.NewCityOrder(1)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
