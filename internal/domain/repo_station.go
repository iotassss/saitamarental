package domain

type StationRepository interface {
	FindById(id ID) (*Station, error)
	Save(station *Station) (*Station, error)
}
