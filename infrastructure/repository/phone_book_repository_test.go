package repositories_test

import (
	"testing"

	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Set up the test environment by creating a repository and a test phone book
func setupPhoneBookTestEnvironment(t *testing.T) (*repositories.PGRepository, *entities.PhoneBook) {
	repo := repositories.NewGormDatabase()

	phoneBook := entities.PhoneBook{
		UserId: uuid.New(),
		Name:   "test phone book name",
	}

	// Create a test phone book in the repository
	err := repo.DB.Create(&phoneBook).Error
	assert.NoError(t, err)

	return repo, &phoneBook
}

// Clean up the test environment by deleting the created phone book
func cleanupPhoneBookTestEnvironment(t *testing.T, repo *repositories.PGRepository, phoneBook *entities.PhoneBook) {
	err := repo.DB.Unscoped().Delete(phoneBook).Error
	assert.NoError(t, err)
}

// Test creating a phone book
func TestCreatePhoneBook(t *testing.T) {
	t.Run("withValidData", func(t *testing.T) {

		repo, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, repo, phoneBook)

		newPhoneBook := entities.PhoneBook{
			UserId: phoneBook.UserId,
			Name:   "new phone book name",
		}

		// Create a new phone book in the repository
		pBook, err := repo.CreatePhoneBook(newPhoneBook)

		assert.NoError(t, err)
		assert.Equal(t, phoneBook.UserId, pBook.UserId)
		assert.Equal(t, "new phone book name", pBook.Name)

		// Delete the created phone book
		err = repo.DB.Unscoped().Delete(pBook).Error
		assert.NoError(t, err)
	})
}

// Test getting a list of phone books
func TestGetPhoneBookList(t *testing.T) {
	t.Run("withValidData", func(t *testing.T) {

		repo, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, repo, phoneBook)

		// Get the list of phone books from the repository and assert the results
		pBooks, err := repo.GetPhoneBookList(*phoneBook)
		assert.NoError(t, err)
		assert.NotEmpty(t, pBooks)
		assert.Len(t, pBooks, 1)
	})
}

// Test getting a phone book by its ID
func TestGetPhoneBookById(t *testing.T) {

	t.Run("withValidData", func(t *testing.T) {

		repo, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, repo, phoneBook)

		// Get the phone book by its ID and assert the results
		result, err := repo.GetPhoneBookById(*phoneBook)
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
		assert.Equal(t, phoneBook.Name, result.Name)
		assert.Equal(t, phoneBook.UserId, result.UserId)
	})

	t.Run("withInvalidPhoneBookID", func(t *testing.T) {

		repo, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, repo, phoneBook)

		phoneBookWithInvalidPhoneBookID := entities.PhoneBook{
			UserId: phoneBook.UserId,
		}
		phoneBookWithInvalidPhoneBookID.ID = uuid.New()

		// Get the phone book with an invalid phone book ID
		_, err := repo.GetPhoneBookById(phoneBookWithInvalidPhoneBookID)
		assert.Error(t, err)

	})
}

// Test updating a phone book by its ID
func TestUpdatePhoneBookById(t *testing.T) {
	t.Run("withValidData", func(t *testing.T) {

		repo, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, repo, phoneBook)

		// Update the name of the phone book
		phoneBook.Name = "updated phone book name"
		_, err := repo.UpdatePhoneBookById(*phoneBook)
		assert.NoError(t, err)

		// Get the updated phone book and assert the results
		result, err := repo.GetPhoneBookById(*phoneBook)
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
		assert.Equal(t, "updated phone book name", result.Name)
	})

	t.Run("withInvalidPhoneBookID", func(t *testing.T) {

		repo, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, repo, phoneBook)

		phoneBookWithInvalidPhoneBookID := entities.PhoneBook{
			UserId: phoneBook.UserId,
		}
		phoneBookWithInvalidPhoneBookID.ID = uuid.New()

		_, err := repo.UpdatePhoneBookById(phoneBookWithInvalidPhoneBookID)
		assert.Error(t, err)
	})
}

// Test deleting a phone book by its ID
func TestDeletePhoneBookById(t *testing.T) {
	t.Run("withValidData", func(t *testing.T) {

		repo, phoneBook := setupPhoneBookTestEnvironment(t)
		defer cleanupPhoneBookTestEnvironment(t, repo, phoneBook)

		err := repo.DeletePhoneBookById(*phoneBook)
		assert.NoError(t, err)

		// Verify that the phone book is deleted
		var pBook entities.PhoneBook
		err = repo.DB.Where("user_id = ? AND id = ?", phoneBook.UserId, phoneBook.ID).First(&pBook).Error
		assert.Error(t, err)
	})
}
