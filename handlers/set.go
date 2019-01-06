package handlers

import (
	"github.com/labstack/echo"
	"go-fitness/api"
	"net/http"
)

func Set(service *api.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		schema, err := service.SetSchema()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		result := service.ExecuteQuery(c.QueryParam("query"), schema)
		return c.JSONPretty(http.StatusOK, result, "    ")
	}
}
