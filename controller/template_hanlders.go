package controllers

import (
	"net/http"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	"github.com/cyneptic/letsgo-smspanel/internal/core/service"
	"github.com/labstack/echo/v4"
)

type TemplateHandler struct {
	svc ports.TemplateContract
}

func NewTemplateHandler() *TemplateHandler {
	svc := service.NewTemplateService()
	return &TemplateHandler{
		svc: svc,
	}
}

func AddTemplateRoutes(e *echo.Echo) {
	handler := NewTemplateHandler()
	e.POST("/create-temp", handler.CreateTemplateHandler)
	e.POST("/create-temp-content", handler.CreateTemplateContentHandler)
}

func (h *TemplateHandler) CreateTemplateHandler(c echo.Context) error {
	var jsonTemp entities.Template
	err := c.Bind(&jsonTemp)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err = h.svc.CreateTemplate(jsonTemp)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, jsonTemp)
}
func (h *TemplateHandler) CreateTemplateContentHandler(c echo.Context) error {
	return nil
}
