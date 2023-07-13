package repositories_test

import (
	"fmt"
	"testing"

	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// ! Setup Request and Cleanup ContactList
func setupTestEnvironmetContactList(t *testing.T) (*repositories.PGRepository, *entities.PhoneBook, *entities.Contact) {
	repo := repositories.NewGormDatabase()
	phoneBook := entities.PhoneBook{
		DBModel: entities.DBModel{
			ID: uuid.New(),
		},
		UserId: uuid.New(),
		Name:   "test",
	}

	//todo create phoneBookList
	_, err := repo.CreatePhoneBookList(phoneBook)
	assert.NoError(t, err)

	var phoneBookSample entities.PhoneBook

	err = repo.DB.Where("user_id=?", phoneBook.UserId).First(&phoneBookSample).Error
	assert.NoError(t, err)
	phoneBookSample.UserId = phoneBook.UserId

	//todo CreateContact
	contactlist1 := entities.Contact{
		UserID:      uuid.New(),
		PhoneBookID: phoneBookSample.ID,
	}
	_, err = repo.CreateContact(contactlist1)

	assert.Nil(t, err)
	return repo, &phoneBook, &contactlist1

}

func cleanupTestEnviromentContactList(t *testing.T, repo *repositories.PGRepository, phoneBook *entities.PhoneBook, contact *entities.Contact) {
	err := repo.DB.Unscoped().Delete(phoneBook).Error
	assert.NoError(t, err)
	err = repo.DB.Unscoped().Where("user_id", contact.UserID).Delete(contact).Error
	assert.NoError(t, err)

}

// ! Setup Request and Cleanup User
func setupTestEnvironmetUser(t *testing.T) (*repositories.PGRepository, *entities.User) {
	repo := repositories.NewGormDatabase()
	user := entities.User{
		Name:        "test",
		PhoneNumber: "09121232231",
	}

	user, err := repo.DB.Create(&user)
	assert.NoError(t, err)

	var userSample entities.User
	fmt.Println("idsample", user.ID)
	err = repo.DB.Where("id=?", user.ID).First(&userSample).Error
	assert.NoError(t, err)
	userSample.ID = user.ID

	assert.Nil(t, err)
	return repo, &user

}

func cleanupTestEnviromentUser(t *testing.T, repo *repositories.PGRepository, user *entities.User) {
	err := repo.DB.Unscoped().Delete(user).Error
	assert.NoError(t, err)
}

// ! Setup Request and Cleanup Number

// !tests
func TestRequestContactList(t *testing.T) {
	repo, pbook, contact := setupTestEnvironmetContactList(t)
	defer cleanupTestEnviromentContactList(t, repo, pbook, contact)
	var phoneBookSample entities.PhoneBook
	err := repo.DB.Where("user_id=?", pbook.UserId).First(&phoneBookSample).Error
	assert.Nil(t, err)
	contactlist, err := repo.RequestContactList(phoneBookSample.UserId)
	assert.Nil(t, err)
	assert.NotNil(t, contactlist)
	assert.Len(t, contactlist, 1)
}

func TestRequestUser(t *testing.T) {
	repo, user := setupTestEnvironmetUser(t)
	defer cleanupTestEnviromentUser(t, repo, user)
	var userSample entities.User
	err := repo.DB.Where("id=?", user.ID).First(&userSample).Error
	assert.Nil(t, err)
	assert.NotEmpty(t, userSample)
}
