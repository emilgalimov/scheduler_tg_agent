package task_creator

import (
	"context"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/app"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/model"
)

const Name = "taskCreator"

type taskCreator struct {
	startMessage      state
	createName        state
	createDescription state
	final             state
	currentState      state

	activeAction model.ActiveLiveAction
	data         taskCreatorData

	service *app.Service
}

func NewTaskCreator(activeAction model.ActiveLiveAction, service *app.Service) (*taskCreator, error) {

	var data taskCreatorData

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

	createName := &createNameState{
		tc: tc,
	}

	createDescription := &createDescriptionState{
		tc: tc,
	}

	final := &finalState{
		tc: tc,
	}
	tc.createName = createName
	tc.createDescription = createDescription
	tc.startMessage = startMessage
	tc.final = final

	switch activeAction.State {
	case "createName":
		tc.currentState = createName
	case "createDescription":
		tc.currentState = createDescription
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
