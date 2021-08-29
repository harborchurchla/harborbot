package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/harborchurchla/harborbot/internal/services"
	"time"
)

type API struct {
	*gin.Engine
	*services.ScheduleService
}

const (
	GetTeamScheduleAction = "get-team-schedule"
)

func New(ss *services.ScheduleService) *API {
	a := &API{gin.Default(), ss}
	a.Engine.GET("/ping", a.ping)
	a.Engine.GET("/actions", a.listActions)
	a.Engine.GET("/actions/:id", a.executeAction)
	a.Engine.GET("/schedule/:team", a.getTeamSchedule)
	a.Engine.GET("/schedule/:team/upcoming", a.getUpcomingTeamSchedule)

	return a
}

func (a *API) ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (a *API) listActions(ctx *gin.Context) {
	ctx.JSON(200, []gin.H{{
		"id": GetTeamScheduleAction,
	}})
}

func (a *API) executeAction(ctx *gin.Context) {
	action := ctx.Param("id")
	switch action {
	case GetTeamScheduleAction:
		team := ctx.Query("team")
		schedule, err := a.ScheduleService.GetByTeam(team)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": fmt.Sprintf("error while fetching %s team schedule: %s", team, err.Error()),
			})
			return
		}
		entries := schedule.GetFutureEntries()
		scheduleText := ""
		for _, entry := range entries {
			scheduleText += fmt.Sprintf("%s: %s\n", time.Time(entry.Date).Format("1/2/2006"), entry.TeamMembers)
		}
		sheetText := fmt.Sprintf("_Feel free to make changes on the google sheet:_ https://docs.google.com/spreadsheets/d/%s", a.ScheduleService.GetSheetID())

		ctx.JSON(200, gin.H{
			"message": fmt.Sprintf("_Here's the upcoming schedule for the %s team:_\n\n%s\n\n%s", team, scheduleText, sheetText),
			"metadata": gin.H{
				"schedule": schedule,
			}})
		return
	default:
		ctx.JSON(400, gin.H{
			"message": fmt.Sprintf("error executing unknown action %s", action),
		})
	}
}

func (a *API) getTeamSchedule(ctx *gin.Context) {
	team := ctx.Param("team")
	schedule, err := a.ScheduleService.GetByTeam(team)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": fmt.Sprintf("error while fetching %s team schedule: %s", team, err.Error()),
		})
		return
	}

	ctx.JSON(200, schedule)
}

func (a *API) getUpcomingTeamSchedule(ctx *gin.Context) {
	ctx.JSON(400, gin.H{
		"message": "not-implemented",
	})
}
