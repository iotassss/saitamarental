package gormrepo

import (
	"errors"

	"github.com/iotassss/saitamarental/internal/domain"
	"gorm.io/gorm"
)

type AverageRentModel struct {
	ID          uint    `gorm:"primaryKey;column:id"`
	CityID      uint    `gorm:"column:city_id;not null"`
	Layout      string  `gorm:"column:layout;not null;size:10"`
	Amount      string  `gorm:"column:amount;not null;size:16"`
	FloatAmount float64 `gorm:"column:float_amount;not null"`
}

func (AverageRentModel) TableName() string {
	return "average_rents"
}

func toAverageRentModel(ar *domain.AverageRent) *AverageRentModel {
	return &AverageRentModel{
		ID:          uint(ar.ID()),
		CityID:      uint(ar.CityID()),
		Layout:      string(ar.Layout()),
		Amount:      ar.Amount().String(),
		FloatAmount: ar.Amount().Number(),
	}
}

func toAverageRentDomain(m *AverageRentModel) (*domain.AverageRent, error) {
	amount, err := domain.NewAverageRentAmount(m.Amount)
	if err != nil {
		return nil, err
	}
	return domain.NewAverageRent(
		domain.ID(m.ID),
		domain.ID(m.CityID),
		amount,
		domain.Layout(m.Layout),
	)
}

type AverageRentRepo struct {
	db *gorm.DB
}

func NewAverageRentRepo(db *gorm.DB) *AverageRentRepo {
	return &AverageRentRepo{db: db}
}

func (r *AverageRentRepo) FindById(id domain.ID) (*domain.AverageRent, error) {
	var model AverageRentModel
	if err := r.db.First(&model, "id = ?", uint(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrEntityNotFound
		}
		return nil, err
	}
	return toAverageRentDomain(&model)
}

func (r *AverageRentRepo) FindByCityIDAndLayout(cityID domain.ID, layout domain.Layout) (*domain.AverageRent, error) {
	var model AverageRentModel
	if err := r.db.
		Where("city_id = ? AND layout = ?", uint(cityID), string(layout)).
		First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrEntityNotFound
		}
		return nil, err
	}
	return toAverageRentDomain(&model)
}

func (r *AverageRentRepo) FindByAmountRangeAndCityIDs(min domain.AverageRentAmount, max domain.AverageRentAmount, cityIDs []domain.ID) ([]*domain.AverageRent, error) {
	var models []AverageRentModel
	idUints := make([]uint, len(cityIDs))
	for i, id := range cityIDs {
		idUints[i] = uint(id)
	}
	if err := r.db.
		Where("city_id IN ? AND float_amount > ? AND float_amount < ?", idUints, min.Number(), max.Number()).
		Find(&models).Error; err != nil {
		return nil, err
	}

	result := make([]*domain.AverageRent, 0, len(models))
	for _, model := range models {
		ar, err := toAverageRentDomain(&model)
		if err != nil {
			return nil, err
		}
		result = append(result, ar)
	}
	return result, nil
}

func (r *AverageRentRepo) Save(ar *domain.AverageRent) (*domain.AverageRent, error) {
	model := toAverageRentModel(ar)
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	id, err := domain.NewID(int(model.ID))
	if err != nil {
		return nil, err
	}
	if err = ar.SetID(id); err != nil {
		return nil, err
	}
	return ar, nil
}
