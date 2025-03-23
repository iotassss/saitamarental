package domain

type AverageRent struct {
	id     ID
	cityID ID
	amount AverageRentAmount
	layout Layout
}

func NewAverageRent(id ID, cityID ID, amount AverageRentAmount, layout Layout) (*AverageRent, error) {
	return &AverageRent{
		id:     id,
		cityID: cityID,
		amount: amount,
		layout: layout,
	}, nil
}

func (r *AverageRent) ID() ID                    { return r.id }
func (r *AverageRent) CityID() ID                { return r.cityID }
func (r *AverageRent) Amount() AverageRentAmount { return r.amount }
func (r *AverageRent) Layout() Layout            { return r.layout }

func (r *AverageRent) Get10PercentUpper() (AverageRentAmount, error) {
	return r.amount.Get10PercentUpper()
}
func (r *AverageRent) Get10PercentLower() (AverageRentAmount, error) {
	return r.amount.Get10PercentLower()
}
func (r *AverageRent) IsLargerLayoutThan(targetAverageRent *AverageRent) bool {
	return r.layout.IsLargerThan(targetAverageRent.Layout())
}
