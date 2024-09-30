package db

type Order struct {
}

func (o Order) Migrate() error {
    conn := GetDBConnection()
    _, err := conn.Exec(`
        CREATE TABLE IF NOT EXISTS orders (
            id VARCHAR(9) PRIMARY KEY, 
            completed BOOLEAN NOT NULL
        );
        `)
    return err
}

func (o Order) Name() string {
    return "create_orders_table"
}
