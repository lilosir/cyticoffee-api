package models

//GetAllCoffee returns all the coffee
func GetAllCoffee() ([]GoodsBrief, error) {
	allCoffee, err := GetAllGoods("coffee")
	if err != nil {
		return nil, err
	}
	return allCoffee, nil
}

// GetCoffee return one type of coffee query with id
func GetCoffee(id int) (GoodsDetail, error) {
	coffeeDetail := GoodsDetail{}
	coffeeDetail, err := GetGoods(id, "coffee")
	if err != nil {
		return coffeeDetail, err
	}
	return coffeeDetail, nil
}
