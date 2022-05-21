package app

import (
	"context"
	"gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
)

func (s *Service) ChatTasks(ctx context.Context, chatId int64) (*api.TasksList, error) {
	user, _ := s.repo.GetUserByChatId(ctx, chatId)

	return s.api.GetUserTasksByTime(ctx, &api.GetUserTasksByTimeRequest{
		UserId:   user.ID,
		TimeFrom: nil,
		TimeTo:   nil,
	})
}
