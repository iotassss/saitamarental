package memory

import (
	"errors"
	"sync"

	"github.com/iotassss/saitamarental/internal/domain"
)

type InMemoryAverageRentRepo struct {
	data map[domain.ID]*domain.AverageRent
	mu   sync.RWMutex
}

func NewInMemoryAverageRentRepo() *InMemoryAverageRentRepo {
	return &InMemoryAverageRentRepo{
		data: make(map[domain.ID]*domain.AverageRent),
	}
}

func (r *InMemoryAverageRentRepo) FindById(id domain.ID) (*domain.AverageRent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	rent, ok := r.data[id]
	if !ok {
		return nil, errors.New("average rent not found")
	}
	return rent, nil
}

func (r *InMemoryAverageRentRepo) FindByCityIDAndLayout(cityID domain.ID, layout domain.Layout) (*domain.AverageRent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, rent := range r.data {
		if rent.CityID() == cityID && rent.Layout() == layout {
			return rent, nil
		}
	}
	return nil, errors.New("average rent not found")
}

func (r *InMemoryAverageRentRepo) FindByAmountRangeAndCityIDs(min domain.AverageRentAmount, max domain.AverageRentAmount, cityIDs []domain.ID) ([]*domain.AverageRent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	idSet := make(map[domain.ID]struct{}, len(cityIDs))
	for _, id := range cityIDs {
		idSet[id] = struct{}{}
	}

	var result []*domain.AverageRent
	for _, rent := range r.data {
		if _, ok := idSet[rent.CityID()]; ok {
			if rent.Amount().IsLargerThan(min) && rent.Amount().IsSmallerThan(max) {
				result = append(result, rent)
			}
		}
	}
	return result, nil
}

func (r *InMemoryAverageRentRepo) Save(averageRent *domain.AverageRent) (*domain.AverageRent, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[averageRent.ID()] = averageRent
	return averageRent, nil
}
