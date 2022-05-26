package app

import (
	"context"
	"gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
)

func (s *Service) CreateTask(ctx context.Context, name string, description string) (uint64, error) {

	task, err := s.api.CreateTask(ctx, &api.Task{
		Name:        name,
		Description: description,
	})
	if err != nil {
		return 0, err
	}
	return task.ID, nil
}
