package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
)

func TestNewStation(t *testing.T) {
	id, err := domain.NewID(1)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	toTokyo, err := domain.NewTravelTime(37)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	toShinjuku, err := domain.NewTravelTime(1234)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	_, err = domain.NewStation(id, toTokyo, toShinjuku)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
