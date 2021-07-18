package menu

import (
	"github.com/labstack/echo/v4"
)

func MenuController(e* echo.Echo) { 
	e.GET("/menu", getRandomDish)
}


