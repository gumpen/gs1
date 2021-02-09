package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUsers(c echo.Context) error {
	users, err := h.UserRepository.GetAll()
	// users, err := repository.UserRepository.GetAll()
	if err != nil {
		return echo.NewHTTPError(500)
	}
	return c.JSON(http.StatusOK, users)
}
