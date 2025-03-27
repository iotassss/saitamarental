package domain

func ExtractCityIDsFromCityMap(saitamaCities map[ID]*City) []ID {
	var saitamaCityIDs []ID
	for id := range saitamaCities {
		saitamaCityIDs = append(saitamaCityIDs, id)
	}
	return saitamaCityIDs
}
