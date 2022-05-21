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

func (r *router) ProcessMessage(update tgbotapi.Update, ctx context.Context) string {
	switch update.Message.Command() {
	case "start":
		r.service.CreateUser(ctx, update.Message.Chat.ID)
		return "Добро пожаловать"
	case "all_tasks":
		list := r.service.TasksList(ctx)
		return tasksListToTGText(list.Tasks)
	case "my_tasks":
		r.service.ChatTasks(ctx, update.Message.Chat.ID)
	}

	str := update.Message.Text
	switch {
	case strings.Contains(str, "subscribe"):
		r.service.Subscribe(ctx, update.Message.Chat.ID, getTaskID(str))
	case strings.Contains(str, "unsubscribe"):
		r.service.Unsubscribe(ctx, update.Message.Chat.ID, getTaskID(str))
	}
	return "error"
}

func tasksListToTGText(tasks []*api.Task) string {

	var result string

	for _, task := range tasks {
		result += "Name: " + task.Name
		result += "\n"
		result += "Description: " + task.Description
		result += "\n\n"
		result += "Stages\n"

		for _, taskStage := range task.TaskStages {
			result += "Name: " + taskStage.Name
			result += "\n"
			result += "Description: " + taskStage.Description
			result += "\n"
			result += "MinutesFromStart: " + strconv.FormatUint(taskStage.MinutesFromStart, 10)
			result += "\n"
			result += "DurationMinutes: " + strconv.FormatUint(taskStage.DurationMinutes, 10)
			result += "\n\n"
		}
	}

	return result

}

func getTaskID(str string) uint64 {
	ID, _ := strconv.ParseUint(
		strings.Split(str, "_")[1],
		0,
		64,
	)

	return ID
}
