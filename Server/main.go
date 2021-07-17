package main

import (
	"fmt"

	"random10bis/src/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	fmt.Printf("Hello, World!")

	e := echo.New()
	
	e.Use(middleware.Logger())

	api.RootController(e)

	e.Logger.Fatal(e.Start(":9000"))
}