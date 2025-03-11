package domain

type Rent struct {
	id     ID
	amount RentAmount
	layout Layout
}

func NewRent(id ID, amount RentAmount, layout Layout) *Rent {
	return &Rent{
		id:     id,
		amount: amount,
		layout: layout,
	}
}
