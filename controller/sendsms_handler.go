package controller

import (
	"net/http"

	"github.com/cyneptic/letsgo-smspanel/infrastructure/provider"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	"github.com/cyneptic/letsgo-smspanel/internal/core/service"
	"github.com/labstack/echo/v4"
)

type SendSMSHandler struct {
	svc ports.SendSMSServiceContract
}

func NewSnedSMSHandler() *SendSMSHandler {
	pv := provider.NewSendSMSProviderClient()
	svc := service.NewSendSMSService(pv)
	return &SendSMSHandler{
		svc: svc,
	}
}

func AddSendSMSRouters(e *echo.Echo) {
	handler := NewSnedSMSHandler()
	e.GET("/sendtocontactlist", handler.SendToContactListhandler)
}

func (h *SendSMSHandler) SendToContactListhandler(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := c.Bind(&json_map)
	if err != nil {
		return err
	} else {
		message := json_map["message"]
		return c.JSON(http.StatusOK, message)
	}
}
