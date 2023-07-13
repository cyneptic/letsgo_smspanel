package repositories_test

import (
	"testing"

	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/stretchr/testify/assert"
)

func TestSinglePriceChange(t *testing.T) {
	db := repositories.NewGormDatabase()

	err := db.EditSingleMessagePrice(25)
	assert.NoError(t, err)

	var price int
	q := db.DB.Model(&entities.Prices{}).First(&price)
	assert.NoError(t, q.Error)

	assert.Equal(t, 25, price)
}
