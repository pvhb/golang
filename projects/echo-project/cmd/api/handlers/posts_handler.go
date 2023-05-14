package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pvhb/golang/projects/echo-project/cmd/api/service"
)

func GetPostsHandler(c echo.Context) error {
	data, err := service.GetAll()
	if err != nil {
		c.String(http.StatusBadGateway, err.Error())
	}

	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}

func GetPostByIDHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadGateway, err.Error())
	}

	data, err := service.GetByID(idx)
	if err != nil {
		return c.String(http.StatusBadGateway, err.Error())
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}
