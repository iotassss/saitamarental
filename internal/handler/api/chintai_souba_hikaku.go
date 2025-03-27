package handler

// /chintai/souba/hikaku/tokyo-vs-saitama

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/saitamarental/internal/domain"
	"github.com/iotassss/saitamarental/internal/repository/memory"
	"github.com/iotassss/saitamarental/internal/unittest"
	"github.com/iotassss/saitamarental/internal/usecase"
	"gorm.io/gorm"
)

type ChintaiSoubaHikakuTokyoVsSaitamaHandler struct {
	db      *gorm.DB
	usecase usecase.ChintaiSoubaHikakuUsecase
}

func NewChintaiSoubaHikakuTokyoVsSaitamaHandler(
	db *gorm.DB,
	usecase usecase.ChintaiSoubaHikakuUsecase,
) *ChintaiSoubaHikakuTokyoVsSaitamaHandler {
	return &ChintaiSoubaHikakuTokyoVsSaitamaHandler{
		db:      db,
		usecase: usecase,
	}
}

func (h *ChintaiSoubaHikakuTokyoVsSaitamaHandler) Execute(c *gin.Context) {
	cityCode := c.Query("city_code")
	if cityCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "city_code is required"})
		return
	}
	layout := c.Query("layout")
	if layout == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "layout is required"})
		return
	}
	criteria := c.Query("criteria")
	if criteria == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "criteria is required"})
		return
	}

	input := usecase.ChintaiSoubaHikakuInputData{
		CityCode: cityCode,
		Layout:   layout,
		Criteria: criteria,
	}

	cityRepo := memory.NewInMemoryCityRepo()
	averageRentRepo := memory.NewInMemoryAverageRentRepo()
	presenter := NewPresenter(c)

	// test data
	city1 := unittest.Must(domain.NewCity(
		unittest.Must(domain.NewID(1)),
		unittest.Must(domain.NewCityName("さいたま市大宮区")),
		unittest.Must(domain.ParseCityCode("11201")),
		unittest.Must(domain.NewCityOrder(1)),
	))
	averageRent1 := unittest.Must(domain.NewAverageRent(
		unittest.Must(domain.NewID(1)),
		city1.ID(),
		unittest.Must(domain.NewAverageRentAmount("12.3")),
		unittest.Must(domain.NewLayout("ワンルーム")),
	))
	averageRent2 := unittest.Must(domain.NewAverageRent(
		unittest.Must(domain.NewID(2)),
		city1.ID(),
		unittest.Must(domain.NewAverageRentAmount("45.6")),
		unittest.Must(domain.NewLayout("1K/1DK")),
	))
	city2 := unittest.Must(domain.NewCity(
		unittest.Must(domain.NewID(2)),
		unittest.Must(domain.NewCityName("川口市")),
		unittest.Must(domain.ParseCityCode("11202")),
		unittest.Must(domain.NewCityOrder(2)),
	))
	averageRent3 := unittest.Must(domain.NewAverageRent(
		unittest.Must(domain.NewID(3)),
		city2.ID(),
		unittest.Must(domain.NewAverageRentAmount("7.8")),
		unittest.Must(domain.NewLayout("2LDK/3K")),
	))
	averageRent4 := unittest.Must(domain.NewAverageRent(
		unittest.Must(domain.NewID(4)),
		city2.ID(),
		unittest.Must(domain.NewAverageRentAmount("9.0")),
		unittest.Must(domain.NewLayout("3LDK/4K~")),
	))

	// save test data
	cityRepo.Save(city1)
	cityRepo.Save(city2)
	averageRentRepo.Save(averageRent1)
	averageRentRepo.Save(averageRent2)
	averageRentRepo.Save(averageRent3)
	averageRentRepo.Save(averageRent4)

	err := h.usecase.Execute(input, presenter, cityRepo, averageRentRepo)
	if err != nil {
		slog.Error("usecase.Execute failed", slog.Any("error", err))
		return
	}
}

type Presenter struct {
	ctx *gin.Context
}

func NewPresenter(ctx *gin.Context) *Presenter {
	return &Presenter{
		ctx: ctx,
	}
}

func (p *Presenter) Present(outputData usecase.ChintaiSoubaHikakuOutputData) error {
	p.ctx.JSON(http.StatusOK, gin.H{
		"base_tokyo_city": outputData.BaseTokyoCity,
		"saitama_cities":  outputData.SaitamaCities,
	})
	return nil
}

func (p *Presenter) PresentError(err error) error {
	p.ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return nil
}
