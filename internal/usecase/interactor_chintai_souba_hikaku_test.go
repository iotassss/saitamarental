package usecase_test

import (
	"fmt"
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
	"github.com/iotassss/saitamarental/internal/repository/memory"
	"github.com/iotassss/saitamarental/internal/unittest"
	"github.com/iotassss/saitamarental/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestExecute_Valid(t *testing.T) {
	// repository
	cityRepo := memory.NewInMemoryCityRepo()
	averageRentRepo := memory.NewInMemoryAverageRentRepo()

	// test data
	baseCity1 := unittest.Must(domain.NewCity(
		unittest.Must(domain.NewID(11)),
		unittest.Must(domain.NewCityName("千代田区")),
		unittest.Must(domain.ParseCityCode("13101")),
		unittest.Must(domain.NewCityOrder(1)),
	))
	baseAverageRent1 := unittest.Must(domain.NewAverageRent(
		unittest.Must(domain.NewID(11)),
		baseCity1.ID(),
		unittest.Must(domain.NewAverageRentAmount("12.3")),
		unittest.Must(domain.NewLayout("ワンルーム")),
	))
	city1 := unittest.Must(domain.NewCity(
		unittest.Must(domain.NewID(1)),
		unittest.Must(domain.NewCityName("さいたま市大宮区")),
		unittest.Must(domain.ParseCityCode("111031")),
		unittest.Must(domain.NewCityOrder(1)),
	))
	averageRent1 := unittest.Must(domain.NewAverageRent(
		unittest.Must(domain.NewID(1)),
		city1.ID(),
		unittest.Must(domain.NewAverageRentAmount("12.3")),
		unittest.Must(domain.NewLayout("1K/1DK")),
	))
	averageRent2 := unittest.Must(domain.NewAverageRent(
		unittest.Must(domain.NewID(2)),
		city1.ID(),
		unittest.Must(domain.NewAverageRentAmount("45.6")),
		unittest.Must(domain.NewLayout("1K/1DK")),
	))
	city2 := unittest.Must(domain.NewCity(
		unittest.Must(domain.NewID(2)),
		unittest.Must(domain.NewCityName("川口市")),
		unittest.Must(domain.ParseCityCode("112038")),
		unittest.Must(domain.NewCityOrder(2)),
	))
	averageRent3 := unittest.Must(domain.NewAverageRent(
		unittest.Must(domain.NewID(3)),
		city2.ID(),
		unittest.Must(domain.NewAverageRentAmount("7.8")),
		unittest.Must(domain.NewLayout("2LDK/3K")),
	))
	averageRent4 := unittest.Must(domain.NewAverageRent(
		unittest.Must(domain.NewID(4)),
		city2.ID(),
		unittest.Must(domain.NewAverageRentAmount("9.0")),
		unittest.Must(domain.NewLayout("3LDK/4K~")),
	))

	// save test data
	cityRepo.Save(baseCity1)
	averageRentRepo.Save(baseAverageRent1)
	cityRepo.Save(city1)
	cityRepo.Save(city2)
	averageRentRepo.Save(averageRent1)
	averageRentRepo.Save(averageRent2)
	averageRentRepo.Save(averageRent3)
	averageRentRepo.Save(averageRent4)

	// input test data
	input := usecase.ChintaiSoubaHikakuInputData{
		CityCode: "13101",
		Layout:   "ワンルーム",
		Criteria: "larger",
	}

	// execute
	interactor := usecase.NewChintaiSoubaHikakuInteractor()
	presenter := &presenter{}
	err := interactor.Execute(input, presenter, cityRepo, averageRentRepo)

	assert.NoError(t, err)
}

type presenter struct{}

func (p *presenter) Present(outputData usecase.ChintaiSoubaHikakuOutputData) error {
	if outputData.BaseTokyoCity.Name != "千代田区" {
		return fmt.Errorf("unexpected BaseTokyoCity.Name: %s", outputData.BaseTokyoCity.Name)
	}
	if outputData.BaseTokyoCity.Layout != "ワンルーム" {
		return fmt.Errorf("unexpected BaseTokyoCity.Layout: %s", outputData.BaseTokyoCity.Layout)
	}
	if outputData.BaseTokyoCity.AverageRent != "12.3" {
		return fmt.Errorf("unexpected BaseTokyoCity.AverageRent: %s", outputData.BaseTokyoCity.AverageRent)
	}
	if outputData.BaseTokyoCity.CityCode != "131016" {
		return fmt.Errorf("unexpected BaseTokyoCity.CityCode: %s", outputData.BaseTokyoCity.CityCode)
	}

	if len(outputData.SaitamaCities) != 1 {
		return fmt.Errorf("unexpected SaitamaCities length: %d", len(outputData.SaitamaCities))
	}
	if outputData.SaitamaCities[0].Name != "さいたま市大宮区" {
		return fmt.Errorf("unexpected SaitamaCities[0].Name: %s", outputData.SaitamaCities[0].Name)
	}
	if outputData.SaitamaCities[0].Layout != "1K/1DK" {
		return fmt.Errorf("unexpected SaitamaCities[0].Layout: %s", outputData.SaitamaCities[0].Layout)
	}
	if outputData.SaitamaCities[0].AverageRent != "12.3" {
		return fmt.Errorf("unexpected SaitamaCities[0].AverageRent: %s", outputData.SaitamaCities[0].AverageRent)
	}
	if outputData.SaitamaCities[0].CityCode != "111031" {
		return fmt.Errorf("unexpected SaitamaCities[0].CityCode: %s", outputData.SaitamaCities[0].CityCode)
	}

	return nil
}

func (p *presenter) PresentError(err error) error {
	return err
}
