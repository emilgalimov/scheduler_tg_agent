package app

import (
	"context"
	"gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
)

func (s *Service) Unsubscribe(ctx context.Context, chatID int64, taskID uint64) error {
	user, _ := s.repo.GetUserByChatId(ctx, chatID)

	_, err := s.api.UnsubscribeUser(ctx, &api.UnsubscribeUserRequest{UserID: user.ID, TaskID: taskID})
	return err
}
