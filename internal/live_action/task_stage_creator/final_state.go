package task_stage_creator

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type finalState struct {
	tc *taskCreator
}

func (s *finalState) process(context.Context, *tgbotapi.Message) []tgbotapi.Chattable {
	return nil
}

func (s *finalState) name() string {
	return "final"
}

func (s *finalState) isFinal() bool {
	return true
}
