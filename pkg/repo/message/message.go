package message

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgconn"
	"log"
	"messagio/pkg/db"
	"messagio/pkg/domain/message"
	interfaces "messagio/pkg/repo/message/interface"
)

type messageDataBase struct {
	db *sql.DB
}

func NewMessageDataBase(db *sql.DB) interfaces.MessageRepository {
	return &messageDataBase{
		db: db,
	}
}

func (r *messageDataBase) Migrate(ctx context.Context) error {
	accQuery := `
    CREATE TABLE IF NOT EXISTS message (
		id SERIAL PRIMARY KEY,
		name TEXT ,
		value TEXT,
		status TEXT
	);
    `
	_, err := r.db.ExecContext(ctx, accQuery)
	if err != nil {
		message := db.ErrMigrate.Error() + " message"
		log.Printf("%q: %s\n", message, err.Error())
		return db.ErrMigrate
	}

	return err
}

func (r *messageDataBase) Post(ctx context.Context, newPeople message.Message) (int64, error) {
	var id int64

	err := r.db.QueryRowContext(ctx, "INSERT INTO message(name, value, status) values($1, $2, $3) RETURNING id",
		newPeople.Name, newPeople.Value, newPeople.Status).Scan(&id)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return 0, db.ErrDuplicate
			}
		}
		return 0, err
	}

	return id, nil
}

func (r *messageDataBase) PUT(ctx context.Context, id int64, updatedMessage message.Message) error {
	res, err := r.db.ExecContext(ctx, "UPDATE message SET name = $1, value = $2, status = $3 WHERE id = $4",
		updatedMessage.Name, updatedMessage.Value, updatedMessage.Status, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return db.ErrNotExist
	}

	return nil
}

func (r *messageDataBase) GetAll(ctx context.Context) ([]message.Message, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM message")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []message.Message

	for rows.Next() {
		var msg message.Message
		if err := rows.Scan(&msg.ID, &msg.Name, &msg.Value, &msg.Status); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
