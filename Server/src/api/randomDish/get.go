package randomDish

import (
	"math/rand"
	"net/http"
	"random10bis/src/tenbisAPI"
	"time"

	"github.com/labstack/echo/v4"
)


func getRandomDish(c echo.Context) error {
	restaurantId := c.QueryParam("restaurantId")
	if restaurantId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing param restaurantId")
	}
	
	menu, err := tenbisAPI.GetMenuByRestaurauntId(restaurantId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error getting restaurant menu, did you check if the restaurant exitsts?")
	}

	dishNames := menu.DishNames

	rand.Seed(time.Now().Unix())
	randomDishName := dishNames[rand.Intn(len(dishNames))]
	// TODO : verify params
	return c.String(http.StatusOK, randomDishName) 
}


