package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/model"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) GetUserByChatId(ctx context.Context, chatID int64) (user model.User, err error) {

	//language=PostgreSQL
	const sql = `SELECT id, chat_id FROM users WHERE chat_id = $1`

	err = r.pool.QueryRow(ctx, sql, chatID).Scan(
		&user.ID,
		&user.ChatID,
	)

	return

}

func (r *Repository) CreateUser(ctx context.Context, user model.User) error {
	//language=PostgreSQL
	const sql = `INSERT INTO users(id, chat_id) VALUES ($1, $2)`

	r.pool.QueryRow(ctx, sql, user.ID, user.ChatID)
	return nil
}
