package domain

func ExtractCityIDs(cities []*City) []ID {
	cityIDmap := map[ID]bool{}
	for _, city := range cities {
		cityIDmap[city.ID()] = true
	}
	cityIDs := []ID{}
	for cityID := range cityIDmap {
		cityIDs = append(cityIDs, cityID)
	}

	return cityIDs
}
