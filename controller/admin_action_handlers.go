package controllers

import (
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	"github.com/cyneptic/letsgo-smspanel/internal/core/service"
	"github.com/labstack/echo/v4"
)

type AdminActionHandler struct {
	svc ports.AdminActionServiceContract
}

func NewAdminActionHandler() *AdminActionHandler {
	svc := service.NewAdminService()
	return &AdminActionHandler{
		svc: svc,
	}
}

func AddAdminActionRoutes(e *echo.Echo) {
	handler := NewAdminActionHandler()
	e.POST("/edit-single-price", handler.EditSingleMessagePrice)
	e.POST("/edit-group-price", handler.EditGroupMessagePrice)
	e.POST("/disable-user", handler.DisableUserAccount)
	e.POST("/get-user-history", handler.GetUserHistory)
}

func (h *AdminActionHandler) EditSingleMessagePrice(c echo.Context) error {

	return nil
}

func (h *AdminActionHandler) EditGroupMessagePrice(c echo.Context) error {

	return nil
}

func (h *AdminActionHandler) DisableUserAccount(c echo.Context) error {

	return nil
}

func (h *AdminActionHandler) GetUserHistory(c echo.Context) error {

	return nil
}
