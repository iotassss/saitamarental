package domain

type Station struct {
	id         ID
	toTokyo    TravelTime
	toShinjuku TravelTime
}

func NewStation(id ID, toTokyo TravelTime, toShinjuku TravelTime) (*Station, error) {
	return &Station{
		id:         id,
		toTokyo:    toTokyo,
		toShinjuku: toShinjuku,
	}, nil
}
