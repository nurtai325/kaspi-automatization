package models

type Order struct {
	Code      string `json:"code"`
	Completed bool   `json:"completed"`
}
