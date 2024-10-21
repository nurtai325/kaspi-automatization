package models

import (
	"database/sql"
)

type Client struct {
	Id                 uint         `json:"id"`
	Name               string       `json:"name"`
	Token              string       `json:"token"`
	Phone              string       `json:"phone"`
	ExpirationNotified bool         `json:"expiration_notified"`
	Expires            sql.NullTime `json:"expires"`
	Connected          bool         `json:"connected"`
}
