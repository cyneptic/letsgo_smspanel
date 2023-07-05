package controllers

import (
	"net/http"

	"github.com/cyneptic/letsgo-smspanel/controller/validators"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
	"github.com/cyneptic/letsgo-smspanel/internal/core/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PhoneBookHandler struct {
	svc ports.PhoneBookServiceContract
}

func NewPhoneBookHandler() *PhoneBookHandler {
	svc := service.NewPhoneBookService()

	return &PhoneBookHandler{
		svc: svc,
	}
}

func AddPhoneBookRoutes(e *echo.Echo) {
	handler := NewPhoneBookHandler()
	e.POST("/phone-books", handler.CreatePhoneBookHandler)
	e.GET("/phone-books", handler.ListPhoneBooksHandler)
	e.GET("/phone-books/:id", handler.FindPhoneBookHandler)
	e.PATCH("/phone-books/:id", handler.UpdatePhoneBookHandler)
	e.DELETE("/phone-books/:id", handler.DeletePhoneBooksHandler)
}

func (h *PhoneBookHandler) CreatePhoneBookHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid UserID")
	}

	var phoneBookModel entities.PhoneBook
	if err := c.Bind(&phoneBookModel); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}
	phoneBookModel.UserId = userID

	err = validators.ValidatePhoneBookParam(phoneBookModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newPhoneBook, err := h.svc.CreatePhoneBookList(phoneBookModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, newPhoneBook)
}

func (h *PhoneBookHandler) ListPhoneBooksHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid UserID")
	}

	var phoneBookModel entities.PhoneBook
	phoneBookModel.UserId = userID

	var phoneBookList []entities.PhoneBook

	phoneBookList, err = h.svc.GetPhoneBookList(phoneBookModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, phoneBookList)
}

func (h *PhoneBookHandler) FindPhoneBookHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid UserID")
	}

	phoneBookID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}
	var phoneBookModel entities.PhoneBook
	phoneBookModel.ID = phoneBookID
	phoneBookModel.UserId = userID

	phoneBook, err := h.svc.GetPhoneBookById(phoneBookModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, phoneBook)
}

func (h *PhoneBookHandler) UpdatePhoneBookHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid UserID")
	}

	phoneBookID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	var phoneBookModel entities.PhoneBook
	if err := c.Bind(&phoneBookModel); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}
	phoneBookModel.UserId = userID
	phoneBookModel.ID = phoneBookID

	err = validators.ValidatePhoneBookParam(phoneBookModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updatedPhoneBook, err := h.svc.UpdatePhoneBookById(phoneBookModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, updatedPhoneBook)
}

func (h *PhoneBookHandler) DeletePhoneBooksHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid UserID")
	}

	phoneBookID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}
	var phoneBookModel entities.PhoneBook
	phoneBookModel.UserId = userID
	phoneBookModel.ID = phoneBookID

	if err := h.svc.DeletePhoneBookById(phoneBookModel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return nil
}
