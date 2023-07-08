package controllers

import (
	"net/http"
	"strconv"

	"github.com/cyneptic/letsgo-smspanel/controller/validators"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	"github.com/cyneptic/letsgo-smspanel/internal/core/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AdminActionPriceRequest struct {
	UserID string `json:"user_id"`
	Price  string `json:"price"`
}

type AdminActionDisableUserRequest struct {
	UserID   string `json:"user_id"`
	TargetID string `json:"target_id"`
	Toggle   bool   `json:"toggle"`
}

type AdminActionGetHistoryRequest struct {
	UserID   string `json:"user_id"`
	TargetID string `json:"target_id"`
}

type AdminActionHandler struct {
	svc       ports.AdminActionServiceContract
	validator validators.AdminActionValidator
}

func NewAdminActionHandler() *AdminActionHandler {
	svc := service.NewAdminService()
	return &AdminActionHandler{
		svc:       svc,
		validator: validators.AdminActionValidator{},
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
	var request AdminActionPriceRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := h.validator.PriceValidator(request.UserID, request.Price)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userid, _ := uuid.Parse(request.UserID)
	price, _ := strconv.Atoi(request.Price)

	err = h.svc.EditSingleMessagePrice(userid, price)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *AdminActionHandler) EditGroupMessagePrice(c echo.Context) error {
	var request AdminActionPriceRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := h.validator.PriceValidator(request.UserID, request.Price)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userid, _ := uuid.Parse(request.UserID)
	price, _ := strconv.Atoi(request.Price)

	err = h.svc.EditGroupMessagePrice(userid, price)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *AdminActionHandler) DisableUserAccount(c echo.Context) error {
	var request AdminActionDisableUserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := h.validator.VerifyTwoUUID(request.UserID, request.TargetID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userid, _ := uuid.Parse(request.UserID)
	targetid, _ := uuid.Parse(request.TargetID)

	err = h.svc.DisableUserAccount(userid, targetid, request.Toggle)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *AdminActionHandler) GetUserHistory(c echo.Context) error {
	var request AdminActionGetHistoryRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := h.validator.VerifyTwoUUID(request.UserID, request.TargetID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userid, _ := uuid.Parse(request.UserID)
	targetid, _ := uuid.Parse(request.TargetID)

	messages, err := h.svc.GetUserHistory(userid, targetid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, messages)
}
