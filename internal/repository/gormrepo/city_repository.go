package gormrepo

import (
	"errors"

	"github.com/iotassss/saitamarental/internal/domain"
	"gorm.io/gorm"
)

type CityModel struct {
	gorm.Model
	Name           string `gorm:"column:name;not null;size:100"`
	PrefectureCode string `gorm:"column:prefecture_code;not null;size:2"`
	CityCode       string `gorm:"column:city_code;unique;not null;size:5"`
	Order          int    `gorm:"column:order;not null"`
}

func (CityModel) TableName() string {
	return "cities"
}

func toCityModel(city *domain.City) *CityModel {
	return &CityModel{
		Model:          gorm.Model{ID: uint(city.ID())},
		Name:           city.Name().String(),
		PrefectureCode: city.Code().PrefectureCode().String(),
		CityCode:       city.Code().String(),
		Order:          city.Order().Value(),
	}
}

func toCityDomain(model CityModel) (*domain.City, error) {
	name, err := domain.NewCityName(model.Name)
	if err != nil {
		return nil, err
	}
	code, err := domain.ParseCityCode(model.CityCode)
	if err != nil {
		return nil, err
	}
	order, err := domain.NewCityOrder(model.Order)
	if err != nil {
		return nil, err
	}
	return domain.NewCity(domain.ID(model.ID), name, code, order)
}

type CityRepo struct {
	db *gorm.DB
}

func NewGormCityRepo(db *gorm.DB) *CityRepo {
	return &CityRepo{db: db}
}

func (r *CityRepo) FindById(id domain.ID) (*domain.City, error) {
	var model CityModel
	if err := r.db.First(&model, "id = ?", uint(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrEntityNotFound
		}
		return nil, err
	}
	return toCityDomain(model)
}

func (r *CityRepo) FindByIds(ids []domain.ID) (map[domain.ID]*domain.City, error) {
	var models []CityModel
	uids := make([]uint, len(ids))
	for i, id := range ids {
		uids[i] = uint(id)
	}
	if err := r.db.Where("id IN ?", uids).Find(&models).Error; err != nil {
		return nil, err
	}

	result := make(map[domain.ID]*domain.City)
	for _, model := range models {
		city, err := toCityDomain(model)
		if err != nil {
			return nil, err
		}
		result[city.ID()] = city
	}
	return result, nil
}

func (r *CityRepo) FindByCityCode(code domain.CityCode) (*domain.City, error) {
	var model CityModel
	if err := r.db.First(&model, "city_code = ?", code.String()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrEntityNotFound
		}
		return nil, err
	}
	return toCityDomain(model)
}

func (r *CityRepo) FindByPrefectureCode(code domain.PrefectureCode) (map[domain.ID]*domain.City, error) {
	var models []CityModel
	if err := r.db.Where("prefecture_code = ?", code.String()).Find(&models).Error; err != nil {
		return nil, err
	}

	result := make(map[domain.ID]*domain.City)
	for _, model := range models {
		city, err := toCityDomain(model)
		if err != nil {
			return nil, err
		}
		result[city.ID()] = city
	}
	return result, nil
}

func (r *CityRepo) Save(city *domain.City) (*domain.City, error) {
	model := toCityModel(city)
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	id, err := domain.NewID(int(model.ID))
	if err != nil {
		return nil, err
	}
	if err = city.SetID(id); err != nil {
		return nil, err
	}
	return city, nil
}
