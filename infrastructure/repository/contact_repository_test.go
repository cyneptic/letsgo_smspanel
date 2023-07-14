package repositories_test

import (
	"testing"

	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
)

func setupContactTestEnvironment(t *testing.T) (*repositories.PGRepository, *entities.Contact) {
	repo := repositories.NewGormDatabase()

	contact := entities.Contact{
		UserID:      uuid.New(),
		FirstName:   "Amir",
		LastName:    "Dashti",
		Username:    random.String(15),
		PhoneNumber: "19123456789",
	}

	err := repo.DB.Create(&contact).Error
	assert.NoError(t, err)

	return repo, &contact
}

func cleanupContactTestEnvironment(t *testing.T, repo *repositories.PGRepository, contact *entities.Contact) {
	err := repo.DB.Unscoped().Delete(contact).Error
	assert.NoError(t, err)
}

func TestCreateContactByUsername(t *testing.T) {
	t.Run("withValidData", func(t *testing.T) {
		repo, contact := setupContactTestEnvironment(t)
		defer cleanupContactTestEnvironment(t, repo, contact)

		newContact := entities.Contact{
			UserID:      contact.UserID,
			FirstName:   "test name",
			LastName:    "test family",
			Username:    random.String(15),
			PhoneNumber: "19120000000",
		}

		createdContact, err := repo.CreateContactByUsername(newContact)

		assert.NoError(t, err)
		assert.Equal(t, contact.UserID, createdContact.UserID)
		assert.Equal(t, newContact.FirstName, createdContact.FirstName)
		assert.Equal(t, newContact.LastName, createdContact.LastName)
		assert.Equal(t, newContact.Username, createdContact.Username)
		assert.Equal(t, newContact.PhoneNumber, createdContact.PhoneNumber)

		err = repo.DB.Unscoped().Delete(createdContact).Error
		assert.NoError(t, err)
	})

	t.Run("withExistsUsername", func(t *testing.T) {
		repo, contact := setupContactTestEnvironment(t)
		defer cleanupContactTestEnvironment(t, repo, contact)

		newContact := entities.Contact{
			UserID:      contact.UserID,
			FirstName:   "test",
			LastName:    "test",
			Username:    contact.Username,
			PhoneNumber: "19120000000",
		}

		_, err := repo.CreateContactByUsername(newContact)
		assert.Error(t, err)
	})
}

func TestListContactByUsername(t *testing.T) {
	repo, contact := setupContactTestEnvironment(t)
	defer cleanupContactTestEnvironment(t, repo, contact)

	Contactlist, err := repo.ListContactByUsername(*contact)

	assert.NoError(t, err)
	assert.NotEmpty(t, Contactlist)
	assert.Len(t, Contactlist, 1)
	assert.Equal(t, contact.Username, Contactlist[0].Username)
}

func TestGetContactByUsername(t *testing.T) {
	t.Run("withValidData", func(t *testing.T) {
		repo, contact := setupContactTestEnvironment(t)
		defer cleanupContactTestEnvironment(t, repo, contact)

		result, err := repo.GetContactByUsername(*contact)
		assert.NoError(t, err)
		assert.Equal(t, contact.UserID, result.UserID)
		assert.Equal(t, contact.FirstName, result.FirstName)
		assert.Equal(t, contact.LastName, result.LastName)
		assert.Equal(t, contact.Username, result.Username)
		assert.Equal(t, contact.PhoneNumber, result.PhoneNumber)
	})

	t.Run("withInvalidUsername", func(t *testing.T) {
		repo, contact := setupContactTestEnvironment(t)
		defer cleanupContactTestEnvironment(t, repo, contact)

		var contactWithInvalidUsername entities.Contact
		contactWithInvalidUsername.Username = random.String(14)
		_, err := repo.GetContactByUsername(contactWithInvalidUsername)
		assert.Error(t, err)
	})
}

