package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexPage(c echo.Context) error {
	return c.String(http.StatusOK, "welcome to index")
}
