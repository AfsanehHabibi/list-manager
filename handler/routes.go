package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	lists := v1.Group("/lists")
	lists.GET("", h.getLists)
	lists.POST("", h.addNewList)
	lists.PUT("/:id", h.updateList)
	lists.DELETE("/:id", h.deleteList)
}