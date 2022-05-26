package router

import (
	"context"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/model"
)

type Repository interface {
	GetActionByChatID(ctx context.Context, chatID int64) (model.ActiveLiveAction, error)
	CreateOrUpdateAction(context.Context, model.ActiveLiveAction)
	DeleteActionByChatID(ctx context.Context, chatID int64)
}
