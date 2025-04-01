package gormrepo_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
	"github.com/iotassss/saitamarental/internal/repository/gormrepo"
	"github.com/iotassss/saitamarental/internal/unittest"
	"github.com/stretchr/testify/require"
)

func createTestAverageRent(cityID domain.ID) *domain.AverageRent {
	id := unittest.Must(domain.NewID(0))
	amount := unittest.Must(domain.NewAverageRentAmount("12.3"))
	layout := unittest.Must(domain.NewLayout("1LDK/2K"))
	return unittest.Must(domain.NewAverageRent(id, cityID, amount, layout))
}

func TestAverageRentRepo_SaveAndFindById(t *testing.T) {
	db := setupTestDB(t)
	cityRepo := gormrepo.NewGormCityRepo(db)
	arRepo := gormrepo.NewAverageRentRepo(db)

	city := createTestCity()
	savedCity, err := cityRepo.Save(city)
	require.NoError(t, err)

	ar := createTestAverageRent(savedCity.ID())
	savedAverageRent, err := arRepo.Save(ar)
	require.NoError(t, err)

	found, err := arRepo.FindById(savedAverageRent.ID())
	require.NoError(t, err)
	require.Equal(t, savedAverageRent.Amount().String(), found.Amount().String())
}

func TestAverageRentRepo_FindByCityIDAndLayout(t *testing.T) {
	db := setupTestDB(t)
	cityRepo := gormrepo.NewGormCityRepo(db)
	arRepo := gormrepo.NewAverageRentRepo(db)

	city := createTestCity()
	savedCity, err := cityRepo.Save(city)
	require.NoError(t, err)

	ar := createTestAverageRent(savedCity.ID())
	_, err = arRepo.Save(ar)
	require.NoError(t, err)

	found, err := arRepo.FindByCityIDAndLayout(savedCity.ID(), domain.Layout("1LDK/2K"))
	require.NoError(t, err)
	require.Equal(t, ar.Amount().String(), found.Amount().String())
}

func TestAverageRentRepo_FindByAmountRangeAndCityIDs(t *testing.T) {
	db := setupTestDB(t)
	cityRepo := gormrepo.NewGormCityRepo(db)
	arRepo := gormrepo.NewAverageRentRepo(db)

	city := createTestCity()
	savedCity, err := cityRepo.Save(city)
	require.NoError(t, err)

	ar := createTestAverageRent(savedCity.ID())
	_, err = arRepo.Save(ar)
	require.NoError(t, err)

	min := unittest.Must(domain.NewAverageRentAmount("10.0"))
	max := unittest.Must(domain.NewAverageRentAmount("15.5"))

	result, err := arRepo.FindByAmountRangeAndCityIDs(min, max, []domain.ID{savedCity.ID()})
	require.NoError(t, err)
	require.Len(t, result, 1)
	require.Equal(t, ar.Amount().String(), result[0].Amount().String())
}
