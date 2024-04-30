package models

import (
	"time"
)

type Users struct {
	Username         *string   `json:"username,omitempty" example:"john_vleminckx" gorm:"primaryKey; varchar(20); unique"`
	Producer         *Producers `gorm:"foreignKey:Username; references:Username; constraint:OnDelete:CASCADE;"`
	Password         string    `json:"password" example:"ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"`
	Email            *string   `json:"email,omitempty" example:"mateo@example.com" gorm:"type:varchar(50)"`
	CreatedAt        time.Time `json:"createdAt" example:"Mon Jan 2 15:04:05 MST 2006"`
	ProfilePictureId *string   `json:"profile_picture_id,omitempty" example:"1524689"`
}

type Producers struct {
	ID        int    `json:"id" example:"1" gorm:"primaryKey; autoIncrement; type:int; not null"`
	Username  string `json:"username" example:"john_vleminckx" gorm:"not null; index"`
	Firstname string `json:"firstname" example:"John" gorm:"type:char(20); not null"`
	Lastname  string `json:"lastname" example:"Vleminckx" gorm:"type:char(20); not null"`
	PhoneNum  string `json:"phone_num" example:"0483598799" gorm:"type:char(10); not null"`
	EmailPro  string `json:"email" example:"postmaster@example.com" gorm:"type:varchar(50); not null"`
}

type Category struct {
	ID          string  `json:"id" example:"CAT1" gorm:"primaryKey;type:char(4);not null"`
	Name        string  `json:"name" example:"Legume" gorm:"type:varchar(30);not null;unique"`
	MotherCatID *string `json:"mother_cat,omitempty" example:"CAT2" gorm:"type:char(4);null"`
	MotherCat   *Category
	Description *string `json:"description,omitempty" example:"Ceci est une carrotte" gorm:"type:longtext"`
}

type Product struct {
	ID          string    `json:"id" example:"PROD1" gorm:"primaryKey;type:char(5);not null"`
	Name        string    `json:"name" example:"Laptop" gorm:"type:varchar(30);not null;unique"`
	CategoryID  string    `json:"cat" example:"CAT1" gorm:"type:char(4);not null"`
	Category    *Category `gorm:"foreignKey:CategoryID; references:ID"`
	Description *string   `json:"description,omitempty" example:"A powerful laptop with high-resolution display." gorm:"type:longtext;null"`
}

type Company struct {
	CompanyName string `gorm:"primaryKey"`
	Password    string `gorm:"not null"`
	Alias       string `gorm:"not null;unique"`
	Address     string `gorm:"not null"`
	Mail        string `gorm:"not null"`
	PhoneNum    string `gorm:"not null"`
	VATNum      string `gorm:"not null"`
	Description string `gorm:"type:longtext"`
}

type RelCompProd struct {
	ProducerID  string `gorm:"primaryKey;size:7"`
	CompanyName string `gorm:"primaryKey"`
}
