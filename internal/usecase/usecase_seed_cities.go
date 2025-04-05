package usecase

import (
	"github.com/iotassss/saitamarental/internal/domain"
)

// Seedするファイルが増えてきたらChunkSizeを指定できるようにする
// type SeedCitiesInputData struct {
// 	ChunkSize int
// }

// 現在は全件メモリに展開してSeedする想定

type SeedCitiesOutputData struct {
	SeededCount int
}

type SeedCitiesPresenter interface {
	Present(output SeedCitiesOutputData) error
	PresentError(err error) error
}

type SeedCitiesUsecase interface {
	Execute(
		loader domain.CityDataLoader,
		repo domain.CityRepository,
		presenter SeedCitiesPresenter,
	) error
}
