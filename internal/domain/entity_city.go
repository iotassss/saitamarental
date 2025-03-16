package domain

type City struct {
	id         ID
	name       CityName
	code       CityCode
	order      CityOrder
	stationIDs []ID
	rentList   []Rent
}

func NewCity(id ID, name CityName, code CityCode, order CityOrder, stationIDs []ID, rentList []Rent) (*City, error) {
	stationIDsCopy := append([]ID(nil), stationIDs...)
	rentListCopy := append([]Rent(nil), rentList...)

	return &City{
		id:         id,
		name:       name,
		code:       code,
		order:      order,
		stationIDs: stationIDsCopy,
		rentList:   rentListCopy,
	}, nil
}

// func (c *City) ID() ID           { return c.id }
// func (c *City) Name() CityName   { return c.name }
// func (c *City) Code() CityCode   { return c.code }
// func (c *City) Order() CityOrder { return c.order }
// func (c *City) StationIDs() []ID { return append([]ID(nil), c.stationIDs...) } // スライスコピー
// func (c *City) RentList() []Rent { return append([]Rent(nil), c.rentList...) } // スライスコピー
