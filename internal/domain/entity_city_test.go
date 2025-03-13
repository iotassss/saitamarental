package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarent/internal/domain"
)

func TestNewCity(t *testing.T) {
	id, err := domain.NewID(1)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	name, err := domain.NewCityName("さいたま市")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	code, err := domain.NewCityCode("11", "001")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	order, err := domain.NewCityOrder(1)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	stationIDs := []domain.ID{}
	rentList := []domain.Rent{}

	_, err = domain.NewCity(id, name, code, order, stationIDs, rentList)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
