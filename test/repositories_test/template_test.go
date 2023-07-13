package repositories_test

import (
	"testing"

	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setupTestEnvironmetTemplate(t *testing.T) (*repositories.PGRepository, *entities.Template) {
	repo := repositories.NewGormDatabase()
	temp := entities.Template{
		ID:       uuid.New(),
		TempName: "TestTemp",
		Content:  "Test Create Template",
	}
	err := repo.AddTemplate(temp)
	assert.NoError(t, err)

	var tempSample entities.Template

	err = repo.DB.Where("id=?", temp.ID).First(&tempSample).Error

	assert.NoError(t, err)
	tempSample.ID = temp.ID

	return repo, &temp

}

func cleanupTestEnviromentTemplate(t *testing.T, repo *repositories.PGRepository, temp *entities.Template) {
	err := repo.DB.Unscoped().Delete(temp).Error
	assert.NoError(t, err)
}

func TestAddtemplate(t *testing.T) {
	repo, temp := setupTestEnvironmetTemplate(t)
	defer cleanupTestEnviromentTemplate(t, repo, temp)
	var tempSample entities.Template
	err := repo.DB.Where("id=?", temp.ID).First(&tempSample).Error
	assert.NoError(t, err)
	assert.Equal(t, temp.TempName, tempSample.TempName)
}

func TestGetTemplate(t *testing.T) {
	repo, temp := setupTestEnvironmetTemplate(t)
	defer cleanupTestEnviromentTemplate(t, repo, temp)
	tempSample, err := repo.GetTemplate(temp.TempName)
	assert.Nil(t, err)
	assert.NotEmpty(t, tempSample.TempName)
	assert.NotEmpty(t, tempSample.ID)
	assert.NotEmpty(t, tempSample.Content)
}

func TestGetTemplateInvalidName(t *testing.T) {
	repo, temp := setupTestEnvironmetTemplate(t)
	defer cleanupTestEnviromentTemplate(t, repo, temp)
	invalidTempName := ""
	var invalidTemp entities.Template
	invalidTemp.TempName = invalidTempName
	tempSample, err := repo.GetTemplate(invalidTempName)
	assert.NotNil(t, err)
	assert.Empty(t, tempSample)
}
