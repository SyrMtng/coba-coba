package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int    `form:"id" json:"id" validate:"required"`
	Name     string `form:"name" json:"name" validate:"required"`
	Username string `form:"username" json:"username" validate:"required"`
	Image    string `form:"image" json:"image" validate:"required"`
	Email    string `form:"email" json:"email" validate:"required"`
	Role     string `form:"role" json:"role" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
	Disable  bool
	Videos   []*Video  `gorm:"foreignKey:UserIdVideo"`
	Adverts  []*Advert `gorm:"foreignKey:UserIdAdvert"`
}

type LoginForm struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

func CreateUser(db *gorm.DB, newUser *User) (err error) {
	err = db.Create(newUser).Error
	if err != nil {
		return err
	}
	return nil
}

func FindUserByUsername(db *gorm.DB, user *User, username string) (err error) {
	err = db.Where("username=?", username).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

func FindUserById(db *gorm.DB, user *User, id int) (err error) {
	err = db.Where("id=?", id).First(user).Error
	if err != nil {
		return err
	}
	return nil
}
