package models

import (
	"time"
)

type OwnModel struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-" `
	UpdatedAt time.Time  `json:"-" `
	DeletedAt *time.Time `json:"-"`
}

type User struct {
	OwnModel
	//Id       uint   `json:"user_id"`
	UserName string `json:"user_name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"` //`json:"-"` for first time maybe throw error so check on and change it to -> `json:"password"` then run again!

}
