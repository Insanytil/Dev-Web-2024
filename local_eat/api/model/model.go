package model

type Producers struct {
	Id      int    `json:"id" example:"1"`
	Name    string `json:"name" example:"John"`
	Picture string `json:"picture" example:"John.jpg"`
	Created string `json:"created" example:"2020-01-01"`
}

type Users struct {
	Username string `json:"username" example:"John"`
	Password string `json:"password" example:"1234"`
	Email	string `json:"email" example:"John@example.com"`
	Age	  int    `json:"age" example:"20"`
	Gender	  string	`json:"gender" example:"M"`
	Address	  string	`json:"address" example:"1234 Main St"`
	Locality  int	`json:"locality" example:"1650"`
	Cellphone string `json:"cellphone" example:"1234567890"`
}