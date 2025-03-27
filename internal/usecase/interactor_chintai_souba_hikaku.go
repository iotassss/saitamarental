package usecase

import (
	"github.com/iotassss/saitamarental/internal/domain"
)

type ChintaiSoubaHikakuInteractor struct {
	filters map[domain.ChintaiSoubaHikakuCriteria]domain.ChintaiSoubaHikakuFilter
}

func NewChintaiSoubaHikakuInteractor() *ChintaiSoubaHikakuInteractor {
	return &ChintaiSoubaHikakuInteractor{
		filters: map[domain.ChintaiSoubaHikakuCriteria]domain.ChintaiSoubaHikakuFilter{
			domain.CriteriaLarger: domain.NewLargerLayoutFilter(),
		},
	}
}

func (uc *ChintaiSoubaHikakuInteractor) Execute(
	input ChintaiSoubaHikakuInputData,
	presenter ChintaiSoubaHikakuPresenter,
	cityRepo domain.CityRepository,
	averageRentRepo domain.AverageRentRepository,
) error {
	// baseTokyoCityを取得
	baseCityCode, err := domain.ParseCityCode(input.CityCode)
	if err != nil {
		return presenter.PresentError(err)
	}
	baseTokyoCity, err := cityRepo.FindByCityCode(baseCityCode)
	if err != nil {
		return presenter.PresentError(err)
	}

	// baseAverageRentを取得
	layout, err := domain.NewLayout(input.Layout)
	if err != nil {
		return presenter.PresentError(err)
	}
	baseAverageRent, err := averageRentRepo.FindByCityIDAndLayout(baseTokyoCity.ID(), layout)
	if err != nil {
		return presenter.PresentError(err)
	}

	// saitamaCitiesを取得
	saitamaCities, err := cityRepo.FindByPrefectureCode(domain.SaitamaPrefectureCode)
	if err != nil {
		return presenter.PresentError(err)
	}
	saitamaCityIDs := domain.ExtractCityIDsFromCityMap(saitamaCities)

	// saitamaAverageRentsを取得
	min, err := baseAverageRent.Get10PercentLower()
	if err != nil {
		return presenter.PresentError(err)
	}
	max, err := baseAverageRent.Get10PercentUpper()
	if err != nil {
		return presenter.PresentError(err)
	}
	saitamaAverageRents, err := averageRentRepo.FindByAmountRangeAndCityIDs(min, max, saitamaCityIDs)
	if err != nil {
		return presenter.PresentError(err)
	}

	criteria, err := domain.NewChintaiSoubaHikakuCriteria(input.Criteria)
	if err != nil {
		return presenter.PresentError(err)
	}

	// filterを実行
	filteredAverageRents := uc.filters[criteria].Execute(
		criteria,
		baseAverageRent,
		saitamaAverageRents,
		saitamaCities,
	)

	// presenterに渡す
	outputData := ChintaiSoubaHikakuOutputData{}
	outputData.BaseTokyoCity.Name = baseTokyoCity.Name().String()
	outputData.BaseTokyoCity.Layout = baseAverageRent.Layout().String()
	outputData.BaseTokyoCity.AverageRent = baseAverageRent.Amount().String()
	outputData.BaseTokyoCity.CityCode = baseTokyoCity.Code().String()
	for _, r := range filteredAverageRents {
		outputData.SaitamaCities = append(outputData.SaitamaCities, struct {
			Name        string
			CityCode    string
			Layout      string
			AverageRent string
		}{
			Name:        saitamaCities[r.CityID()].Name().String(),
			CityCode:    saitamaCities[r.CityID()].Code().String(),
			Layout:      r.Layout().String(),
			AverageRent: r.Amount().String(),
		})
	}

	return presenter.Present(outputData)
}
