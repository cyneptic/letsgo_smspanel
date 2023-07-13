package controllers

import (
	"fmt"
	"github.com/cyneptic/letsgo-smspanel/controller/validators"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	"github.com/cyneptic/letsgo-smspanel/internal/core/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type WalletHandler struct {
	svc ports.WalletServiceContracts
}

func NewWalletHandler() *WalletHandler {
	svc := service.NewWalletService()
	return &WalletHandler{
		svc: svc,
	}
}
func AddWalletHRoutes(e *echo.Echo) {
	walletHandler := NewWalletHandler()
	e.POST("/wallet-create", walletHandler.GenerateWalletHandler)

	e.GET("/wallet-charge", walletHandler.ChargeWalletHandler)
	e.POST("/wallet-charge-verify", walletHandler.VerifyChargeWalletHandler)
	e.GET("/wallet-amount", walletHandler.GetWalletAmountHandler)
	e.GET("/wallet-has-amount", walletHandler.HasWalletCreditHandler)
}

func (w *WalletHandler) GenerateWalletHandler(c echo.Context) error {
	paramID := c.QueryParam("id")
	err := validators.GenerateWalletValidation(paramID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userID, err := uuid.Parse(paramID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = w.svc.GenerateWalletService(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "wallet Generate successfully")
}

func (w *WalletHandler) ChargeWalletHandler(c echo.Context) error {

	userID := c.QueryParam("userId")
	userUUId, err := uuid.Parse(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	walletID := c.QueryParam("walletId")
	walletUUId, err := uuid.Parse(walletID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	amountStr := c.QueryParam("amount")

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		// Handle the error if the conversion fails
		return c.JSON(http.StatusBadRequest, "Invalid amount")
	}
	err = validators.ChargeWalletValidation(userID, walletID, amountStr, amount)
	if err != nil {
		// Handle the error if the conversion fails
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	redirectLink, err := w.svc.ChargeWalletService(walletUUId, userUUId, amount)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(http.StatusTemporaryRedirect, redirectLink)
}

func (w *WalletHandler) VerifyChargeWalletHandler(c echo.Context) error {
	refID := c.FormValue("RefId")
	SaleReferenceId := c.FormValue("SaleReferenceId")
	reservationId := c.FormValue("SaleOrderId")
	walletUUId, err := uuid.Parse(reservationId)
	userId := c.Get("id").(string)
	userUUId, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := w.svc.ChargeWalletServiceVerify(refID, walletUUId, userUUId, SaleReferenceId)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, fmt.Sprintf("%v", result))

}

func (w *WalletHandler) GetWalletAmountHandler(c echo.Context) error {
	walletID := c.QueryParam("walletId")
	walletUUId, err := uuid.Parse(walletID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userID := c.QueryParam("userId")
	userUUId, err := uuid.Parse(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = validators.GetWalletAmountValidation(userID, walletID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	amount, err := w.svc.GetWalletAmountService(walletUUId, userUUId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, amount)
}
func (w *WalletHandler) HasWalletCreditHandler(c echo.Context) error {
	walletID := c.QueryParam("walletId")
	amountStr := c.QueryParam("amount")
	err := validators.HasWalletCredit(walletID, amountStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	walletUUId, err := uuid.Parse(walletID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		// Handle the error if the conversion fails
		return c.JSON(http.StatusBadRequest, "Invalid amount")
	}
	hasEnoughAmount, err := w.svc.HasEnoughCreditService(walletUUId, amount)
	if err != nil {
		// Handle the error if the conversion fails
		return c.JSON(http.StatusBadRequest, "Invalid amount")
	}
	return c.JSON(http.StatusOK, hasEnoughAmount)
}
