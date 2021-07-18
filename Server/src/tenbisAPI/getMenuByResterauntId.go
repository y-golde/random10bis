package tenbisAPI

import (
	"random10bis/src/util"
)

type MenuAPIResponse struct {
	CategoriesList []category `json:"categoriesList"`
}

type category struct {
	Name string `json:"categoryName"`
	Desc string `json:"categoryDesc"`
	DishList     []dish `json:"dishList"` 
}

type dish struct {
	Name                string   `json:"dishName"`
	IsPopularDish       bool     `json:"isPopularDish"`
	PopularityScore     int      `json:"dishPopularityScore"`
	Price               float32  `json:"dishPrice"`
	Desctiption         string   `json:"dishDescription"`
}

const tenbisEndpoint = "https://tenbis-static.azureedge.net/restaurant-menu/"

func GetMenuByRestaurauntId(restaurauntId string) (*MenuAPIResponse, error) {
	requestURL := tenbisEndpoint + restaurauntId + "_he" // TODO: change to var

	menuApi := new(MenuAPIResponse)
	err := util.GetJson(requestURL, menuApi)
	if err != nil {
		return &MenuAPIResponse{}, err
	}

	//formattedMenu := formatMenuFromAPI(menuApi);

	return menuApi, nil;
}
