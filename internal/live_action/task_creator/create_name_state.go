package task_creator

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type createNameState struct {
	tc *taskCreator
}

func (s *createNameState) process(ctx context.Context, message *tgbotapi.Message) []tgbotapi.Chattable {
	s.tc.data.Name = message.Text

	s.tc.currentState = s.tc.createDescription
	answerMessage := tgbotapi.NewMessage(message.Chat.ID, "Введите описание")
	return []tgbotapi.Chattable{answerMessage}
}

func (s *createNameState) name() string {
	return "createName"
}

func (s *createNameState) isFinal() bool {
	return false
}
