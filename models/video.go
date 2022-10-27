package models

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Id          int    `form:"id" json:"id" validate:"required"`
	Video       string `form:"video" json:"video" validate:"required"`
	Tumb       	string `form:"tumb" json:"tumb" validate:"required"`
	Title       string `form:"title" json:"title" validate:"required"`
	Description string `form:"description" json:"description" validate:"required"`
	DisableAds  bool   `gorm:"default:false"`
	// UserIdVideo uint   `gorm:"foreignKey:UserIdUserIdVideo"`
}

// CRUD
// Untuk Menampilkan Semua Data Video
func ReadVideo(db *gorm.DB, videos *[]Video)(err error) {
	err = db.Find(videos).Error
	if err != nil {
		return err
	}
	return nil
}
// Untuk Membuat atau menambahkan Data Video
func CreateVideo(db *gorm.DB, newVideo *Video) (err error) {
	err = db.Create(newVideo).Error
	if err != nil {
		return err
	}
	return nil
}