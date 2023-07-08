package service_test

import (
	"testing"

	"github.com/cyneptic/letsgo-smspanel/internal/core/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestEditSinglePriceService(t *testing.T) {
	svc := service.NewAdminService()

	err := svc.EditSingleMessagePrice(uuid.New(), 25)
	assert.Error(t, err) // should error because account is not admin.
}
