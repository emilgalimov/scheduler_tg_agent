package app

import (
	"context"
	"gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
)

func (s *Service) TasksList(ctx context.Context) *api.TasksList {

	list, err := s.api.GetTasksList(ctx, &api.GetTasksListRequest{})

	if err != nil {
		return nil
	}

	return list
}
