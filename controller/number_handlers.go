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
	numberGroup.GET("/generate", handler.GenerateNewNumber)
	numberGroup.GET("/buy", handler.BuyNumber)

}

func (h *NumberHandler) GenerateNewNumber(c echo.Context) error {
	generatedNumber, _ := h.srv.GenerateNumber()
	return c.String(http.StatusOK, generatedNumber)
}

func (h *NumberHandler) BuyNumber(c echo.Context) error {
	generatedNumber, _ := h.srv.GenerateNumber()
	h.srv.BuyNumber(uuid.New().String())
	return c.String(http.StatusOK, generatedNumber)
}
