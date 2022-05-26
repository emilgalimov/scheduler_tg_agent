package task_creator

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type startMessageState struct {
	tc *taskCreator
}

func (s *startMessageState) process(ctx context.Context, message *tgbotapi.Message) []tgbotapi.Chattable {
	s.tc.currentState = s.tc.createName
	answerMessage := tgbotapi.NewMessage(message.Chat.ID, "Введите название")
	return []tgbotapi.Chattable{answerMessage}
}

func (s *startMessageState) name() string {
	return "startMessage"
}

func (s *startMessageState) isFinal() bool {
	return false
}
