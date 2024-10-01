package db

type User struct {
}

func (c User) Migrate() error {
    conn := GetDBConnection()
    _, err := conn.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY, 
            name VARCHAR(50) NOT NULL,
            phone VARCHAR(15) NOT NULL
        );
        `)
    return err
}

func (c User) Name() string {
    return "create_users_table"
}
