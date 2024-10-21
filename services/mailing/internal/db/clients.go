package db

import "database/sql"

type Client struct {
}

func (c *Client) Migrate(conn *sql.DB) error {
    _, err := conn.Exec(`
        CREATE TABLE IF NOT EXISTS clients (
        id SERIAL PRIMARY KEY, 
        name VARCHAR(50) NOT NULL,
        token VARCHAR(50) NOT NULL,
        phone VARCHAR(15) NOT NULL,
        expiration_notified BOOLEAN NOT NULL,
        connected BOOLEAN NOT NULL,
        expires TIMESTAMP WITHOUT TIME ZONE
        );
        `)
    return err
}

func (c Client) Name() string {
    return "create_clients_table"
}
