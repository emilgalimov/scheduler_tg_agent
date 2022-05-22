package app

import (
	"context"
	"gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (s *Service) ChatTasks(ctx context.Context, chatId int64) (*api.TasksList, error) {
	user, _ := s.repo.GetUserByChatId(ctx, chatId)
	t := time.Now()
	startOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1)

	return s.api.GetUserTasksByTime(ctx, &api.GetUserTasksByTimeRequest{
		UserId:   user.ID,
		TimeFrom: timestamppb.New(startOfDay),
		TimeTo:   timestamppb.New(endOfDay),
	})
}