func TestUpdateContactByUsername(t *testing.T) {

	t.Run("withValidData", func(t *testing.T) {
		repo, contact := setupContactTestEnvironment(t)
		defer cleanupContactTestEnvironment(t, repo, contact)
		editedContact := entities.Contact{
			UserID:      contact.UserID,
			Username:    random.String(14),
			FirstName:   "Amir Hossein",
			LastName:    "-T",
			PhoneNumber: "19100000000",
		}

		updatedContact, err := repo.UpdateContactByUsername(contact.Username, editedContact)
		assert.NoError(t, err)
		assert.Equal(t, editedContact.UserID, updatedContact.UserID)
		assert.Equal(t, editedContact.Username, updatedContact.Username)
		assert.Equal(t, editedContact.FirstName, updatedContact.FirstName)
		assert.Equal(t, editedContact.LastName, updatedContact.LastName)
		assert.Equal(t, editedContact.PhoneNumber, updatedContact.PhoneNumber)
	})

	t.Run("withSomeData", func(t *testing.T) {
		repo, contact := setupContactTestEnvironment(t)
		defer cleanupContactTestEnvironment(t, repo, contact)

		editedContact := entities.Contact{
			UserID:    contact.UserID,
			FirstName: "alireza",
		}

		updatedContact, err := repo.UpdateContactByUsername(contact.Username, editedContact)
		assert.NoError(t, err)
		assert.Equal(t, editedContact.FirstName, updatedContact.FirstName)
	})

	t.Run("withInvalidUsername", func(t *testing.T) {
		repo, contact := setupContactTestEnvironment(t)
		defer cleanupContactTestEnvironment(t, repo, contact)

		editedContact := entities.Contact{
			UserID:      contact.UserID,
			Username:    random.String(15),
			FirstName:   "Amir Hossein",
			LastName:    "-T",
			PhoneNumber: "19100000000",
		}

		invalidUsername := random.String(12)
		_, err := repo.UpdateContactByUsername(invalidUsername, editedContact)
		assert.Error(t, err)

	})

	t.Run("withUsernameAlreadyExists", func(t *testing.T) {
		repo, contact := setupContactTestEnvironment(t)
		defer cleanupContactTestEnvironment(t, repo, contact)

		newContact := entities.Contact{
			UserID:      contact.UserID,
			FirstName:   "test name",
			Username:    random.String(16),
			PhoneNumber: "19120000000",
		}

		err := repo.DB.Create(&newContact).Error
		assert.NoError(t, err)

		contactwithUsernameAlreadyExists := entities.Contact{
			UserID:   contact.UserID,
			Username: newContact.Username,
		}

		_, err = repo.UpdateContactByUsername(contact.Username, contactwithUsernameAlreadyExists)
		assert.Error(t, err)

		err = repo.DB.Unscoped().Delete(newContact).Error
		assert.NoError(t, err)
	})
}

func TestDeleteContactByUsername(t *testing.T) {
	t.Run("withValidData", func(t *testing.T) {
		repo, contact := setupContactTestEnvironment(t)
		defer cleanupContactTestEnvironment(t, repo, contact)

		err := repo.DeleteContactByUsername(*contact)
		assert.NoError(t, err)

		var deletedContact entities.Contact
		err = repo.DB.Where("user_id = ? AND username = ?", contact.UserID, contact.Username).First(&deletedContact).Error
		assert.Error(t, err)
	})

	t.Run("withInvalidUsername", func(t *testing.T) {
		repo, contact := setupContactTestEnvironment(t)
		defer cleanupContactTestEnvironment(t, repo, contact)

		contactwithInvalidUsername := entities.Contact{
			UserID:   contact.UserID,
			Username: random.String(13),
		}

		err := repo.DeleteContactByUsername(contactwithInvalidUsername)
		assert.NoError(t, err)

		var deletedContact entities.Contact
		err = repo.DB.Where("user_id = ? AND username = ?", contact.UserID, contact.Username).First(&deletedContact).Error
		assert.NoError(t, err)
		assert.Equal(t, contact.Username, deletedContact.Username)
	})
}

func setupContactInPhoneBookTestEnvironment(t *testing.T) (*repositories.PGRepository, *entities.PhoneBook, *entities.Contact) {
	repo := repositories.NewGormDatabase()

	phoneBook := entities.PhoneBook{
		UserId: uuid.New(),
		Name:   "test phone book name",
	}

	err := repo.DB.Create(&phoneBook).Error
	assert.NoError(t, err)

	contact := entities.Contact{
		UserID:      phoneBook.UserId,
		PhoneBookID: phoneBook.ID,
		FirstName:   "Amir",
		LastName:    "Dashti",
		Username:    random.String(15),
		PhoneNumber: "19123456789",
	}

	err = repo.DB.Create(&contact).Error
	assert.NoError(t, err)

	return repo, &phoneBook, &contact
}

func cleanupContactInPhoneBookTestEnvironment(t *testing.T, repo *repositories.PGRepository, phoneBook *entities.PhoneBook, contact *entities.Contact) {
	err := repo.DB.Unscoped().Delete(phoneBook).Error
	assert.NoError(t, err)
	err = repo.DB.Unscoped().Delete(contact).Error
	assert.NoError(t, err)
}

