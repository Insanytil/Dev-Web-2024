package model

import ()

type Producers struct {
	Id      int    `json:"id" example:"1" gorm:"primaryKey; autoIncrement"`
	Name    string `json:"name" example:"John"`
	Picture string `json:"picture" example:"John.jpg"`
	Created string `json:"created" example:"2020-01-01"`
}

type Users struct {
	Username *string `json:"username,omitempty" example:"John" gorm:"primaryKey"`
	Password string  `json:"password" example:"1234"`
	Email    *string `json:"email,omitempty" example:"john@example.com"`
}