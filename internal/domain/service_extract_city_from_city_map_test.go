package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
	"github.com/iotassss/saitamarental/internal/unittest"
)

func TestExtractCityIDsFromCityMap(t *testing.T) {
	id1 := unittest.Must(domain.NewID(1))
	id2 := unittest.Must(domain.NewID(2))
	name1 := unittest.Must(domain.NewCityName("国立市"))
	name2 := unittest.Must(domain.NewCityName("三郷市"))
	code1 := unittest.Must(domain.NewCityCode("13", "215"))
	code2 := unittest.Must(domain.NewCityCode("11", "237"))
	order1 := unittest.Must(domain.NewCityOrder(1))
	order2 := unittest.Must(domain.NewCityOrder(2))

	city1 := unittest.Must(domain.NewCity(id1, name1, code1, order1))
	city2 := unittest.Must(domain.NewCity(id2, name2, code2, order2))

	cities := map[domain.ID]*domain.City{
		id1: city1,
		id2: city2,
	}

	ids := domain.ExtractCityIDsFromCityMap(cities)

	if len(ids) != 2 {
		t.Fatalf("expected 2 IDs, got %d", len(ids))
	}

	idMap := make(map[domain.ID]bool)
	for _, id := range ids {
		idMap[id] = true
	}

	if !idMap[id1] || !idMap[id2] {
		t.Errorf("missing expected IDs in result")
	}
}
