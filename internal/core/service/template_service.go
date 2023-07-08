package service

import (
	repositories "github.com/cyneptic/letsgo-smspanel/infrastructure/repository"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/cyneptic/letsgo-smspanel/internal/core/ports"
)

type TemplateService struct {
	db ports.TemplateRepositoryContract
}

func NewTemplateService() *TemplateService {
	db := repositories.NewGormDatabase()
	return &TemplateService{
		db: db,
	}
}

func (svc *TemplateService) CreateTemplate(temp entities.Template) error        { return nil }
func (svc *TemplateService) CreateTemplateContent(temp entities.Template) error { return nil }
