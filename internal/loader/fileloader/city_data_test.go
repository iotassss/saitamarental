package fileloader_test

import (
	"path/filepath"
	"testing"

	"github.com/iotassss/saitamarental/internal/loader/fileloader"
	"github.com/stretchr/testify/assert"
)

func TestCityFileLoader_Load_Valid(t *testing.T) {
	root, err := filepath.Abs("../../../") // <- ルートまで戻る
	if err != nil {
		t.Fatal(err)
	}

	orderPath := filepath.Join(root, "data_sample/city_order_list.txt")
	cityPath := filepath.Join(root, "data_sample/cities.yaml")

	loader := fileloader.NewCityFileLoader(orderPath, cityPath)

	cities, err := loader.Load()
	assert.NoError(t, err)
	assert.NotEmpty(t, cities)

	firstCity := cities[0]
	assert.Equal(t, "江東区", firstCity.Name().String())
	assert.Equal(t, "13113", firstCity.Code().String())
	assert.Equal(t, 2, firstCity.Order().Value())
	secondCity := cities[1]
	assert.Equal(t, 3, secondCity.Order().Value())
	thirdCity := cities[2]
	assert.Equal(t, 1, thirdCity.Order().Value())
}
