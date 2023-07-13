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

func NewSendSMSHandler() *SendSMSHandler {
	svc := service.NewSendSMSService()
	return &SendSMSHandler{
		svc: svc,
	}
}

func AddSendSMSRouters(e *echo.Echo) {
	// handler := NewSendSMSHandler()
	// e.POST("/send-contactlist", handler.SendToContactListHandler)
	// e.POST("/send-user", handler.SendToUserHandler)
	// e.POST("/send-number", handler.SendToNumberHandler)
	// e.POST("/send-contactlist-interval", handler.SendToContactListIntervalHandler)

}

// !SendToContactListHandler
func (h *SendSMSHandler) SendToContactListHandler(c echo.Context) error {
	var jsonMessage entities.MessageReciver
	err := c.Bind(&jsonMessage)
	if err != nil {
		return err
	}
	messageSample, err := validators.ValidatorReceiveMessage(jsonMessage)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err = h.svc.SendToContactList(messageSample)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Done")
}

// !SendToUserHandler
func (h *SendSMSHandler) SendToUserHandler(c echo.Context) error {
	var jsonMessage entities.MessageReciver
	err := c.Bind(&jsonMessage)
	if err != nil {
		return err
	}
	messageSample, err := validators.ValidatorReceiveMessage(jsonMessage)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err = h.svc.SendToUser(messageSample)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Done")
}

func (h *SendSMSHandler) SendToNumberHandler(c echo.Context) error {
	var jsonMessage entities.MessageReciver
	err := c.Bind(&jsonMessage)
	if err != nil {
		return err
	}
	messageSample, err := validators.ValidatorReceiveMessage(jsonMessage)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err = h.svc.SendToNumber(messageSample)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Done")
}

// !SendToContactListIntervalHandler
func (h *SendSMSHandler) SendToContactListIntervalHandler(c echo.Context) error {
	var jsonMessage entities.MessageReciver
	err := c.Bind(&jsonMessage)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	messageSample, err := validators.ValidatorReceiveMessage(jsonMessage)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	t, err := validators.ValidateTimeDuration(c.QueryParam("t"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = h.svc.SendToContactListInterval(messageSample, t)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Done")

}
