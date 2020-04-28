package models

//GetAllSnacks returns all the snacks
func GetAllSnacks() ([]GoodsBrief, error) {
	allSnacks, err := GetAllGoods("snacks")
	if err != nil {
		return nil, err
	}
	return allSnacks, nil
}

// GetSnacks return one type of fosnacksod query with id
func GetSnacks(id int) (GoodsDetail, error) {
	snacksDetail := GoodsDetail{}
	snacksDetail, err := GetGoods(id, "snacks")
	if err != nil {
		return snacksDetail, err
	}
	return snacksDetail, nil
}
