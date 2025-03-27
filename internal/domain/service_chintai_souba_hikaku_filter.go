package domain

import (
	"sort"
)

type ChintaiSoubaHikakuFilter interface {
	Execute(
		criteria ChintaiSoubaHikakuCriteria,
		baseAverageRent *AverageRent,
		targetAverageRents []*AverageRent,
		targetCities map[ID]*City,
	) []*AverageRent
}

type LargerLayoutFilter struct{}

func NewLargerLayoutFilter() *LargerLayoutFilter {
	return &LargerLayoutFilter{}
}

func (f *LargerLayoutFilter) Execute(
	criteria ChintaiSoubaHikakuCriteria,
	baseAverageRent *AverageRent,
	targetAverageRents []*AverageRent,
	targetCities map[ID]*City,
) []*AverageRent {
	if criteria != "larger" {
		return nil
	}

	var result []*AverageRent

	// 同じ家賃でより広い
	var sameAmountLarger []*AverageRent
	for _, t := range targetAverageRents {
		if t.Amount() == baseAverageRent.Amount() &&
			t.IsLargerLayoutThan(baseAverageRent) {
			sameAmountLarger = append(sameAmountLarger, t)
		}
	}

	var bestSameAmountLarger *AverageRent
	if len(sameAmountLarger) > 0 {
		bestSameAmountLarger = sameAmountLarger[0]
		for _, s := range sameAmountLarger {
			sameAmountLargerCity := targetCities[s.CityID()]
			bestSameAmountLargerCity := targetCities[bestSameAmountLarger.CityID()]
			if sameAmountLargerCity.ComesBefore(bestSameAmountLargerCity) {
				bestSameAmountLarger = s
			}
		}
		result = append(result, bestSameAmountLarger)
	}

	// より広くて order が小さい上位5件
	var larger []*AverageRent
	for _, t := range targetAverageRents {
		if t.IsLargerLayoutThan(baseAverageRent) &&
			t.Amount() != baseAverageRent.Amount() {
			larger = append(larger, t)
		}
	}
	sort.Slice(larger, func(i, j int) bool {
		return targetCities[larger[i].CityID()].ComesBefore(targetCities[larger[j].CityID()])
	})

	for i := 0; i < 5 && i < len(larger); i++ {
		result = append(result, larger[i])
	}

	return result
}

// closer_to_shinjukuの場合（phase2で実装）
// ...