func TestCreateContact(t *testing.T) {
	t.Run("withValidData", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		newContact := entities.Contact{
			UserID:      contact.UserID,
			PhoneBookID: phoneBook.ID,
			FirstName:   "test name",
			LastName:    "test family",
			Username:    random.String(15),
			PhoneNumber: "19120000000",
		}

		createdContact, err := repo.CreateContact(newContact)

		assert.NoError(t, err)
		assert.Equal(t, contact.UserID, createdContact.UserID)
		assert.Equal(t, contact.PhoneBookID, createdContact.PhoneBookID)
		assert.Equal(t, newContact.FirstName, createdContact.FirstName)
		assert.Equal(t, newContact.LastName, createdContact.LastName)
		assert.Equal(t, newContact.Username, createdContact.Username)
		assert.Equal(t, newContact.PhoneNumber, createdContact.PhoneNumber)

		err = repo.DB.Unscoped().Delete(createdContact).Error
		assert.NoError(t, err)
	})

	t.Run("withInvalidPhoneBook", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		newContact := entities.Contact{
			UserID:      contact.UserID,
			PhoneBookID: uuid.New(),
			FirstName:   "test",
			LastName:    "test",
			PhoneNumber: "19120000000",
		}

		_, err := repo.CreateContact(newContact)
		assert.Error(t, err)
	})

	t.Run("withExistsUsername", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		newContact := entities.Contact{
			UserID:      contact.UserID,
			PhoneBookID: phoneBook.ID,
			FirstName:   "test",
			LastName:    "test",
			Username:    contact.Username,
			PhoneNumber: "19120000000",
		}

		_, err := repo.CreateContact(newContact)
		assert.Error(t, err)
	})
}

func TestGetContactList(t *testing.T) {
	repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
	defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

	Contactlist, err := repo.GetContactList(*contact)

	assert.NoError(t, err)
	assert.NotEmpty(t, Contactlist)
	assert.Len(t, Contactlist, 1)
	assert.Equal(t, contact.Username, Contactlist[0].Username)
	assert.Equal(t, contact.PhoneBookID, Contactlist[0].PhoneBookID)
}

func TestGetContactById(t *testing.T) {
	t.Run("withValidData", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		result, err := repo.GetContactById(*contact)
		assert.NoError(t, err)
		assert.Equal(t, contact.UserID, result.UserID)
		assert.Equal(t, contact.PhoneBookID, result.PhoneBookID)
		assert.Equal(t, contact.FirstName, result.FirstName)
		assert.Equal(t, contact.LastName, result.LastName)
		assert.Equal(t, contact.Username, result.Username)
		assert.Equal(t, contact.PhoneNumber, result.PhoneNumber)
	})

	t.Run("withInvalidUserId", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		contactWithInvalidPhoneBookId := entities.Contact{
			UserID:      uuid.New(),
			PhoneBookID: contact.PhoneBookID,
		}
		contactWithInvalidPhoneBookId.ID = contact.ID
		_, err := repo.GetContactById(contactWithInvalidPhoneBookId)
		assert.Error(t, err)
	})

	t.Run("withInvalidContactId", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		contactWithInvalidPhoneBookId := entities.Contact{
			UserID:      contact.UserID,
			PhoneBookID: uuid.New(),
		}
		contactWithInvalidPhoneBookId.ID = contact.ID
		_, err := repo.GetContactById(contactWithInvalidPhoneBookId)
		assert.Error(t, err)
	})

	t.Run("withInvalidContactId", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		contactWithInvalidPhoneBookId := entities.Contact{
			UserID:      contact.UserID,
			PhoneBookID: contact.PhoneBookID,
		}
		contactWithInvalidPhoneBookId.ID = uuid.New()
		_, err := repo.GetContactById(contactWithInvalidPhoneBookId)
		assert.Error(t, err)
	})
}

