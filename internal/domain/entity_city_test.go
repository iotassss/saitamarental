package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewCity(t *testing.T) {
	id, _ := domain.NewID(1)
	name, _ := domain.NewCityName("さいたま市")
	code, _ := domain.NewCityCode("11", "001")
	order, _ := domain.NewCityOrder(1)

	city, err := domain.NewCity(id, name, code, order)
	assert.NoError(t, err)
	assert.Equal(t, id, city.ID())
	assert.Equal(t, name, city.Name())
	assert.Equal(t, code, city.Code())
	assert.Equal(t, order, city.Order())
}

func TestCity_ComesBefore(t *testing.T) {
	id1, _ := domain.NewID(1)
	id2, _ := domain.NewID(2)
	name1, _ := domain.NewCityName("市A")
	name2, _ := domain.NewCityName("市B")
	code1, _ := domain.NewCityCode("11", "001")
	code2, _ := domain.NewCityCode("11", "002")
	order1, _ := domain.NewCityOrder(1)
	order2, _ := domain.NewCityOrder(2)

	city1, _ := domain.NewCity(id1, name1, code1, order1)
	city2, _ := domain.NewCity(id2, name2, code2, order2)

	assert.True(t, city1.ComesBefore(city2))
	assert.False(t, city2.ComesBefore(city1))
	assert.False(t, city1.ComesBefore(city1))
}
