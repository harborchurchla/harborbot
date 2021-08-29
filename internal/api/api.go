package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/harborchurchla/harborbot/internal/services"
)

type API struct {
	*gin.Engine
	*services.ScheduleService
}

func New(ss *services.ScheduleService) *API {
	a := &API{gin.Default(), ss}
	a.Engine.GET("/ping", a.ping)
	a.Engine.GET("/schedule/:team", a.getTeamSchedule)
	a.Engine.GET("/schedule/:team/upcoming", a.getUpcomingTeamSchedule)

	return a
}

func (a *API) ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
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
