package memory

import (
	"errors"
	"sync"

	"github.com/iotassss/saitamarental/internal/domain"
)

type InMemoryCityRepo struct {
	data map[domain.ID]*domain.City
	mu   sync.RWMutex
}

func NewInMemoryCityRepo() *InMemoryCityRepo {
	return &InMemoryCityRepo{
		data: make(map[domain.ID]*domain.City),
	}
}

func (r *InMemoryCityRepo) FindById(id domain.ID) (*domain.City, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	city, ok := r.data[id]
	if !ok {
		return nil, errors.New("city not found")
	}
	return city, nil
}

func (r *InMemoryCityRepo) FindByIds(ids []domain.ID) (map[domain.ID]*domain.City, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make(map[domain.ID]*domain.City)
	for _, id := range ids {
		if city, ok := r.data[id]; ok {
			result[id] = city
		}
	}
	return result, nil
}

func (r *InMemoryCityRepo) FindByCityCode(code domain.CityCode) (*domain.City, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, city := range r.data {
		if city.Code().Equals(code) {
			return city, nil
		}
	}
	return nil, errors.New("city not found")
}

func (r *InMemoryCityRepo) FindByPrefectureCode(code domain.PrefectureCode) (map[domain.ID]*domain.City, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make(map[domain.ID]*domain.City)
	for id, city := range r.data {
		if city.Code().PrefectureCode().Equals(code) {
			result[id] = city
		}
	}
	return result, nil
}

func (r *InMemoryCityRepo) Save(city *domain.City) (*domain.City, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[city.ID()] = city
	return city, nil
}
