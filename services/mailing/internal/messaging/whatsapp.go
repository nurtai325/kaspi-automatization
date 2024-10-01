package messaging

type whatsapp struct {
}

func (w *whatsapp) Message(phone, text string) error {
	return nil
}
