package task_stage_creator

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

type createTaskIDState struct {
	tc *taskCreator
}

func (s *createTaskIDState) process(ctx context.Context, message *tgbotapi.Message) []tgbotapi.Chattable {
	var err error

	s.tc.data.TaskID, err = strconv.ParseUint(message.Text, 10, 64)

	if err != nil {
		return nil
	}

	s.tc.currentState = s.tc.createName
	answerMessage := tgbotapi.NewMessage(message.Chat.ID, "Введите Название")
	return []tgbotapi.Chattable{answerMessage}
}

func (s *createTaskIDState) name() string {
	return "createTaskID"
}

func (s *createTaskIDState) isFinal() bool {
	return false
}
