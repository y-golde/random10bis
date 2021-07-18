package randomDish

import (
	"math/rand"
	"net/http"
	"random10bis/src/entities"
	"random10bis/src/tenbisAPI"
	"time"

	"github.com/labstack/echo/v4"
)


func getRandomDish(c echo.Context) error {
	restaurantId := c.QueryParam("restaurantId")
	if restaurantId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing param restaurantId")
	}
	
	rawMenu, err := tenbisAPI.GetMenuByRestaurauntId(restaurantId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error getting restaurant menu, did you check if the restaurant exitsts?")
	}

	menu := formatMenuFromAPI(rawMenu);

	dishNames := menu.DishNames

	rand.Seed(time.Now().Unix())
	randomDishName := dishNames[rand.Intn(len(dishNames))]
	return c.String(http.StatusOK, randomDishName) 
}

func formatMenuFromAPI(menuFromApi *tenbisAPI.MenuAPIResponse) entities.Menu {
	var dishes []string

	categories := menuFromApi.CategoriesList; 

	for _, category := range categories {
		for _, dish := range category.DishList {
			dishes = append(dishes, dish.Name)
		}	
	}

	return entities.Menu{ DishNames: dishes }
}
