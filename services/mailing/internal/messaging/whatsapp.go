package messaging

import (
	"context"

	"github.com/nurtai325/kaspi/mailing/internal/models"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

type whatsapp struct {
}

func (w *whatsapp) Message(message models.Message) error {
	client, ok := clients.Get(message.Sender)
	if !ok {
		return ErrNoClient
	}

	_, err := client.SendMessage(context.Background(), types.NewJID(
		message.Receiver,
		"s.whatsapp.net"),
		&waE2E.Message{
			Conversation: proto.String(message.Text),
		},
	)

	return err
}
