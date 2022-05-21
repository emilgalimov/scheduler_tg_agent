package app

import (
	"context"
	"gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"

	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/model"
)

func (s *Service) CreateUser(ctx context.Context, chatId int64) {

	if _, err := s.repo.GetUserByChatId(ctx, chatId); err == nil {
		return
	}

	scUser, _ := s.api.CreateUser(ctx, &api.CreateUserRequest{})

	user := model.User{
		ID:     scUser.ID,
		ChatID: chatId,
	}

	_ = s.repo.CreateUser(ctx, user)
}
