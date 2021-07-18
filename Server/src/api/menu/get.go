package menu

import (
	"net/http"
	"random10bis/src/tenbisAPI"

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

	return c.JSON(http.StatusOK, menu)
}
