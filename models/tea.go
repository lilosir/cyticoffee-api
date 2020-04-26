package models

//GetAllTea returns all the coffee
func GetAllTea() ([]GoodsBrief, error) {
	allTea, err := GetAllGoods("tea")
	if err != nil {
		return nil, err
	}
	return allTea, nil
}

// GetTea return one type of tea query with id
func GetTea(id int) (GoodsDetail, error) {
	teaDetail := GoodsDetail{}
	teaDetail, err := GetGoods(id, "tea")
	if err != nil {
		return teaDetail, err
	}
	return teaDetail, nil
}
