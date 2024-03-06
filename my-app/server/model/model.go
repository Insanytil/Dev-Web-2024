package model

type Technology struct {
	Name    string `json:"name"`
	Details string `json:"details"`
}

type Producers struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Created string `json:"created"`
}
