package models

type Order struct {
	Id        string `json:"id"`
	Completed bool   `json:"completed"`
}
