package router

import (
	"context"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/app"
	"strconv"
	"strings"
)

type router struct {
	service *app.Service
}

func NewRouter(service *app.Service) *router {
	return &router{service: service}
}

func (r *router) ProcessMessage(update tgbotapi.Update, ctx context.Context) []tgbotapi.Chattable {
	switch update.Message.Command() {
	case "start":
		r.service.CreateUser(ctx, update.Message.Chat.ID)
		return []tgbotapi.Chattable{tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать")}

	case "all_tasks":
		list := r.service.TasksList(ctx)
		return tasksListToTGText(update.Message.Chat.ID, list.Tasks, false)

	case "my_tasks":
		list, err := r.service.ChatTasks(ctx, update.Message.Chat.ID)
		if err != nil {
			return []tgbotapi.Chattable{tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка "+err.Error())}
		}
		return tasksListToTGText(update.Message.Chat.ID, list.Tasks, true)
	}

	str := update.Message.Text
	switch {
	case strings.Contains(str, "Подписаться"):
		err := r.service.Subscribe(ctx, update.Message.Chat.ID, getTaskID(str))
		if err != nil {
			return []tgbotapi.Chattable{tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка "+err.Error())}
		}
		message := tgbotapi.NewMessage(update.Message.Chat.ID, "Успешно")
		message.ReplyMarkup = tgbotapi.NewReplyKeyboard()
		return []tgbotapi.Chattable{message}

	case strings.Contains(str, "Отписаться"):
		r.service.Unsubscribe(ctx, update.Message.Chat.ID, getTaskID(str))
	}
	return []tgbotapi.Chattable{tgbotapi.NewMessage(update.Message.Chat.ID, "Введите заного")}
}

func tasksListToTGText(chatId int64, tasks []*api.Task, isUnsubscribe bool) (messages []tgbotapi.Chattable) {

	var buttons []tgbotapi.KeyboardButton

	for _, task := range tasks {
		var text string
		text += "ID: " + strconv.FormatUint(task.ID, 10)
		text += "\n"
		text += "Name: " + task.Name
		text += "\n"
		text += "Description: " + task.Description
		text += "\n\n"
		text += "Stages\n"

		for _, taskStage := range task.TaskStages {
			text += "Name: " + taskStage.Name
			text += "\n"
			text += "Description: " + taskStage.Description
			text += "\n"
			text += "MinutesFromStart: " + strconv.FormatUint(taskStage.MinutesFromStart, 10)
			text += "\n"
			text += "DurationMinutes: " + strconv.FormatUint(taskStage.DurationMinutes, 10)
			text += "\n\n"
		}

		message := tgbotapi.NewMessage(chatId, text)
		messages = append(messages, message)

		if isUnsubscribe {
			buttons = append(buttons, tgbotapi.NewKeyboardButton("Отписаться "+strconv.FormatUint(task.ID, 10)))
			continue
		}
		buttons = append(buttons, tgbotapi.NewKeyboardButton("Подписаться "+strconv.FormatUint(task.ID, 10)))
	}

	message := tgbotapi.NewMessage(chatId, "Для подписки выберите ID задачи")
	message.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

	messages = append(messages, message)
	return messages
}

func getTaskID(str string) uint64 {
	ID, _ := strconv.ParseUint(
		strings.Split(str, " ")[1],
		0,
		64,
	)

	return ID
}
