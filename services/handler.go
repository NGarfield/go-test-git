package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Servicer
}

func NerHandler(service Servicer) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetAllHandler(c echo.Context) error {
	test,err := h.service.GetAllService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, test)
}
