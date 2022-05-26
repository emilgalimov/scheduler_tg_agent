package task_stage_creator

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type state interface {
	process(ctx context.Context, message *tgbotapi.Message) []tgbotapi.Chattable
	name() string
	isFinal() bool
}
