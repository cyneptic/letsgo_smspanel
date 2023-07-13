package controllers

import (
	"errors"
	"github.com/cyneptic/letsgo-smspanel/controller/validators"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	"github.com/cyneptic/letsgo-smspanel/internal/core/service"
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
	ok, id := validators.UUIDValidation(c.Get("id").(string))
	if !ok {
		return errors.New("invalid user id")
	}
	generatedNumber, _ := h.srv.GenerateNumber()
	err := h.srv.BuyNumber(id, generatedNumber)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"number": generatedNumber,
	},
	)
}

func (h *NumberHandler) GetSharedNumber(c echo.Context) error {
	shared, _ := h.srv.GetSharedNumber()
	return c.JSON(http.StatusOK, echo.Map{
		"shared_numbers": shared,
	},
	)
}

func (h *NumberHandler) SubscribeNumber(c echo.Context) error {
	ok, id := validators.UUIDValidation(c.Get("id").(string))
	if !ok {
		return errors.New("invalid user id")
	}
	generatedNumber, err := h.srv.GenerateNumber()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = h.srv.SubscribeNumber(id, generatedNumber)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "success")
}
