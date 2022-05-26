package app

import (
	"context"
	"gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
)

func (s *Service) CreateTaskStage(ctx context.Context, stage *api.TaskStage, taskId uint64) (uint64, error) {

	taskStage, err := s.api.CreateTaskStage(ctx, &api.CreateTaskStageRequest{
		TaskID:    taskId,
		TaskStage: stage,
	})
	if err != nil {
		return 0, err
	}
	return taskStage.ID, nil
}
