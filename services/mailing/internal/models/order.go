package models

type Order struct {
	Id          string `json:"id"`
	Completed   bool   `json:"completed"`
	Phone       string `json:"phone"`
	Customer    string `json:"customer"`
	Sum         int64  `json:"sum"`
	ProductCode string `json:"product_code"`
	Entries     string `json:"entries"`
}

type Entry struct {
	Id           string `json:"id"`
	Price        int64  `json:"price"`
	Quantity     int    `json:"quantity"`
	ProductCode  string `json:"product_code"`
	ProductName  string `json:"product_name"`
	DeliveryCost int64  `json:"delivery_cost"`
}
