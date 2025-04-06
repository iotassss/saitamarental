package domain

type CityDataLoader interface {
	Load() ([]*City, error)
}

// ロードするファイルが膨大になってきたらCityDataStreamLoader interfaceを実装して
// ストリームで読み込むようにする

// 参考
// type CityDataStreamLoader interface {
// 	LoadStream() (<-chan *City, error)
// 	Close() error
// }
