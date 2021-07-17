package index

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func postIndex(c echo.Context) error {
	return c.String(http.StatusOK, "post index")
}