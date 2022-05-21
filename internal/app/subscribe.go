package app

import (
	"context"
	"gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
)

func (s *Service) Subscribe(ctx context.Context, chatId int64, taskID uint64) error {
	user, _ := s.repo.GetUserByChatId(ctx, chatId)

	_, err := s.api.SubscribeUser(ctx, &api.SubscribeUserRequest{UserID: user.ID, TaskID: taskID})

	return err
}
