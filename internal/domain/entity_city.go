package domain

type City struct {
	id    ID
	name  CityName
	code  CityCode
	order CityOrder
}

func NewCity(id ID, name CityName, code CityCode, order CityOrder) (*City, error) {
	return &City{
		id:    id,
		name:  name,
		code:  code,
		order: order,
	}, nil
}

func (c *City) ID() ID           { return c.id }
func (c *City) Name() CityName   { return c.name }
func (c *City) Code() CityCode   { return c.code }
func (c *City) Order() CityOrder { return c.order }

// func (c *City) StationIDs() []ID { return append([]ID(nil), c.stationIDs...) } // スライスコピー
// func (c *City) RentList() []Rent { return append([]Rent(nil), c.rentList...) } // スライスコピー

func (c *City) SetID(id ID) error {
	if c.id != 0 {
		return ErrIDAlreadySet
	}
	c.id = id
	return nil
}

func (c *City) ComesBefore(city *City) bool {
	return c.order.ComesBefore(city.order)
}
