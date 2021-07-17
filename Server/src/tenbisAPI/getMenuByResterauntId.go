package tenbisAPI

import (
	"random10bis/src/entities"
	"random10bis/src/util"
)

type menuAPIResponse struct {
	CategoriesList []category `json:"categoriesList"`
}

type category struct {
	DishList     []dish `json:"dishList"` 
}

type dish struct {
	DishName            string `json:"dishName"`
	IsPopularDish       bool `json:"isPopularDish"`
	DishPopularityScore int `json:"dishPopularityScore"`
}

const tenbisEndpoint = "https://tenbis-static.azureedge.net/restaurant-menu/"

func GetMenuByRestaurauntId(restaurauntId string) (entities.Menu, error) {
	requestURL := tenbisEndpoint + restaurauntId + "_he" // TODO: change to var

	menuApi := new(menuAPIResponse)
	err := util.GetJson(requestURL, menuApi)
	if err != nil {
		return entities.Menu{}, err
	}

	formattedMenu := formatMenuFromAPI(menuApi);

	return formattedMenu, nil;
}

func formatMenuFromAPI(menuFromApi *menuAPIResponse) entities.Menu {
	var dishes []string

	categories := menuFromApi.CategoriesList; 

	for _, category := range categories {
		for _, dish := range category.DishList {
			
			dishes = append(dishes, dish.DishName)
		}	
	}

	return entities.Menu{ DishNames: dishes }
}