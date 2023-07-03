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

type ContactHandler struct {
	svc ports.ContactServiceContract
}

func NewContactHandler() *ContactHandler {
	svc := service.NewContactService()

	return &ContactHandler{
		svc: svc,
	}
}
func AddContactRoutes(e *echo.Echo) {
	handler := NewContactHandler()
	e.POST("/contacts", handler.CreateContactByUsernameHandler)
	e.GET("/contacts", handler.FindContactByUsernameHandler)
	e.PATCH("/phone-books/:id/contacts/:contact-id", handler.UpdateContactByUsernameHandler)
	e.DELETE("/phone-books/:id/contacts/:contact-id", handler.DeleteContacByUsernametHandler)
	e.POST("/phone-books/:id/contacts", handler.CreateContactHandler)
	e.GET("/phone-books/:id/contacts", handler.ListContactsHandler)
	e.GET("/phone-books/:id/contacts/:contact-id", handler.FindContactHandler)
	e.PATCH("/phone-books/:id/contacts/:contact-id", handler.UpdateContactHandler)
	e.DELETE("/phone-books/:id/contacts/:contact-id", handler.DeleteContactHandler)
}

func (h *ContactHandler) CreateContactByUsernameHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid UserID")
	}

	var contactModel entities.Contact
	if err := c.Bind(&contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}
	contactModel.UserID = userID
	contactModel.PhoneBookID = uuid.Nil

	if err := validators.ValidateUsername(contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := validators.ValidateContactParam(contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	contact, err := h.svc.CreateContact(contactModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, contact)
}
func (h *ContactHandler) FindContactByUsernameHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid UserID")
	}

	var contactModel entities.Contact
	if err := c.Bind(&contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}

	if err := validators.ValidateUsername(contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	contactModel.UserID = userID

	contact, err := h.svc.GetContactByUsername(contactModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, contact)
}
func (h *ContactHandler) UpdateContactByUsernameHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid UserID")
	}

	var contactModel entities.Contact
	if err := c.Bind(&contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}

	if err := validators.ValidateUsername(contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	contactModel.UserID = userID

	if err := validators.ValidateUpdateContactParam(contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	contact, err := h.svc.UpdateContactByUsername(contactModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, contact)
}
func (h *ContactHandler) DeleteContacByUsernametHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid UserID")
	}

	var contactModel entities.Contact
	if err := c.Bind(&contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}

	if err := validators.ValidateUsername(contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	contactModel.UserID = userID

	if err := h.svc.DeleteContactByUsername(contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return nil
}
func (h *ContactHandler) CreateContactHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid UserID")
	}

	phoneBookID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	var contactModel entities.Contact
	if err := c.Bind(&contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}
	contactModel.UserID = userID
	contactModel.PhoneBookID = phoneBookID

	if err := validators.ValidateContactParam(contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	contact, err := h.svc.CreateContact(contactModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, contact)
}
func (h *ContactHandler) ListContactsHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid UserID")
	}

	phoneBookID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	var contactModel entities.Contact
	contactModel.UserID = userID
	contactModel.PhoneBookID = phoneBookID

	contactList, err := h.svc.GetContactList(contactModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, contactList)
}
func (h *ContactHandler) FindContactHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid UserID")
	}

	phoneBookID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	contactID, err := uuid.Parse(c.Param("contact-id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}
	var contactModel entities.Contact
	contactModel.UserID = userID
	contactModel.PhoneBookID = phoneBookID
	contactModel.ID = contactID

	contact, err := h.svc.GetContactById(contactModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, contact)
}
func (h *ContactHandler) UpdateContactHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid UserID")
	}

	phoneBookID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	contactID, err := uuid.Parse(c.Param("contact-id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	var contactModel entities.Contact
	if err := c.Bind(&contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}
	contactModel.UserID = userID
	contactModel.ID = contactID
	contactModel.PhoneBookID = phoneBookID

	if err := validators.ValidateUpdateContactParam(contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	contact, err := h.svc.UpdateContactById(contactModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, contact)
}
func (h *ContactHandler) DeleteContactHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid UserID")
	}

	phoneBookID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	contactID, err := uuid.Parse(c.Param("contact-id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}
	var contactModel entities.Contact
	contactModel.UserID = userID
	contactModel.ID = contactID
	contactModel.PhoneBookID = phoneBookID

	if err := h.svc.DeleteContactById(contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return nil
}
