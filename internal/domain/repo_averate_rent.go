package domain

type AverageRentRepository interface {
	FindById(id ID) (*AverageRent, error)
	FindByCityIDAndLayout(cityID ID, layout Layout) (*AverageRent, error)
	FindByAmountRangeAndCityIDs(min AverageRentAmount, max AverageRentAmount, cityIDs []ID) ([]*AverageRent, error)
	Save(averageRent *AverageRent) (*AverageRent, error)
}
