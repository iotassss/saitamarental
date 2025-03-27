package usecase

import (
	"github.com/iotassss/saitamarental/internal/domain"
)

type ChintaiSoubaHikakuInputData struct {
	CityCode string
	Layout   string
	Criteria string
}

type ChintaiSoubaHikakuOutputData struct {
	BaseTokyoCity struct {
		Name        string
		CityCode    string
		Layout      string
		AverageRent string
	}
	SaitamaCities []struct {
		Name        string
		CityCode    string
		Layout      string
		AverageRent string
	}
}

type ChintaiSoubaHikakuPresenter interface {
	Present(outputData ChintaiSoubaHikakuOutputData) error
	PresentError(err error) error
}

type ChintaiSoubaHikakuUsecase interface {
	Execute(
		input ChintaiSoubaHikakuInputData,
		presenter ChintaiSoubaHikakuPresenter,
		cityRepo domain.CityRepository,
		averageRentRepo domain.AverageRentRepository,
	) error
}
