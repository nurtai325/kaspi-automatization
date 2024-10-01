package db

type Client struct {
}

func (c Client) Migrate() error {
    conn := GetDBConnection()
    _, err := conn.Exec(`
        CREATE TABLE IF NOT EXISTS clients (
            id SERIAL PRIMARY KEY, 
            name VARCHAR(50) NOT NULL,
            token VARCHAR(50) NOT NULL,
            phone VARCHAR(15) NOT NULL
        );
        `)
    return err
}

func (c Client) Name() string {
    return "create_clients_table"
}
