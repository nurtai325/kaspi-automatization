package messaging

import (
	"errors"
	"sync"

	"go.mau.fi/whatsmeow"
)

var (
	clients     = ClientsMap{clients: make(map[string]*whatsmeow.Client)}
	ErrNoClient = errors.New("sender doesn't have a whatsapp client session")
)

type ClientsMap struct {
	mu      sync.Mutex
	clients map[string]*whatsmeow.Client
}

func Add(phone string, cl *whatsmeow.Client) {
	clients.mu.Lock()
	defer clients.mu.Unlock()
	clients.clients[phone] = cl
}

func (c *ClientsMap) Get(phone string) (*whatsmeow.Client, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	client, ok := c.clients[phone]
	return client, ok
}
