package models

import (
	"time"

	"affan/helpers"

	"github.com/luthfikw/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Username  string     `gorm:"not null" json:"username" valid:"required~username is required"`
	Email     string     `gorm:"not null" json:"email" valid:"required~email is required"`
	Password  string     `gorm:"not null" json:"-" valid:"required~password is required"`
	Photos    []Photo    `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPassword(u.Password)

	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPassword(u.Password)

	err = nil
	return
}