package service

import (
	"fmt"
	"regexp"
	"strings"

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

func (svc *TemplateService) CreateTemplate(temp entities.Template) error {
	err := svc.db.AddTemplate(temp)
	if err != nil {
		panic(err)
	}
	return nil
}
func (svc *TemplateService) GetTemplateMapContent(tempName string) (string, map[string]string, error) {
	temp, err := svc.db.GetTemplate(tempName)
	MapTemp := make(map[string]string)
	if err != nil {
		return "", make(map[string]string), err
	}
	r, _ := regexp.Compile(`{([a-z\_\-]+)}`)
	res := r.FindStringSubmatch(temp.Content)
	for k, v := range res {
		if k%2 != 0 {
			_, ok := MapTemp[v]
			if !ok {
				MapTemp[v] = ""
			}
		}
	}
	return temp.Content, MapTemp, nil
}

func (svc *TemplateService) GenerateTemplate(content string, tempMap map[string]string) (string, error) {
	for k, v := range tempMap {
		content = strings.Replace(content, fmt.Sprint("{", k, "}"), v, -1)
	}
	return content, nil
}

func (svc *TemplateService) GetAllTemplates() ([]entities.Template, error) {
	templates, err := svc.db.AllTemplates()
	if err != nil {
		return templates, err
	}
	return templates, err
}
