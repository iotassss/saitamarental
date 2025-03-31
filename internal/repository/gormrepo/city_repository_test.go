package gormrepo_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
	"github.com/iotassss/saitamarental/internal/repository/gormrepo"
	"github.com/iotassss/saitamarental/internal/unittest"
	"github.com/stretchr/testify/require"
)

func createTestCity() *domain.City {
	name := unittest.Must(domain.NewCityName("さいたま市桜区"))
	code := unittest.Must(domain.ParseCityCode("13001"))
	order := unittest.Must(domain.NewCityOrder(1))
	city := unittest.Must(domain.NewCity(0, name, code, order))
	return city
}

func TestCityRepo_SaveAndFind(t *testing.T) {
	db := setupTestDB(t)
	repo := gormrepo.NewGormCityRepo(db)
	city := createTestCity()

	saved, err := repo.Save(city)
	require.NoError(t, err)
	require.Equal(t, city.Name().String(), saved.Name().String())

	fetched, err := repo.FindByCityCode(city.Code())
	require.NoError(t, err)
	require.Equal(t, city.Name().String(), fetched.Name().String())
}

func TestCityRepo_FindById(t *testing.T) {
	db := setupTestDB(t)
	repo := gormrepo.NewGormCityRepo(db)
	city := createTestCity()
	saved, err := repo.Save(city)
	require.NoError(t, err)

	found, err := repo.FindById(saved.ID())
	require.NoError(t, err)
	require.Equal(t, saved.Name().String(), found.Name().String())
}

func TestCityRepo_FindByPrefectureCode(t *testing.T) {
	db := setupTestDB(t)
	repo := gormrepo.NewGormCityRepo(db)
	city := createTestCity()
	saved, err := repo.Save(city)
	require.NoError(t, err)

	foundMap, err := repo.FindByPrefectureCode(city.Code().PrefectureCode())
	require.NoError(t, err)
	found, ok := foundMap[saved.ID()]
	require.True(t, ok)
	require.Equal(t, saved.Name().String(), found.Name().String())
}

func TestCityRepo_FindByIds(t *testing.T) {
	db := setupTestDB(t)
	repo := gormrepo.NewGormCityRepo(db)
	city := createTestCity()
	saved, err := repo.Save(city)
	require.NoError(t, err)

	foundMap, err := repo.FindByIds([]domain.ID{saved.ID()})
	require.NoError(t, err)
	found, ok := foundMap[saved.ID()]
	require.True(t, ok)
	require.Equal(t, saved.Name().String(), found.Name().String())
}
