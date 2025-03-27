package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewAverageRent(t *testing.T) {
	id, err := domain.NewID(1)
	assert.NoError(t, err)
	cityID, err := domain.NewID(1)
	assert.NoError(t, err)
	amount, err := domain.NewAverageRentAmount("10.0")
	assert.NoError(t, err)
	layout := domain.Layout("1LDK")

	rent, err := domain.NewAverageRent(id, cityID, amount, layout)
	assert.NoError(t, err)
	assert.Equal(t, id, rent.ID())
	assert.Equal(t, cityID, rent.CityID())
	assert.Equal(t, amount, rent.Amount())
	assert.Equal(t, layout, rent.Layout())
}

func TestAverageRent_Get10PercentUpper(t *testing.T) {
	amount, _ := domain.NewAverageRentAmount("10.0")
	id, _ := domain.NewID(1)
	cityID, _ := domain.NewID(1)
	rent, _ := domain.NewAverageRent(id, cityID, amount, domain.Layout("1LDK"))

	upper, err := rent.Get10PercentUpper()
	assert.NoError(t, err)
	assert.Equal(t, "11.0", upper.String())
}

func TestAverageRent_Get10PercentLower(t *testing.T) {
	amount, _ := domain.NewAverageRentAmount("10.0")
	id, _ := domain.NewID(1)
	cityID, _ := domain.NewID(1)
	rent, _ := domain.NewAverageRent(id, cityID, amount, domain.Layout("1LDK"))

	lower, err := rent.Get10PercentLower()
	assert.NoError(t, err)
	assert.Equal(t, "9.0", lower.String())
}

func TestAverageRent_IsLargerLayoutThan(t *testing.T) {
	layout1 := domain.Layout(domain.Layout1K1DK)
	layout2 := domain.Layout(domain.LayoutOneRoom)
	id1, _ := domain.NewID(1)
	id2, _ := domain.NewID(2)
	cityID1, _ := domain.NewID(1)
	cityID2, _ := domain.NewID(2)

	rent1, _ := domain.NewAverageRent(id1, cityID1, mustRentAmount("11.0"), layout1)
	rent2, _ := domain.NewAverageRent(id2, cityID2, mustRentAmount("10.0"), layout2)

	assert.True(t, rent1.IsLargerLayoutThan(rent2))
	assert.False(t, rent2.IsLargerLayoutThan(rent1))
}

// テスト補助関数
func mustRentAmount(s string) domain.AverageRentAmount {
	a, err := domain.NewAverageRentAmount(s)
	if err != nil {
		panic(err)
	}
	return a
}
