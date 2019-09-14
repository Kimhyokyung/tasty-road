package models

import "errors"

type lunch struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Location string `json:"location"`
	Contact  string `json:"contact"`
}

var lunchList = []lunch{
	{ID: 0, Name: "마음은콩밭에", Category: "한식"},
	{ID: 1, Name: "서초지하식당", Category: "Undefined"},
	{ID: 2, Name: "맘스터치", Category: "패스트푸드"},
}

func GetLunchList() []lunch {
	return lunchList
}

func GetLunchById(lunchId int) (*lunch, error) {
	for _, lunch := range lunchList {
		if lunch.ID == lunchId {
			return &lunch, nil
		}
	}
	return nil, errors.New("Lunch Not Found")
}
