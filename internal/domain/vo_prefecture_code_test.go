package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewPrefectureCode_Valid(t *testing.T) {
	code, err := domain.NewPrefectureCode("11") // 埼玉
	assert.NoError(t, err)
	assert.Equal(t, "11", code.String())

	code, err = domain.NewPrefectureCode("13") // 東京
	assert.NoError(t, err)
	assert.Equal(t, "13", code.String())
}
