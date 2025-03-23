package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewLayout_Valid(t *testing.T) {
	layoutStr := "2LDK/3K"
	layout, err := domain.NewLayout(layoutStr)
	assert.NoError(t, err)
	assert.Equal(t, layoutStr, layout.String())
}

func TestLayout_IsLargerThan(t *testing.T) {
	small, _ := domain.NewLayout("1K/1DK")
	large, _ := domain.NewLayout("2LDK/3K")
	largest, _ := domain.NewLayout("3LDK/4K~")

	assert.True(t, large.IsLargerThan(small))
	assert.True(t, largest.IsLargerThan(small))
	assert.True(t, largest.IsLargerThan(large))

	assert.False(t, small.IsLargerThan(large))
	assert.False(t, large.IsLargerThan(largest))
	assert.False(t, small.IsLargerThan(small)) // 同じなら false
}
