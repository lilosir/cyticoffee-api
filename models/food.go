package models

//GetAllFood returns all the food
func GetAllFood() ([]GoodsBrief, error) {
	allFood, err := GetAllGoods("food")
	if err != nil {
		return nil, err
	}
	return allFood, nil
}

// GetFood return one type of food query with id
func GetFood(id int) (GoodsDetail, error) {
	foodDetail := GoodsDetail{}
	foodDetail, err := GetGoods(id, "food")
	if err != nil {
		return foodDetail, err
	}
	return foodDetail, nil
}
