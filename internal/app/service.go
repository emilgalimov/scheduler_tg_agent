package app

import "gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"

type Service struct {
	repo Repository
	api  api.SmartCalendarClient
}

func NewService(repo Repository, api api.SmartCalendarClient) *Service {
	return &Service{
		repo: repo,
		api:  api,
	}
}
