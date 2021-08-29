package entities

const (
	GetTeamScheduleAction         = "get-team-schedule"
	GetWhoIsServingThisWeekAction = "get-who-is-serving-this-week"
)

type Action struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type ActionResult struct {
	Message  string                 `json:"message"`
	Metadata map[string]interface{} `json:"metadata"`
}
