package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	controllers "github.com/cyneptic/letsgo-smspanel/controller"
	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupPhoneBookTestEnvironment(t *testing.T) (*entities.User, *entities.PhoneBook) {
	// Set up the test environment by creating a user and a phone book
	repo := repositories.NewGormDatabase()

	user := entities.User{
		Name:        "test user",
		PhoneNumber: "09876543210",
		Email:       "test@domain.com",
	}

	err := repo.DB.Create(&user).Error
	assert.NoError(t, err)

	phoneBook := entities.PhoneBook{
		UserId: user.ID,
		Name:   "test phone book name",
	}

	err = repo.DB.Create(&phoneBook).Error
	assert.NoError(t, err)

	return &user, &phoneBook
}

func cleanupPhoneBookTestEnvironment(t *testing.T, user *entities.User, phoneBook *entities.PhoneBook) {
	// Clean up the test environment by deleting the user and the phone book
	repo := repositories.NewGormDatabase()
	err := repo.DB.Unscoped().Delete(phoneBook).Error
	assert.NoError(t, err)
	err = repo.DB.Unscoped().Delete(user).Error
	assert.NoError(t, err)
}

func TestCreatePhoneBookHandler(t *testing.T) {
	// Test the CreatePhoneBookHandler function

	prepareRequest := func(t *testing.T, userId string, bodyContent string) *httptest.ResponseRecorder {
		// Prepare the HTTP request for creating a phone book
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/phone-books", strings.NewReader(bodyContent))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Set("id", userId)
		handler := controllers.NewPhoneBookHandler()

		err := handler.CreatePhoneBookHandler(c)
		assert.NoError(t, err)
		return rec
	}

	t.Run("withValidData", func(t *testing.T) {
		// Set up the test environment
		user, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, user, phoneBook)

		bodyContent := `{"name": "new test phone book"}`

		// Make the request
		rec := prepareRequest(t, user.ID.String(), bodyContent)

		// Verify the response
		assert.Equal(t, http.StatusOK, rec.Code)

		var response entities.PhoneBook
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
		assert.Equal(t, "new test phone book", response.Name)
		assert.Equal(t, user.ID, response.UserId)

		// Clean up the test data
		repo := repositories.NewGormDatabase()
		err = repo.DB.Unscoped().Delete(response).Error
		assert.NoError(t, err)
	})

	t.Run("withInvalidUserID", func(t *testing.T) {
		// Set up the test environment
		user, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, user, phoneBook)

		bodyContent := `{"name": "new test phone book"}`
		invalidUserID := "string-that-is-not-a-UUID"

		// Make the request
		rec := prepareRequest(t, invalidUserID, bodyContent)

		// Verify the response
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("withEmptyName", func(t *testing.T) {
		// Set up the test environment
		user, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, user, phoneBook)

		bodyContent := `{"name": ""}`

		// Make the request
		rec := prepareRequest(t, user.ID.String(), bodyContent)

		// Verify the response
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestListPhoneBooksHandler(t *testing.T) {
	// Test the ListPhoneBooksHandler function

	prepareRequest := func(t *testing.T, userId string) *httptest.ResponseRecorder {
		// Prepare the HTTP request for listing phone books
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/phone-books", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Set("id", userId)
		handler := controllers.NewPhoneBookHandler()

		err := handler.ListPhoneBooksHandler(c)
		assert.NoError(t, err)
		return rec
	}

	t.Run("withValidData", func(t *testing.T) {
		// Set up the test environment
		user, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, user, phoneBook)

		// Make the request
		rec := prepareRequest(t, user.ID.String())

		// Verify the response
		assert.Equal(t, http.StatusOK, rec.Code)

		var responseBody []entities.PhoneBook
		err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.NotEmpty(t, responseBody)
		assert.Len(t, responseBody, 1)
		assert.Equal(t, "test phone book name", responseBody[0].Name)
		assert.Equal(t, user.ID, responseBody[0].UserId)
	})

	t.Run("withInvalidUserID", func(t *testing.T) {
		// Set up the test environment
		user, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, user, phoneBook)

		invalidUserID := "string-that-is-not-a-UUID"

		// Make the request
		rec := prepareRequest(t, invalidUserID)

		// Verify the response
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestFindPhoneBookHandler(t *testing.T) {
	// Test the FindPhoneBookHandler function

	prepareRequest := func(t *testing.T, userId, phoneBookId string) *httptest.ResponseRecorder {
		// Prepare the HTTP request for finding a phone book
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/phone-books", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Set("id", userId)
		c.SetParamNames("id")
		c.SetParamValues(phoneBookId)

		handler := controllers.NewPhoneBookHandler()

		err := handler.FindPhoneBookHandler(c)
		assert.NoError(t, err)
		return rec
	}

	t.Run("withValidData", func(t *testing.T) {
		// Set up the test environment
		user, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, user, phoneBook)

		// Make the request
		rec := prepareRequest(t, user.ID.String(), phoneBook.ID.String())

		// Verify the response
		assert.Equal(t, http.StatusOK, rec.Code)

		var responseBody entities.PhoneBook
		err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.NotEmpty(t, responseBody)
		assert.Equal(t, phoneBook.Name, responseBody.Name)
		assert.Equal(t, user.ID, responseBody.UserId)
	})

	t.Run("withInvalidUserID", func(t *testing.T) {
		// Set up the test environment
		user, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, user, phoneBook)

		invalidUserID := uuid.NewString()

		// Make the request
		rec := prepareRequest(t, invalidUserID, phoneBook.ID.String())

		// Verify the response
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("withInvalidPhoneBookID", func(t *testing.T) {
		// Set up the test environment
		user, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, user, phoneBook)

		invalidPhoneBookID := "string-that-is-not-a-UUID"

		// Make the request
		rec := prepareRequest(t, user.ID.String(), invalidPhoneBookID)

		// Verify the response
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestUpdatePhoneBookHandler(t *testing.T) {
	// Test the UpdatePhoneBookHandler function

	prepareRequest := func(t *testing.T, userId, phoneBookId string, bodyContent string) *httptest.ResponseRecorder {
		// Prepare the HTTP request for updating a phone book
		e := echo.New()
		req := httptest.NewRequest(http.MethodPatch, "/phone-books", strings.NewReader(bodyContent))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Set("id", userId)
		c.SetParamNames("id")
		c.SetParamValues(phoneBookId)

		handler := controllers.NewPhoneBookHandler()

		err := handler.UpdatePhoneBookHandler(c)
		assert.NoError(t, err)
		return rec
	}

	t.Run("withValidData", func(t *testing.T) {
		// Set up the test environment
		user, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, user, phoneBook)

		bodyContent := `{"name": "edit phone book"}`

		// Make the request
		rec := prepareRequest(t, user.ID.String(), phoneBook.ID.String(), bodyContent)

		// Verify the response
		assert.Equal(t, http.StatusOK, rec.Code)

		var responseBody entities.PhoneBook
		err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.NotEmpty(t, responseBody)
		assert.Equal(t, "edit phone book", responseBody.Name)
		assert.Equal(t, user.ID, responseBody.UserId)
	})

	t.Run("withInvalidUserID", func(t *testing.T) {
		// Set up the test environment
		user, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, user, phoneBook)

		invalidUserID := uuid.NewString()
		bodyContent := `{"name": "edit phone book"}`

		// Make the request
		rec := prepareRequest(t, invalidUserID, phoneBook.ID.String(), bodyContent)

		// Verify the response
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("withInvalidPhoneBookID", func(t *testing.T) {
		// Set up the test environment
		user, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, user, phoneBook)

		bodyContent := `{"name": "edit phone book"}`
		invalidPhoneBookID := "string-that-is-not-a-UUID"

		// Make the request
		rec := prepareRequest(t, user.ID.String(), invalidPhoneBookID, bodyContent)

		// Verify the response
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("withEmptyName", func(t *testing.T) {
		// Set up the test environment
		user, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, user, phoneBook)

		bodyContent := `{"name": ""}`

		// Make the request
		rec := prepareRequest(t, user.ID.String(), phoneBook.ID.String(), bodyContent)

		// Verify the response
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestDeletePhoneBooksHandler(t *testing.T) {
	// Test the DeletePhoneBooksHandler function

	prepareRequest := func(t *testing.T, userId, phoneBookId string) *httptest.ResponseRecorder {
		// Prepare the HTTP request for deleting a phone book
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/phone-books", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Set("id", userId)
		c.SetParamNames("id")
		c.SetParamValues(phoneBookId)

		handler := controllers.NewPhoneBookHandler()

		err := handler.DeletePhoneBooksHandler(c)
		assert.NoError(t, err)
		return rec
	}

	t.Run("withValidData", func(t *testing.T) {
		// Set up the test environment
		user, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, user, phoneBook)

		// Make the request
		rec := prepareRequest(t, user.ID.String(), phoneBook.ID.String())

		// Verify the response
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("withInvalidUserID", func(t *testing.T) {
		// Set up the test environment
		user, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, user, phoneBook)

		invalidUserID := "string-that-is-not-a-UUID"

		// Make the request
		rec := prepareRequest(t, invalidUserID, phoneBook.ID.String())

		// Verify the response
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("withInvalidPhoneBookID", func(t *testing.T) {
		// Set up the test environment
		user, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, user, phoneBook)

		invalidPhoneBookID := "string-that-is-not-a-UUID"

		// Make the request
		rec := prepareRequest(t, user.ID.String(), invalidPhoneBookID)

		// Verify the response
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
