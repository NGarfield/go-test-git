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
	test, err := h.service.GetAllService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, test)
}

func (h *Handler) InsertTestDataHandler(e echo.Context) error {
	var req Test
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	err := h.service.InsertTestDataService(&req)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, "Create Data Complete.")
}
