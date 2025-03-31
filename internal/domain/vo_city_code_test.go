package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewCityCode(t *testing.T) {
	pref, err := domain.NewPrefectureCode("11")
	assert.NoError(t, err)

	code, err := domain.NewCityCode(pref, "001")
	assert.NoError(t, err)
	assert.Equal(t, "11", code.String()[:2])
	assert.Equal(t, 5, len(code.String())) // "11" + "001"
}

func TestCityCode_StringFormat(t *testing.T) {
	pref, _ := domain.NewPrefectureCode("11")
	code, _ := domain.NewCityCode(pref, "001")

	str := code.String()
	assert.Len(t, str, 5)
	assert.Equal(t, "11", str[:2])
	assert.Equal(t, "001", str[2:5])
}

func TestParseCityCode_Valid(t *testing.T) {
	original, _ := domain.NewCityCode(mustPrefectureCode("11"), "001")
	parsed, err := domain.ParseCityCode(original.String())

	assert.NoError(t, err)
	assert.Equal(t, original.String(), parsed.String())
}

// テスト補助関数（NewPrefectureCode 成功前提）
func mustPrefectureCode(code string) domain.PrefectureCode {
	pc, err := domain.NewPrefectureCode(code)
	if err != nil {
		panic(err)
	}
	return pc
}
