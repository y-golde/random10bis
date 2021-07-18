package api

import (
	"random10bis/src/api/index"
	"random10bis/src/api/menu"
	"random10bis/src/api/randomDish"

	"github.com/labstack/echo/v4"
)

func RootController(e *echo.Echo) {
	menu.MenuController(e)
	index.IndexController(e)
	randomDish.RandomDishController(e)
}