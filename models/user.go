package models

import (
	"LATIHAN1/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type levelmodel string

const (
	SuperAdmin levelmodel = "superadmin"
	Admin      levelmodel = "admin"
	Users      levelmodel = "user"
)

type User struct {
	GormModel
	Username string     `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email    string     `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required"`
	Password string     `gorm:"not null" json:"password" form:"password" valid:"required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Photos   []Photo    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Age      uint       `gorm:"not null;" json:"age" form:"age" valid:"required~Your age is required"`
	Comments []Comment  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	Level    levelmodel `sql:"type:levelmodel"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil 
	return
}