package messaging

type Messenger interface {
	Message(phone, text string) error
}

func New() Messenger {
	return &whatsapp{}
}
