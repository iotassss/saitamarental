package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewCityOrder_Valid(t *testing.T) {
	order, err := domain.NewCityOrder(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, order.Value())
}

func TestNewCityOrder_Invalid(t *testing.T) {
	_, err := domain.NewCityOrder(0)
	assert.Error(t, err)

	_, err = domain.NewCityOrder(-5)
	assert.Error(t, err)
}

func TestCityOrder_ComesBefore(t *testing.T) {
	order1, _ := domain.NewCityOrder(1)
	order2, _ := domain.NewCityOrder(2)

	assert.True(t, order1.ComesBefore(order2))
	assert.False(t, order2.ComesBefore(order1))
	assert.False(t, order1.ComesBefore(order1))
}
