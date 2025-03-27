package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestAverageRentAmount_NormalCases(t *testing.T) {
	r, err := domain.NewAverageRentAmount("10.0")
	assert.NoError(t, err)
	assert.Equal(t, "10.0", r.String())
	assert.Equal(t, 10.0, r.Number())

	upper, err := r.Get10PercentUpper()
	assert.NoError(t, err)
	assert.Equal(t, "11.0", upper.String())
	assert.Equal(t, 11.0, upper.Number())

	lower, err := r.Get10PercentLower()
	assert.NoError(t, err)
	assert.Equal(t, "9.0", lower.String())
	assert.Equal(t, 9.0, lower.Number())
}

func TestAverageRentAmount_Get10PercentUpper(t *testing.T) {
	r, err := domain.NewAverageRentAmount("10.0")
	assert.NoError(t, err)

	upper, err := r.Get10PercentUpper()
	assert.NoError(t, err)
	assert.Equal(t, "11.0", upper.String())
}

func TestAverageRentAmount_Get10PercentLower(t *testing.T) {
	r, err := domain.NewAverageRentAmount("10.0")
	assert.NoError(t, err)

	lower, err := r.Get10PercentLower()
	assert.NoError(t, err)
	assert.Equal(t, "9.0", lower.String())
}

func TestAverageRentAmount_Get10PercentUpper_WithDecimal(t *testing.T) {
	r, err := domain.NewAverageRentAmount("10.5")
	assert.NoError(t, err)

	upper, err := r.Get10PercentUpper()
	assert.NoError(t, err)
	assert.Equal(t, "11.6", upper.String())
}

func TestAverageRentAmount_Get10PercentLower_WithDecimal(t *testing.T) {
	r, err := domain.NewAverageRentAmount("10.5")
	assert.NoError(t, err)

	lower, err := r.Get10PercentLower()
	assert.NoError(t, err)
	assert.Equal(t, "9.5", lower.String())
}
