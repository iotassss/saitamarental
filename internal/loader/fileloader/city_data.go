package fileloader

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/iotassss/saitamarental/internal/domain"
	"gopkg.in/yaml.v3"
)

type CityFileLoader struct {
	orderListFileName string
	cityDataFileName  string
}

func NewCityFileLoader(
	orderListFileName string,
	cityDataFileName string,
) *CityFileLoader {
	return &CityFileLoader{
		orderListFileName: orderListFileName,
		cityDataFileName:  cityDataFileName,
	}
}

// YAML構造
type rawCity struct {
	Name         string             `yaml:"name"`
	CityCode     string             `yaml:"city_code"`
	AverageRents map[string]float64 `yaml:"average_rents"`
}

func (l *CityFileLoader) Load() ([]*domain.City, error) {
	// order list 読み込み
	orderMap, err := loadCityOrder(l.orderListFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to load %s: %w", l.orderListFileName, err)
	}

	// Cityデータ 読み込み
	citiesFile, err := os.ReadFile(l.cityDataFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", l.cityDataFileName, err)
	}

	var rawCities []rawCity
	if err := yaml.Unmarshal(citiesFile, &rawCities); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	var cities []*domain.City
	for _, rc := range rawCities {
		code, err := domain.ParseCityCode(rc.CityCode)
		if err != nil {
			return nil, fmt.Errorf("invalid city code '%s': %w", rc.CityCode, err)
		}
		name, err := domain.NewCityName(rc.Name)
		if err != nil {
			return nil, fmt.Errorf("invalid city name '%s': %w", rc.Name, err)
		}

		order, ok := orderMap[rc.Name]
		if !ok {
			return nil, fmt.Errorf("city name '%s' not found in city_order_list.txt", rc.Name)
		}
		orderVal, err := domain.NewCityOrder(order)
		if err != nil {
			return nil, fmt.Errorf("invalid city order '%d': %w", order, err)
		}

		city, err := domain.NewCity(0, name, code, orderVal)
		if err != nil {
			return nil, fmt.Errorf("failed to create city: %w", err)
		}

		cities = append(cities, city)
	}

	return cities, nil
}

func loadCityOrder(path string) (map[string]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	order := make(map[string]int)
	scanner := bufio.NewScanner(f)
	idx := 1
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		_, err := domain.NewCityName(line)
		if err != nil {
			return nil, fmt.Errorf("invalid city name '%s': %w", line, err)
		}
		if line != "" {
			order[line] = idx
			idx++
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return order, nil
}
