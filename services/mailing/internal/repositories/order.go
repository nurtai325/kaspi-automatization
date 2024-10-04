package repositories

import (
	"database/sql"
	"errors"

	"github.com/nurtai325/kaspi/mailing/internal/db"
	"github.com/nurtai325/kaspi/mailing/internal/models"
)

var (
	ErrRecordIsPresent = errors.New("record is already present in the database")
	ErrCompletingOrder = errors.New("failed to complete an order")
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
        `INSERT INTO orders(id, completed, sum, phone, product_code, customer, entries) 
         VALUES($1, $2, $3, $4, $5, $6, $7)
         ON CONFLICT(id) DO NOTHING;`,
        order.Id,
        order.Completed,
        order.Sum,
        order.Phone,
        order.ProductCode,
        order.Customer,
        order.Entries,
        )

    return err
}

func (o *orderRepository) Complete(id string) error {
    result, err := o.conn.Exec(
        "UPDATE orders SET completed = true WHERE id = $1",
        id,
        )
    if err != nil {
        return err
    }

    r, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if r != 1 {
        return ErrCompletingOrder
    }

    return nil
}
