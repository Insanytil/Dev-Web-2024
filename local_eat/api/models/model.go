package model

import (
	"time"
)

type Users struct {
	Username         *string   `json:"username,omitempty" example:"john_vleminckx" gorm:"primaryKey"`
	Password         string    `json:"password" example:"ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"`
	Email            *string   `json:"email,omitempty" example:"mateo@example.com"`
	CreatedAt        time.Time `json:"created" example:"Mon Jan 2 15:04:05 MST 2006"`
	ProfilePictureId *string   `json:"profile_picture_id,omitempty" example:"1524689"`
}

type Producers struct {
	ID        int     `json:"id" example:"1" gorm:"primaryKey; autoIncrement"`
	Username  *string `json:"username,omitempty" example:"john_vleminckx"`
	Users     Users   `json:"users" gorm:"foreignKey:Username; references:Username"`
	Firstname string  `json:"firstname" example:"John"`
	Lastname  string  `json:"lastname" example:"Vleminckx"`
	PhoneNum  string  `json:"phone_num" example:"0483598799"`
	EmailPro  string  `json:"email" example:"postmaster@example.com"`
}
