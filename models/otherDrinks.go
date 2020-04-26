package models

//GetAllOtherDrinks returns all the food
func GetAllOtherDrinks() ([]GoodsBrief, error) {
	allOtherDrinks, err := GetAllGoods("other_drinks")
	if err != nil {
		return nil, err
	}
	return allOtherDrinks, nil
}

// GetOtherDrinks return one type of food query with id
func GetOtherDrinks(id int) (GoodsDetail, error) {
	otherDrinksDetail := GoodsDetail{}
	otherDrinksDetail, err := GetGoods(id, "other_drinks")
	if err != nil {
		return otherDrinksDetail, err
	}
	return otherDrinksDetail, nil
}
