package repositories

import (
	"database/sql"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/nurtai325/kaspi/mailing/internal/db"
	"github.com/nurtai325/kaspi/mailing/internal/models"
)

type CachedClients struct {
	mu      sync.Mutex
	clients []models.Client
}

func (c *CachedClients) Set(freshClients []models.Client) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.clients = freshClients
}

var cachedClients CachedClients
var clientsCacheFresh atomic.Bool

func NewClient() ClientRepository {
	conn := db.New()
	return &clientRepository{
		conn: conn,
	}
}

type clientRepository struct {
	conn *sql.DB
}

func (c *clientRepository) Get() ([]models.Client, error) {
	if clientsCacheFresh.Load() {
		cachedClients.mu.Lock()
		defer cachedClients.mu.Unlock()
		clientsLen := len(cachedClients.clients)
		clients := make([]models.Client, clientsLen, clientsLen)
		copy(clients, cachedClients.clients)
		return clients, nil
	}

	r, err := c.conn.Query("SELECT id, name, phone, expires, connected, token FROM clients ORDER BY id ASC;")
	if err != nil {
		return nil, err
	}

	clients := []models.Client{}
	for r.Next() {
		var client models.Client
		err = r.Scan(&client.Id, &client.Name, &client.Phone, &client.Expires, &client.Connected, &client.Token)
		if err != nil {
			return nil, err
		}
		client.Expires.Time = client.Expires.Time.Add(time.Hour * 5)
		clients = append(clients, client)
	}

	cachedClients.Set(clients)
	clientsCacheFresh.Store(true)

	return clients, nil
}

func (c *clientRepository) Insert(client models.Client) error {
	_, err := c.conn.Exec(
		"INSERT INTO clients(name, token, phone, expiration_notified, expires, connected) VALUES($1, $2, $3, $4, $5, $6);",
		client.Name, client.Token, client.Phone, false, time.Now().UTC(), client.Connected,
	)
	if err != nil {
		return err
	}

	clientsCacheFresh.Store(false)
	return nil
}

func (c *clientRepository) Extend(id int, duration int, unit string) error {
	r := c.conn.QueryRow("SELECT expires FROM clients WHERE id = $1 LIMIT 1;", id)
	var expireDate time.Time
	err := r.Scan(&expireDate)
	if err != nil {
		return err
	}

	days := 0
	months := 0

	if unit == "days" {
		days = duration
	} else {
		months = duration
	}

	if expireDate.After(time.Now()) {
		expireDate = expireDate.AddDate(0, months, days)
	} else {
		expireDate = time.Now().UTC().AddDate(0, months, days)
	}

	_, err = c.conn.Exec("UPDATE clients SET expires = $1 WHERE id = $2", expireDate, id)
	if err != nil {
		return err
	}
	clientsCacheFresh.Store(false)

	return nil
}

func (c *clientRepository) ConnectWh(id int) error {
	_, err := c.conn.Exec("UPDATE clients SET connected = $1 WHERE id = $2", true, id)
	if err != nil {
		return err
	}

	clientsCacheFresh.Store(false)
	return nil
}

func (c *clientRepository) Deactivate(id int) error {
	_, err := c.conn.Exec("UPDATE clients SET expires = $1 WHERE id = $2", time.Now().UTC(), id)
	if err != nil {
		return err
	}
	clientsCacheFresh.Store(false)

	return nil
}
