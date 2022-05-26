package task_stage_creator

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type createDescriptionState struct {
	tc *taskCreator
}

func (s *createDescriptionState) process(ctx context.Context, message *tgbotapi.Message) []tgbotapi.Chattable {
	s.tc.data.Description = message.Text

	s.tc.currentState = s.tc.createMinutesFromStart
	answerMessage := tgbotapi.NewMessage(message.Chat.ID, "Введите количество минут с начала задания")
	return []tgbotapi.Chattable{answerMessage}
}

func (s *createDescriptionState) name() string {
	return "createDescription"
}

func (s *createDescriptionState) isFinal() bool {
	return false
}
