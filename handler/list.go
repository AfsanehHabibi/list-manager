package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) getLists(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not yet implemented")
}

func (h *Handler) updateList(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not yet implemented")
}

func (h *Handler) addNewList(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not yet implemented")
}

func (h *Handler) deleteList(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not yet implemented")
}