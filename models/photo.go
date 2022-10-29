package models

import (
	"time"

	"github.com/luthfikw/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Title     string `json:"title" form:"title" valid:"required~Title of your photo is required"`
	Caption   string `json:"caption" form:"caption"`
	PhotoUrl  string `json:"photo_url" form:"photo_url" valid:"required~Photo URL of your photo is required"`
	UserID    uint   `json:"user_id"`
	User      *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}