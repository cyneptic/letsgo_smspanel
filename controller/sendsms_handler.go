package controllers

import (
	"net/http"

	"github.com/cyneptic/letsgo-smspanel/controller/validators"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	"github.com/cyneptic/letsgo-smspanel/internal/core/service"
	"github.com/labstack/echo/v4"
)

type SendSMSHandler struct {
	svc ports.SendSMSServiceContract
}

func NewSnedSMSHandler() *SendSMSHandler {
	svc := service.NewSendSMSService()
	return &SendSMSHandler{
		svc: svc,
	}
}

func AddSendSMSRouters(e *echo.Echo) {
	handler := NewSnedSMSHandler()
	e.POST("/sendcontactlist", handler.SendToContactListHandler)
	e.POST("/senduser", handler.SendToUserHandler)
	e.POST("/sendnumber", handler.SendToNumberHandler)
	e.POST("/sendcontactlistinterval", handler.SendToContactListIntervalHandler)

}

// !SendToContactListHandler
func (h *SendSMSHandler) SendToContactListHandler(c echo.Context) error {
	var jsonMessage entities.MessageReciver
	err := c.Bind(&jsonMessage)
	if err != nil {
		return err
	}
	messageSample, err := validators.ValidatorReciveMessage(jsonMessage)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err = h.svc.SendToContactList(messageSample)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Fine")
}

// !SendToUserHandler
func (h *SendSMSHandler) SendToUserHandler(c echo.Context) error {
	var jsonMessage entities.MessageReciver
	err := c.Bind(&jsonMessage)
	if err != nil {
		return err
	}
	messageSample, err := validators.ValidatorReciveMessage(jsonMessage)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err = h.svc.SendToUser(messageSample)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Fine")
}

func (h *SendSMSHandler) SendToNumberHandler(c echo.Context) error {
	var jsonMessage entities.MessageReciver
	err := c.Bind(&jsonMessage)
	if err != nil {
		return err
	}
	messageSample, err := validators.ValidatorReciveMessage(jsonMessage)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err = h.svc.SendToNumber(messageSample)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Fine")
}

// !SendToContactListIntervalHandler
func (h *SendSMSHandler) SendToContactListIntervalHandler(c echo.Context) error {
	var jsonMessage entities.MessageReciver
	err := c.Bind(&jsonMessage)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	messageSample, err := validators.ValidatorReciveMessage(jsonMessage)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	t, err := validators.ValidateNumber(c.QueryParam("t"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = h.svc.SendToContactListInterval(messageSample, t)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Fine")

}
