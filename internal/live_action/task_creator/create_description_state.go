package task_creator

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type createDescriptionState struct {
	tc *taskCreator
}

func (s *createDescriptionState) process(ctx context.Context, message *tgbotapi.Message) []tgbotapi.Chattable {
	s.tc.data.Description = message.Text

	id, err := s.tc.service.CreateTask(ctx, s.tc.data.Name, s.tc.data.Description)
	if err != nil {
		return nil
	}

	s.tc.currentState = s.tc.final
	answerMessage := tgbotapi.NewMessage(
		message.Chat.ID,
		fmt.Sprintf("Задача успешно создана с ID = %v", id),
	)
	return []tgbotapi.Chattable{answerMessage}
}

func (s *createDescriptionState) name() string {
	return "createDescription"
}

func (s *createDescriptionState) isFinal() bool {
	return false
}
