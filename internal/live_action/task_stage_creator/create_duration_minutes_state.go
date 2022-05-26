package task_stage_creator

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
	"strconv"
)

type createDurationMinutesState struct {
	tc *taskCreator
}

func (s *createDurationMinutesState) process(ctx context.Context, message *tgbotapi.Message) []tgbotapi.Chattable {
	minutes, err := strconv.ParseUint(message.Text, 10, 64)

	if err != nil {
		return nil
	}

	s.tc.data.MinutesFromStart = minutes

	id, err := s.tc.service.CreateTaskStage(ctx, &api.TaskStage{
		Name:             s.tc.data.Name,
		Description:      s.tc.data.Description,
		MinutesFromStart: s.tc.data.MinutesFromStart,
		DurationMinutes:  s.tc.data.DurationMinutes,
	}, s.tc.data.TaskID)

	if err != nil {
		return nil
	}

	s.tc.currentState = s.tc.final
	answerMessage := tgbotapi.NewMessage(
		message.Chat.ID,
		fmt.Sprintf("Этап задачи успешно создан с ID = %v", id),
	)

	return []tgbotapi.Chattable{answerMessage}
}

func (s *createDurationMinutesState) name() string {
	return "createDurationMinutes"
}

func (s *createDurationMinutesState) isFinal() bool {
	return false
}
