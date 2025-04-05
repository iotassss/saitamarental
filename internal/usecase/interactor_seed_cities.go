package usecase

import "github.com/iotassss/saitamarental/internal/domain"

type seedCitiesInteractor struct{}

func NewSeedCitiesInteractor() *seedCitiesInteractor {
	return &seedCitiesInteractor{}
}

func (uc *seedCitiesInteractor) Execute(
	loader domain.CityDataLoader,
	repo domain.CityRepository,
	presenter SeedCitiesPresenter,
) error {
	cities, err := loader.Load()
	if err != nil {
		return presenter.PresentError(err)
	}

	// TODO: SaveAllを定義してChunkingを検討
	// 現在はデータ数が少ないので、1件ずつ保存
	count := 0
	for _, city := range cities {
		_, err := repo.Save(city)
		if err != nil {
			return presenter.PresentError(err)
		}
		count++
	}

	return presenter.Present(SeedCitiesOutputData{SeededCount: count})
}
