package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestExtractCityIDs(t *testing.T) {
	id1, _ := domain.NewID(1)
	id2, _ := domain.NewID(2)

	city1, _ := domain.NewCity(id1, mustCityName("市A"), mustCityCode("11", "001"), mustCityOrder(1))
	city2, _ := domain.NewCity(id2, mustCityName("市B"), mustCityCode("11", "002"), mustCityOrder(2))

	ids := domain.ExtractCityIDs([]*domain.City{city1, city2, city1}) // 重複含む

	assert.Len(t, ids, 2)
	assert.Contains(t, ids, id1)
	assert.Contains(t, ids, id2)
}

// 補助関数
func mustCityName(s string) domain.CityName {
	v, _ := domain.NewCityName(s)
	return v
}

func mustCityCode(prefStr, city string) domain.CityCode {
	pref, _ := domain.NewPrefectureCode(prefStr)
	v, _ := domain.NewCityCode(pref, city)
	return v
}

func mustCityOrder(n int) domain.CityOrder {
	v, _ := domain.NewCityOrder(n)
	return v
}
