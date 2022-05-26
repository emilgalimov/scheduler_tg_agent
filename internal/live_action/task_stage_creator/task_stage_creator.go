package task_stage_creator

import (
	"context"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/app"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/model"
)

const Name = "taskStageCreator"

type taskCreator struct {
	startMessage           state
	createTaskID           state
	createName             state
	createDescription      state
	createMinutesFromStart state
	createDurationMinutes  state
	final                  state

	currentState state

	activeAction model.ActiveLiveAction
	data         taskStageCreatorData

	service *app.Service
}

// NewTaskCreator TODO вынести интерфейс многоступенчатого криейтера
func NewTaskCreator(activeAction model.ActiveLiveAction, service *app.Service) (*taskCreator, error) {

	var data taskStageCreatorData

	err := json.Unmarshal(activeAction.Data, &data)
	if err != nil {
		return nil, err
	}

	tc := &taskCreator{
		activeAction: activeAction,
		data:         data,
		service:      service,
	}

	startMessage := &startMessageState{
		tc: tc,
	}

	createTaskID := &createTaskIDState{
		tc: tc,
	}
	createName := &createNameState{
		tc: tc,
	}

	createDescription := &createDescriptionState{
		tc: tc,
	}

	createMinutesFromStart := &createMinutesFromStartState{
		tc: tc,
	}

	createDurationMinutes := &createDurationMinutesState{
		tc: tc,
	}

	final := &finalState{
		tc: tc,
	}

	//TODO привязывать друг к другу как в цепочке событий
	//TODO вынести вступительное сообщение в следующий статус
	tc.startMessage = startMessage
	tc.createTaskID = createTaskID
	tc.createName = createName
	tc.createDescription = createDescription
	tc.createMinutesFromStart = createMinutesFromStart
	tc.createDurationMinutes = createDurationMinutes
	tc.final = final

	switch activeAction.State {
	case "createTaskID":
		tc.currentState = createTaskID
	case "createName":
		tc.currentState = createName
	case "createDescription":
		tc.currentState = createDescription
	case "createMinutesFromStart":
		tc.currentState = createMinutesFromStart
	case "createDurationMinutes":
		tc.currentState = createDurationMinutes
	case "final":
		tc.currentState = final
	default:
		tc.currentState = startMessage
	}

	return tc, nil
}

func (tc *taskCreator) setState(s state) {
	tc.currentState = s
}

func (tc *taskCreator) Process(ctx context.Context, message *tgbotapi.Message) []tgbotapi.Chattable {
	return tc.currentState.process(ctx, message)
}

func (tc *taskCreator) GetCurrentAction() model.ActiveLiveAction {
	action := tc.activeAction
	action.Name = Name
	action.State = tc.currentState.name()
	action.Data, _ = json.Marshal(tc.data)

	return action
}

func (tc *taskCreator) IsFinished() bool {
	return tc.currentState.isFinal()
}
