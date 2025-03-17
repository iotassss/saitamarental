package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
)

func TestNewTravelTime(t *testing.T) {
	_, err := domain.NewTravelTime(37)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
