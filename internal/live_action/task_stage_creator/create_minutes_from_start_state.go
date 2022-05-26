package task_stage_creator

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

type createMinutesFromStartState struct {
	tc *taskCreator
}

func (s *createMinutesFromStartState) process(ctx context.Context, message *tgbotapi.Message) []tgbotapi.Chattable {
	minutes, err := strconv.ParseUint(message.Text, 10, 64)

	if err != nil {
		return nil
	}

	s.tc.data.MinutesFromStart = minutes

	s.tc.currentState = s.tc.createDurationMinutes
	answerMessage := tgbotapi.NewMessage(message.Chat.ID, "Введите длительность этапа задачи")
	return []tgbotapi.Chattable{answerMessage}
}

func (s *createMinutesFromStartState) name() string {
	return "createMinutesFromStart"
}

func (s *createMinutesFromStartState) isFinal() bool {
	return false
}
