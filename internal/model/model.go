package model

type User struct {
	ID     uint64
	ChatID int64
}

type ActiveLiveAction struct {
	ChatID int64
	Name   string
	State  string
	Data   []byte
}
