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
	e.GET("/contacts", handler.ListContactByUsernameHandler)
	e.GET("/contacts/:username", handler.FindContactByUsernameHandler)
	e.PATCH("contacts/:username", handler.UpdateContactByUsernameHandler)
	e.DELETE("contacts/:username", handler.DeleteContacByUsernametHandler)
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
		return c.JSON(http.StatusBadRequest, "invalid user id")
	}

	var contactModel entities.Contact
	if err := c.Bind(&contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}
	contactModel.UserID = userID
	contactModel.PhoneBookID = uuid.Nil

	if err := validators.ValidateContactByUsernameParam(contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	contact, err := h.svc.CreateContactByUsername(contactModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, contact)
}

func (h *ContactHandler) ListContactByUsernameHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid user id")
	}

	var contactModel entities.Contact
	contactModel.UserID = userID

	contacts, err := h.svc.ListContactByUsername(contactModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, contacts)
}

func (h *ContactHandler) FindContactByUsernameHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid user id")
	}

	var contactModel entities.Contact
	contactModel.UserID = userID
	contactModel.Username = c.Param("username")

	if err := validators.ValidateUsername(contactModel.Username); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

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
		return c.JSON(http.StatusBadRequest, "invalid user id")
	}

	username := c.Param("username")

	if err := validators.ValidateUsername(username); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var contactModel entities.Contact
	contactModel.UserID = userID

	if err := c.Bind(&contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}

	if err := validators.ValidateUpdateContactParam(contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	contact, err := h.svc.UpdateContactByUsername(username, contactModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, contact)
}

func (h *ContactHandler) DeleteContacByUsernametHandler(c echo.Context) error {
	userId := c.Get("id").(string)

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid user id")
	}

	var contactModel entities.Contact
	contactModel.UserID = userID
	contactModel.Username = c.Param("username")

	if err := validators.ValidateUsername(contactModel.Username); err != nil {
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
		return c.JSON(http.StatusBadRequest, "invalid user id")
	}

	phoneBookID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid phone book id")
	}

	var contactModel entities.Contact
	if err := c.Bind(&contactModel); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}
	contactModel.UserID = userID
	contactModel.PhoneBookID = phoneBookID

	if err := validators.ValidateContact(contactModel); err != nil {
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
		return c.JSON(http.StatusBadRequest, "invalid user id")
	}

	phoneBookID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid phone book id")
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
		return c.JSON(http.StatusBadRequest, "invalid user id")
	}

	phoneBookID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid phone book id")
	}

	contactID, err := uuid.Parse(c.Param("contact-id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid contact id")
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
		return c.JSON(http.StatusBadRequest, "invalid user id")
	}

	phoneBookID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid phone book id")
	}

	contactID, err := uuid.Parse(c.Param("contact-id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid contact id")
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
		return c.JSON(http.StatusBadRequest, "invalid user id")
	}

	phoneBookID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid phone book id")
	}

	contactID, err := uuid.Parse(c.Param("contact-id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid contact id")
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
