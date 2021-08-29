package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/harborchurchla/harborbot/internal/entities"
	"github.com/harborchurchla/harborbot/internal/utils"
	"sort"
	"time"
)

type ActionService struct {
	*ScheduleService
}

func NewActionService(ss *ScheduleService) *ActionService {
	return &ActionService{ss}
}

func (s *ActionService) ListActions() []*entities.Action {
	return []*entities.Action{
		{ID: entities.GetTeamScheduleAction, Description: ""},
		{ID: entities.GetWhoIsServingThisWeekAction, Description: ""},
	}
}

func (s *ActionService) ExecuteByID(id string, params map[string][]string) (*entities.ActionResult, error) {
	switch id {
	case entities.GetTeamScheduleAction:
		if len(params["team"]) != 1 {
			return nil, fmt.Errorf("one param 'team' is required for action %s", id)
		}
		return s.ExecuteGetTeamScheduleAction(params["team"][0])
	case entities.GetWhoIsServingThisWeekAction:
		if len(params["team"]) != 1 {
			return nil, fmt.Errorf("one param 'team' is required for action %s", id)
		}
		return s.ExecuteGetWhoIsServingThisWeekAction(params["team"][0])
	default:
		return nil, fmt.Errorf("error executing unknown action %s", id)
	}
}

func (s *ActionService) ExecuteGetTeamScheduleAction(team string) (*entities.ActionResult, error) {
	schedule, err := s.ScheduleService.GetByTeam(team)
	if err != nil {
		return nil, fmt.Errorf("error while fetching %s team schedule: %v", team, err)
	}
	entries := schedule.GetFutureEntries()
	scheduleText := ""
	for _, entry := range entries {
		scheduleText += fmt.Sprintf("%s: *%s*\n", time.Time(entry.Date).Format("1/2/2006"), utils.ReplaceNameWithSlackMention(entry.TeamMembers))
	}
	sheetText := fmt.Sprintf("_Feel free to make changes on the google sheet:_ https://docs.google.com/spreadsheets/d/%s", s.ScheduleService.GetSheetID())

	return &entities.ActionResult{
		Message: fmt.Sprintf(
			"_Here's the upcoming schedule for the %s team:_\n\n%s\n\n%s",
			team,
			scheduleText,
			sheetText,
		),
		Metadata: gin.H{
			"schedule": schedule,
		},
	}, nil
}

func (s *ActionService) ExecuteGetWhoIsServingThisWeekAction(team string) (*entities.ActionResult, error) {
	schedule, err := s.ScheduleService.GetByTeam(team)
	if err != nil {
		return nil, fmt.Errorf("error while fetching %s team schedule: %v", team, err)
	}
	entries := schedule.GetFutureEntries()
	sort.SliceStable(entries, func(i, j int) bool {
		return time.Time(entries[i].Date).Before(time.Time(entries[j].Date))
	})
	next := entries[0]
	sheetText := fmt.Sprintf("_Feel free to make changes on the google sheet:_ https://docs.google.com/spreadsheets/d/%s", s.ScheduleService.GetSheetID())

	return &entities.ActionResult{
		Message: fmt.Sprintf(
			"_Here's who's serving this %s:_\n\n*%s*\n\n%s",
			time.Time(next.Date).Format("Mon Jan _2"),
			utils.ReplaceNameWithSlackMention(next.TeamMembers),
			sheetText,
		),
		Metadata: gin.H{
			"schedule": schedule,
		},
	}, nil
}
