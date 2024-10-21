package db

import "database/sql"

type Customer struct {
}

func (c *Customer) Migrate(conn *sql.DB) error {
    _, err := conn.Exec(`
        CREATE TABLE IF NOT EXISTS customers (
        id SERIAL PRIMARY KEY, 
        name VARCHAR(50) NOT NULL,
        phone VARCHAR(15) NOT NULL
        );
        `)
    return err
}

func (c Customer) Name() string {
    return "create_customers_table"
}
