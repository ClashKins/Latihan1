package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID  uint
	PhotoID uint   `json:"photo_id" form:"photo_id" valid:"required~Message on a comment is required"`
	Message string `json:"message" form:"message" valid:"required~Message on a comment is required"`
	User    *User
	Photo   *Photo
}

func (u *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreated := govalidator.ValidateStruct(u)

	if errCreated != nil {
		err = errCreated
		return
	}

	err = nil
	return
}