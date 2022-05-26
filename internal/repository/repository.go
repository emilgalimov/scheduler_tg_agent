package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/model"
)

type Repository struct {
	pool *pgxpool.Pool
}

func (r *Repository) GetActionByChatID(ctx context.Context, chatID int64) (action model.ActiveLiveAction, err error) {
	//language=PostgreSQL
	const sql = "SELECT chat_id, name, state, data FROM active_live_actions WHERE chat_id = $1"

	err = r.pool.QueryRow(ctx, sql, chatID).Scan(
		&action.ChatID,
		&action.Name,
		&action.State,
		&action.Data,
	)

	return
}

func (r *Repository) CreateOrUpdateAction(ctx context.Context, action model.ActiveLiveAction) {
	//language=PostgreSQL
	const sql = `INSERT INTO active_live_actions (chat_id, name, state, data) 
					VALUES ($1, $2, $3, $4) ON CONFLICT (chat_id)
					DO UPDATE SET name = $2, state = $3, data = $4`

	r.pool.QueryRow(ctx, sql, action.ChatID, action.Name, action.State, action.Data)
}

func (r *Repository) DeleteActionByChatID(ctx context.Context, chatID int64) {
	//language=PostgreSQL
	const sql = `DELETE FROM active_live_actions WHERE chat_id = $1`

	r.pool.QueryRow(ctx, sql, chatID)
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
