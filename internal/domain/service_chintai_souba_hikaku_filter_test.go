package domain_test

import (
	"testing"

	"github.com/iotassss/saitamarental/internal/domain"
	"github.com/iotassss/saitamarental/internal/unittest"
)

func TestLargerLayoutFilter_Execute(t *testing.T) {
	// 基準都市と平均家賃
	cityBase := unittest.Must(domain.NewCity(
		unittest.Must(domain.NewID(1)),
		unittest.Must(domain.NewCityName("千代田区")),
		unittest.Must(domain.ParseCityCode("13101")),
		unittest.Must(domain.NewCityOrder(1)),
	))
	baseAverageRent := unittest.Must(domain.NewAverageRent(
		unittest.Must(domain.NewID(10)),
		cityBase.ID(),
		unittest.Must(domain.NewAverageRentAmount("10.0")),
		unittest.Must(domain.NewLayout("ワンルーム")),
	))

	// 比較対象データ
	city2 := unittest.Must(domain.NewCity(
		unittest.Must(domain.NewID(2)),
		unittest.Must(domain.NewCityName("中央区")),
		unittest.Must(domain.ParseCityCode("13102")),
		unittest.Must(domain.NewCityOrder(2)),
	))
	sameAmountLargerRent := unittest.Must(domain.NewAverageRent(
		unittest.Must(domain.NewID(11)),
		city2.ID(),
		unittest.Must(domain.NewAverageRentAmount("10.0")),
		unittest.Must(domain.NewLayout("1K/1DK")), // 広い
	))

	city3 := unittest.Must(domain.NewCity(
		unittest.Must(domain.NewID(3)),
		unittest.Must(domain.NewCityName("港区")),
		unittest.Must(domain.ParseCityCode("13103")),
		unittest.Must(domain.NewCityOrder(3)),
	))
	differentAmountLargerRent := unittest.Must(domain.NewAverageRent(
		unittest.Must(domain.NewID(12)),
		city3.ID(),
		unittest.Must(domain.NewAverageRentAmount("9.8")),
		unittest.Must(domain.NewLayout("2LDK/3K")), // より広い
	))

	targetAverageRents := []*domain.AverageRent{sameAmountLargerRent, differentAmountLargerRent}
	targetCities := map[domain.ID]*domain.City{
		city2.ID(): city2,
		city3.ID(): city3,
	}

	filter := domain.NewLargerLayoutFilter()
	resultAverageRent := filter.Execute(domain.CriteriaLarger, baseAverageRent, targetAverageRents, targetCities)

	if len(resultAverageRent) != 2 {
		t.Fatalf("expected 2 resultAverageRents, got %d", len(resultAverageRent))
	}
	if resultAverageRent[0].ID() != sameAmountLargerRent.ID() {
		t.Errorf("first resultAverageRent should be sameAmountLargerRent")
	}
	if resultAverageRent[1].ID() != differentAmountLargerRent.ID() {
		t.Errorf("second resultAverageRent should be differentAmountLargerRent")
	}
}
