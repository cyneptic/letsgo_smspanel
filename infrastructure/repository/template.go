package repositories

import (
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
)

func (pc *PGRepository) AddTemplate(temp entities.Template) error {
	res := pc.DB.Create(&temp)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (pc *PGRepository) GetTemplate(tempname string) (entities.Template, error) {
	var temp entities.Template
	res := pc.DB.Where("tempname=?", tempname).First(&temp)
	if res.Error != nil {
		return entities.Template{}, res.Error
	}
	return temp, nil

}

func (pc *PGRepository) AllTemplates() ([]entities.Template, error) {
	var templates []entities.Template
	res := pc.DB.Find(&templates)
	if res.Error != nil {
		return templates, res.Error
	}
	return templates, nil
}
