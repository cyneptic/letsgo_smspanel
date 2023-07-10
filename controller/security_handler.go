package controllers

import (
	"net/http"

	"github.com/cyneptic/letsgo-smspanel/controller/validators"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	"github.com/cyneptic/letsgo-smspanel/internal/core/service"
	"github.com/labstack/echo/v4"
)

type SecurityHnadler struct {
	svc ports.SecurityServiceContract
}

func NewSecurityHandler() *SecurityHnadler {
	svc := service.NewSecurityService()
	return &SecurityHnadler{
		svc: svc,
	}
}
func AddSecurityRoutes(e *echo.Echo) {
	handler := NewSecurityHandler()
	e.GET("/search-message", handler.searchMessageHandler)
	e.POST("/add-blaclist-word", handler.addWordToBlackListHandler)
	e.DELETE("/remove-blacklist-word", handler.removeWordFromBlackListHandler)
	e.POST("/add-blacklist-regex", handler.addRegexToBlackListHandler)
	e.DELETE("/remove-blacklist-regex", handler.removeRegexFromBlackListHandler)
}
func (h *SecurityHnadler) searchMessageHandler(c echo.Context) error {
	word := c.QueryParam("word")
	err := validators.SearchMessageValidator(word)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	messages, err := h.svc.SearchInMessages(word)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, messages)
}

// done
func (h *SecurityHnadler) addWordToBlackListHandler(c echo.Context) error {

	newWord := new(entities.BlacklistWord)
	err := c.Bind(&newWord)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	err = validators.AddWordToBlackListValidator(*newWord)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = h.svc.AddWordToBlackList(*newWord)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "New Word Added To BlackList")
}

// done
func (h *SecurityHnadler) removeWordFromBlackListHandler(c echo.Context) error {
	word := c.QueryParam("word")
	err := validators.RemoveWordFromBlackListValidator(word)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	err = h.svc.RemoveWordFromBlackList(word)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "word removed from blackList")
}

// done
func (h *SecurityHnadler) addRegexToBlackListHandler(c echo.Context) error {
	newRegex := new(entities.BlacklistRegex)
	err := c.Bind(&newRegex)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	err = validators.AddRegexToBlackListValidator(*newRegex)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	err = h.svc.AddRegexToBlackList(*newRegex)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "New Regex Added To BlackList")
}

// BUG HERE    //
func (h *SecurityHnadler) removeRegexFromBlackListHandler(c echo.Context) error {
	regex := c.QueryParam("regex")

	err := h.svc.RemoveRegexFromBlackList(regex)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "regex removed from BlackList")
}
