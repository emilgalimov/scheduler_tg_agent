package app

import (
	"context"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/model"
)

type Repository interface {
	CreateUser(context.Context, model.User) error
	GetUserByChatId(ctx context.Context, chatID int64) (model.User, error)
}
