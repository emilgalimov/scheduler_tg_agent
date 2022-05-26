package task_stage_creator

type taskStageCreatorData struct {
	TaskID           uint64 `json:"task_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	MinutesFromStart uint64 `json:"minutes_from_start"`
	DurationMinutes  uint64 `json:"duration_minutes"`
}
