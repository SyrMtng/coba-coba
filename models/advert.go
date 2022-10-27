package models

import (
	"gorm.io/gorm"
)

type Advert struct {
	gorm.Model
	Id           int    `form:"id" json:"id" validate:"required"`
	Advert       string `form:"advert" json:"advert" validate:"required"`
	View         int    `form:"view" json:"view" validate:"required"`
	UserIdAdvert uint   `gorm:"foreignKey:UserIdUserIdAdvert"`
	IdMIklan     uint
}
