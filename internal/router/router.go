package router

import (
	"context"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/app"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/live_action/task_creator"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/live_action/task_stage_creator"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/model"
	"strconv"
	"strings"
)

type router struct {
	service *app.Service
	repo    Repository
}

func NewRouter(service *app.Service, repo Repository) *router {
	return &router{
		service: service,
		repo:    repo,
	}
}

func (r *router) ProcessMessage(update tgbotapi.Update, ctx context.Context) []tgbotapi.Chattable {

	if action, err := r.repo.GetActionByChatID(ctx, update.Message.Chat.ID); err == nil {
		if action.Name == task_creator.Name {
			return r.processCreateTaskAction(update, ctx, action)
		}
		if action.Name == task_stage_creator.Name {
			return r.processCreateTaskStageAction(update, ctx, action)
		}
	}

	switch update.Message.Command() {
	case "start":
		r.service.CreateUser(ctx, update.Message.Chat.ID)
		return []tgbotapi.Chattable{tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать")}

	case "all_tasks":
		list := r.service.TasksList(ctx)
		if list == nil {
			return []tgbotapi.Chattable{tgbotapi.NewMessage(update.Message.Chat.ID, "Введите заново")}
		}
		return tasksListToTGText(update.Message.Chat.ID, list.Tasks, false)

	case "my_tasks":
		list, err := r.service.ChatTasks(ctx, update.Message.Chat.ID)
		if err != nil {
			return []tgbotapi.Chattable{tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка "+err.Error())}
		}
		if list == nil {
			return []tgbotapi.Chattable{tgbotapi.NewMessage(update.Message.Chat.ID, "Введите заново")}
		}
		return tasksListToTGText(update.Message.Chat.ID, list.Tasks, true)

	case "create_task":
		return r.processCreateTaskAction(
			update,
			ctx,
			model.ActiveLiveAction{
				Data:   []byte(`{}`),
				ChatID: update.Message.Chat.ID,
			})
	case "create_task_stage":
		return r.processCreateTaskStageAction(
			update,
			ctx,
			model.ActiveLiveAction{
				Data:   []byte(`{}`),
				ChatID: update.Message.Chat.ID,
			})
	}

	str := update.Message.Text
	switch {
	case strings.Contains(str, "Подписаться"):
		err := r.service.Subscribe(ctx, update.Message.Chat.ID, getTaskID(str))
		if err != nil {
			return []tgbotapi.Chattable{tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка "+err.Error())}
		}
		message := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы успешно подписались")
		return []tgbotapi.Chattable{message}

	case strings.Contains(str, "Отписаться"):
		err := r.service.Unsubscribe(ctx, update.Message.Chat.ID, getTaskID(str))
		if err != nil {
			return []tgbotapi.Chattable{tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка "+err.Error())}
		}
		message := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы успешно отписались")
		return []tgbotapi.Chattable{message}

	}
	return []tgbotapi.Chattable{tgbotapi.NewMessage(update.Message.Chat.ID, "Введите заново")}
}

func (r *router) processCreateTaskAction(update tgbotapi.Update, ctx context.Context, action model.ActiveLiveAction) []tgbotapi.Chattable {
	creator, err := task_creator.NewTaskCreator(action, r.service)
	if err != nil {
		return nil
	}

	returnMessages := creator.Process(ctx, update.Message)
	newAction := creator.GetCurrentAction()

	if creator.IsFinished() {
		r.repo.DeleteActionByChatID(ctx, update.Message.Chat.ID)
		if err != nil {
			return nil
		}
		return returnMessages
	}
	r.repo.CreateOrUpdateAction(ctx, newAction)

	return returnMessages
}

func (r *router) processCreateTaskStageAction(update tgbotapi.Update, ctx context.Context, action model.ActiveLiveAction) []tgbotapi.Chattable {
	creator, err := task_stage_creator.NewTaskCreator(action, r.service)
	if err != nil {
		return nil
	}

	returnMessages := creator.Process(ctx, update.Message)
	newAction := creator.GetCurrentAction()

	if creator.IsFinished() {
		r.repo.DeleteActionByChatID(ctx, update.Message.Chat.ID)
		if err != nil {
			return nil
		}
		return returnMessages
	}
	r.repo.CreateOrUpdateAction(ctx, newAction)

	return returnMessages
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

	if len(messages) == 0 {
		return nil
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
