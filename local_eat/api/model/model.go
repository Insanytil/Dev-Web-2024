package model

type Producers struct {
	Id      int    `json:"id" example:"1"`
	Name    string `json:"name" example:"John"`
	Picture string `json:"picture" example:"John.jpg"`
	Created string `json:"created" example:"2020-01-01"`
}
