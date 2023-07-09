package controllers

import (
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	"github.com/cyneptic/letsgo-smspanel/internal/core/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type NumberHandler struct {
	srv ports.NumberServiceContract
}

func NewNumberHandler() *NumberHandler {
	return &NumberHandler{
		srv: service.NewNumberService(),
	}
}

func RegisterNumberHandler(ctx *echo.Echo) {
	handler := NewNumberHandler()
	numberGroup := ctx.Group("/api/number")
	numberGroup.GET("/buy", handler.BuyNumber)
	numberGroup.GET("/shared-number", handler.GetSharedNumber)
	numberGroup.GET("/subscribe", handler.SubscribeNumber)

}

func (h *NumberHandler) BuyNumber(c echo.Context) error {
	generatedNumber, _ := h.srv.GenerateNumber()
	err := h.srv.BuyNumber(uuid.New().String())
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, generatedNumber)
}

func (h *NumberHandler) GetSharedNumber(c echo.Context) error {
	shared, _ := h.srv.GetSharedNumber()
	return c.JSON(http.StatusOK, shared)
}

func (h *NumberHandler) SubscribeNumber(c echo.Context) error {
	generatedNumber, _ := h.srv.GenerateNumber()

	userId := c.Get("id").(string)

	err := h.srv.SubscribeNumber(userId, generatedNumber)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "success")
}
