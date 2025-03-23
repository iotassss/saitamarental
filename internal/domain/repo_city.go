package domain

type CityRepository interface {
	FindById(id ID) (*City, error)
	FindMapByIds(ids []ID) (map[ID]*City, error)
	FindByCityCode(code CityCode) (*City, error)
	FindByPrefectureCode(code PrefectureCode) ([]*City, error)
	Save(city *City) (*City, error)
}
