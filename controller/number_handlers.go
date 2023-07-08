package controllers

import (
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	"github.com/cyneptic/letsgo-smspanel/internal/core/service"
	"github.com/labstack/echo/v4"
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
}

func (h *NumberHandler) GenerateNewNumber(c echo.Context) error {
	return nil
}
