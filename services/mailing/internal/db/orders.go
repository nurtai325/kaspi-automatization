package db

import "database/sql"

type Order struct {
}

func (o *Order) Migrate(conn *sql.DB) error {
    _, err := conn.Exec(`
        CREATE TABLE IF NOT EXISTS orders (
        id VARCHAR(9) PRIMARY KEY, 
        completed BOOLEAN NOT NULL,
        sum BIGINT NOT NULL,
        phone VARCHAR(11) NOT NULL,
        product_code VARCHAR(9) NOT NULL,
        customer VARCHAR(70) NOT NULL,
        entries TEXT NOT NULL
        );
        `)
    return err
}

func (o Order) Name() string {
    return "create_orders_table"
}
