package domain

import (
	"errors"
	"fmt"
	"strings"
)

var cities = []string{
	// Saitama
	"上尾市", "朝霞市", "伊奈町", "入間市", "小鹿野町", "小川町", "桶川市", "越生町",
	"春日部市", "加須市", "神川町", "上里町", "川口市", "川越市", "川島町", "北本市", "行田市",
	"久喜市", "熊谷市", "鴻巣市", "越谷市",
	"さいたま市", "さいたま市西区", "さいたま市北区", "さいたま市大宮区", "さいたま市見沼区",
	"さいたま市中央区", "さいたま市桜区", "さいたま市浦和区", "さいたま市南区",
	"さいたま市緑区", "さいたま市岩槻区",
	"坂戸市", "幸手市", "狭山市", "志木市", "白岡市", "杉戸町", "草加市",
	"秩父市", "鶴ヶ島市", "ときがわ町", "所沢市", "戸田市",
	"長瀞町", "滑川町", "新座市",
	"蓮田市", "鳩山町", "羽生市", "飯能市", "東秩父村", "東松山市", "日高市", "深谷市",
	"富士見市", "ふじみ野市", "本庄市",
	"松伏町", "三郷市", "美里町", "皆野町", "宮代町", "三芳町", "毛呂山町",
	"八潮市", "横瀬町", "吉川市", "吉見町", "寄居町",
	"嵐山町",
	"和光市", "蕨市",
	// Tokyo
	"千代田区", "中央区", "港区", "新宿区", "文京区", "台東区", "墨田区", "江東区", "品川区", "目黒区",
	"大田区", "世田谷区", "渋谷区", "中野区", "杉並区", "豊島区", "北区", "荒川区", "板橋区", "練馬区",
	"足立区", "葛飾区", "江戸川区",
	"八王子市", "立川市", "武蔵野市", "三鷹市", "青梅市", "府中市", "昭島市", "調布市", "町田市", "小金井市",
	"小平市", "日野市", "東村山市", "国分寺市", "国立市", "福生市", "狛江市", "東大和市", "清瀬市", "東久留米市",
	"武蔵村山市", "多摩市", "稲城市", "羽村市", "あきる野市", "西東京市",
	"瑞穂町", "日の出町", "檜原村", "奥多摩町",
	"大島町", "利島村", "新島村", "神津島村",
	"三宅村", "御蔵島村",
	"八丈町", "青ヶ島村",
	"小笠原村",
}

type CityName struct {
	properNounName string
	class          string
}

func NewCityName(name string) (CityName, error) {
	if name == "" {
		return CityName{}, fmt.Errorf("city name cannot be empty")
	}

	found := false
	for _, city := range cities {
		if city == name {
			found = true
			break
		}
	}
	if !found {
		return CityName{}, fmt.Errorf("city name '%s' is not recognized", name)
	}

	properNounName, class, err := parseCityName(name)
	if err != nil {
		return CityName{}, err
	}

	return CityName{
		properNounName: properNounName,
		class:          class,
	}, nil
}

func parseCityName(name string) (string, string, error) {
	suffixes := []string{"市", "区", "町", "村"}

	for _, suffix := range suffixes {
		if strings.HasSuffix(name, suffix) {
			properNounName := strings.TrimSuffix(name, suffix)
			return properNounName, suffix, nil
		}
	}

	return "", "", errors.New("unknown city name format")
}

func (c CityName) ProperNounName() string {
	return c.properNounName
}

func (c CityName) Class() string {
	return c.class
}

func (c CityName) String() string {
	return c.properNounName + c.class
}
