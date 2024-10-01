package repositories

import (
	"database/sql"
	"errors"

	"github.com/nurtai325/kaspi/mailing/internal/db"
	"github.com/nurtai325/kaspi/mailing/internal/models"
)

var (
	ErrRecordIsPresent = errors.New("record is already present in the database")
)

func Order() OrderRepository {
	conn := db.GetDBConnection()
	return &orderRepository{
		conn: conn,
	}
}

type orderRepository struct {
	conn *sql.DB
}

func (o *orderRepository) Insert(order models.Order) error {
	_, err := o.conn.Exec(
		"INSERT INTO orders(id, completed) VALUES($1, $2)",
		order.Id,
		order.Completed,
	)

	return err
}
