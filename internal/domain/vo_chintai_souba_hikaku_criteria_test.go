package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewChintaiSoubaHikakuCriteria_Valid(t *testing.T) {
	criteriaStrs := []string{
		"larger",
		"closer_to_shinjuku",
		"closer_to_tokyo",
		"larger_closer_to_shinjuku",
		"larger_closer_to_tokyo",
	}

	for _, s := range criteriaStrs {
		criteria, err := domain.NewChintaiSoubaHikakuCriteria(s)
		assert.NoError(t, err)
		assert.Equal(t, s, criteria.String())
	}
}
