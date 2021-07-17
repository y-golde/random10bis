package randomDish

import (
	"github.com/labstack/echo/v4"
)

func RandomDishController(e* echo.Echo) { 
	e.GET("/randomDish", getRandomDish)
}


