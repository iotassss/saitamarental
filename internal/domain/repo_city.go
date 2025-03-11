package domain

type CityRepository interface {
	FindById(id ID) (*City, error)
	Save(city *City) (*City, error)
}
