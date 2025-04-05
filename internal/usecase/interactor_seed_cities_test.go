package usecase_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
	"github.com/iotassss/saitamarental/internal/repository/memory"
	"github.com/iotassss/saitamarental/internal/unittest"
	"github.com/iotassss/saitamarental/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestSeedCitiesInteractor_Execute_Valid(t *testing.T) {
	// モックのリポジトリを作成
	cityRepo := memory.NewInMemoryCityRepo()

	// モックのデータローダーを作成
	mockLoader := &MockCityDataLoader{
		cities: []*domain.City{
			unittest.Must(domain.NewCity(
				unittest.Must(domain.NewID(1)),
				unittest.Must(domain.NewCityName("さいたま市")),
				unittest.Must(domain.ParseCityCode("11101")),
				unittest.Must(domain.NewCityOrder(1)),
			)),
			unittest.Must(domain.NewCity(
				unittest.Must(domain.NewID(2)),
				unittest.Must(domain.NewCityName("川口市")),
				unittest.Must(domain.ParseCityCode("11201")),
				unittest.Must(domain.NewCityOrder(2)),
			)),
		},
	}

	// モックのプレゼンターを作成
	mockPresenter := &MockSeedCitiesPresenter{
		seededCount: 0,
	}

	// インタラクターを作成
	interactor := usecase.NewSeedCitiesInteractor()

	// Execute を実行
	err := interactor.Execute(mockLoader, cityRepo, mockPresenter)

	// エラーがないことを確認
	assert.NoError(t, err)

	// プレゼンターに渡されたデータが正しいことを確認
	assert.Equal(t, 2, mockPresenter.seededCount)

	// リポジトリに保存されたデータが正しいことを確認
	cities, err := cityRepo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, cities, 2)
	assert.Equal(t, "さいたま市", cities[0].Name().String())
	assert.Equal(t, "川口市", cities[1].Name().String())
	assert.Equal(t, "11101", cities[0].Code().String())
	assert.Equal(t, "11201", cities[1].Code().String())
	assert.Equal(t, 1, cities[0].Order().Value())
	assert.Equal(t, 2, cities[1].Order().Value())
}

type MockCityDataLoader struct {
	cities []*domain.City
}

func (m *MockCityDataLoader) Load() ([]*domain.City, error) {
	return m.cities, nil
}

type MockSeedCitiesPresenter struct {
	seededCount int
}

func (p *MockSeedCitiesPresenter) Present(outputData usecase.SeedCitiesOutputData) error {
	p.seededCount = outputData.SeededCount
	return nil
}

func (p *MockSeedCitiesPresenter) PresentError(err error) error {
	return err
}

// type SeedCitiesOutputData struct {
// 	SeededCount int
// }

// type SeedCitiesPresenter interface {
// 	Present(output SeedCitiesOutputData) error
// 	PresentError(err error) error
// }
