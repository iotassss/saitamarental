package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
)

func TestNewCityName(t *testing.T) {
	_, err := domain.NewCityName("さいたま市")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
