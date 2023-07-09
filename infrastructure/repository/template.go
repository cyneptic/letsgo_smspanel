package repositories

import (
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
)

func (pc *PGRepository) AddTemplate(temp entities.Template) error {
	t := entities.Template{ID: temp.ID,
		TempName: temp.TempName,
		Content:  temp.Content}
	res := pc.DB.Create(&t)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (pc *PGRepository) GetTemplate(tempname string) (entities.Template, error) {
	var temp entities.Template
	res := pc.DB.Where("tempname=?", tempname).First(&temp)
	if res.Error != nil {
		panic("Can't Found Template By Name")
	}
	return temp, nil

}
