package messaging

import (
	"github.com/nurtai325/kaspi/mailing/internal/models"
)

type Messenger interface {
	Message(message models.Message) error
}

func New(clientId uint) Messenger {
	return &whatsapp{}
}