func TestUpdateContactById(t *testing.T) {

	t.Run("withValidData", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		editedContact := entities.Contact{
			UserID:      contact.UserID,
			PhoneBookID: contact.PhoneBookID,
			Username:    random.String(15),
			FirstName:   "Amir Hossein",
			LastName:    "-T",
			PhoneNumber: "19100000000",
		}
		editedContact.ID = contact.ID

		updatedContact, err := repo.UpdateContactById(editedContact)
		assert.NoError(t, err)
		assert.Equal(t, editedContact.UserID, updatedContact.UserID)
		assert.Equal(t, editedContact.Username, updatedContact.Username)
		assert.Equal(t, editedContact.FirstName, updatedContact.FirstName)
		assert.Equal(t, editedContact.LastName, updatedContact.LastName)
		assert.Equal(t, editedContact.PhoneNumber, updatedContact.PhoneNumber)
	})

	t.Run("withSomeData", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		editedContact := entities.Contact{
			UserID:      contact.UserID,
			PhoneBookID: contact.PhoneBookID,
			FirstName:   "Amir Hossein",
		}
		editedContact.ID = contact.ID

		updatedContact, err := repo.UpdateContactById(editedContact)
		assert.NoError(t, err)
		assert.Equal(t, editedContact.FirstName, updatedContact.FirstName)
	})

	t.Run("withInvalidUserId", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		contactWithInvalidUserId := entities.Contact{
			UserID:      uuid.New(),
			PhoneBookID: contact.PhoneBookID,
			Username:    random.String(15),
			FirstName:   "Amir Hossein",
			LastName:    "-T",
			PhoneNumber: "19100000000",
		}
		contactWithInvalidUserId.ID = contact.ID

		_, err := repo.UpdateContactById(contactWithInvalidUserId)
		assert.Error(t, err)
	})

	t.Run("withInvalidPhoneBookId", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		contactWithInvalidPhoneBookId := entities.Contact{
			UserID:      contact.UserID,
			PhoneBookID: uuid.New(),
			Username:    random.String(15),
			FirstName:   "Amir Hossein",
			LastName:    "-T",
			PhoneNumber: "19100000000",
		}
		contactWithInvalidPhoneBookId.ID = contact.ID

		_, err := repo.UpdateContactById(contactWithInvalidPhoneBookId)
		assert.Error(t, err)
	})

	t.Run("withInvalidContactId", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		contactWithInvalidId := entities.Contact{
			UserID:      contact.UserID,
			PhoneBookID: contact.PhoneBookID,
			Username:    random.String(15),
			FirstName:   "Amir Hossein",
			LastName:    "-T",
			PhoneNumber: "19100000000",
		}
		contactWithInvalidId.ID = uuid.New()

		_, err := repo.UpdateContactById(contactWithInvalidId)
		assert.Error(t, err)
	})

	t.Run("withUsernameAlreadyExists", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		newContact := entities.Contact{
			UserID:      contact.UserID,
			PhoneBookID: contact.PhoneBookID,
			Username:    random.String(15),
			FirstName:   "alireza",
			LastName:    "rezaei",
			PhoneNumber: "19100000001",
		}

		err := repo.DB.Create(&newContact).Error
		assert.NoError(t, err)

		contactwithUsernameAlreadyExists := entities.Contact{
			UserID:   contact.UserID,
			Username: newContact.Username,
		}

		_, err = repo.UpdateContactById(contactwithUsernameAlreadyExists)
		assert.Error(t, err)

		err = repo.DB.Unscoped().Delete(newContact).Error
		assert.NoError(t, err)
	})
}

func TestDeleteContactById(t *testing.T) {
	t.Run("withValidData", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		err := repo.DeleteContactById(*contact)
		assert.NoError(t, err)

		var deletedContact entities.Contact
		err = repo.DB.Where("user_id = ? AND id = ? AND phone_book_id = ?", contact.UserID, contact.ID, contact.PhoneBookID).First(&deletedContact).Error
		assert.Error(t, err)
	})

	t.Run("withInvalidUserId", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		contactwithInvalidUserId := entities.Contact{
			UserID:      uuid.New(),
			PhoneBookID: contact.PhoneBookID,
		}
		contactwithInvalidUserId.ID = contact.ID

		err := repo.DeleteContactById(contactwithInvalidUserId)
		assert.NoError(t, err)

		var deletedContact entities.Contact
		err = repo.DB.Where("user_id = ? AND id = ? AND phone_book_id = ?", contact.UserID, contact.ID, contact.PhoneBookID).First(&deletedContact).Error
		assert.NoError(t, err)
		assert.Equal(t, contact.Username, deletedContact.Username)
	})

	t.Run("withInvalidPhoneBookId", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		contactwithInvalidPhoneBookId := entities.Contact{
			UserID:      contact.UserID,
			PhoneBookID: uuid.New(),
		}
		contactwithInvalidPhoneBookId.ID = contact.ID

		err := repo.DeleteContactById(contactwithInvalidPhoneBookId)
		assert.NoError(t, err)

		var deletedContact entities.Contact
		err = repo.DB.Where("user_id = ? AND id = ? AND phone_book_id = ?", contact.UserID, contact.ID, contact.PhoneBookID).First(&deletedContact).Error
		assert.NoError(t, err)
		assert.Equal(t, contact.Username, deletedContact.Username)
	})

	t.Run("withInvalidContactId", func(t *testing.T) {
		repo, phoneBook, contact := setupContactInPhoneBookTestEnvironment(t)
		defer cleanupContactInPhoneBookTestEnvironment(t, repo, phoneBook, contact)

		contactwithInvalidId := entities.Contact{
			UserID:      contact.UserID,
			PhoneBookID: contact.PhoneBookID,
		}
		contactwithInvalidId.ID = uuid.New()

		err := repo.DeleteContactById(contactwithInvalidId)
		assert.NoError(t, err)

		var deletedContact entities.Contact
		err = repo.DB.Where("user_id = ? AND id = ? AND phone_book_id = ?", contact.UserID, contact.ID, contact.PhoneBookID).First(&deletedContact).Error
		assert.NoError(t, err)
		assert.Equal(t, contact.Username, deletedContact.Username)
	})
}