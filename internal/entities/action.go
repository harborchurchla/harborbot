package entities

const (
	GetTeamScheduleAction         = "get-team-schedule"
	GetWhoIsServingThisWeekAction = "get-who-is-serving-this-week"
)

type Action struct {
	ID          string
	Description string
}

type ActionResult struct {
	Message  string
	Metadata map[string]interface{}
}
