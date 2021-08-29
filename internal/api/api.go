package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/harborchurchla/harborbot/internal/entities"
	"github.com/harborchurchla/harborbot/internal/services"
)

type API struct {
	*gin.Engine
	*services.ScheduleService
	*services.ActionService
}

func New(ss *services.ScheduleService, as *services.ActionService) *API {
	a := &API{gin.Default(), ss, as}
	a.Engine.GET("/", a.healthz)
	a.Engine.GET("/actions", a.listActions)
	a.Engine.GET("/actions/:id", a.executeAction)

	return a
}

func (a *API) healthz(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"ok": true,
	})
}

func (a *API) listActions(ctx *gin.Context) {
	ctx.JSON(200, a.ActionService.ListActions())
}

func (a *API) executeAction(ctx *gin.Context) {
	action := ctx.Param("id")
	switch action {
	case entities.GetTeamScheduleAction:
		team := ctx.Query("team")
		if team == "" {
			ctx.JSON(400, gin.H{
				"message": fmt.Sprintf("query param 'team' is required for action: %s", action),
			})
			return
		}
		result, err := a.ActionService.ExecuteGetTeamScheduleAction(team)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": fmt.Sprintf("error while executing action %s for team %s: %s", action, team, err.Error()),
			})
			return
		}
		ctx.JSON(200, result)
		return
	case entities.GetWhoIsServingThisWeekAction:
		team := ctx.Query("team")
		if team == "" {
			ctx.JSON(400, gin.H{
				"message": fmt.Sprintf("query param 'team' is required for action: %s", action),
			})
			return
		}
		result, err := a.ActionService.ExecuteGetWhoIsServingThisWeekAction(team)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": fmt.Sprintf("error while executing action %s for team %s: %s", action, team, err.Error()),
			})
			return
		}
		ctx.JSON(200, result)
		return
	default:
		ctx.JSON(400, gin.H{
			"message": fmt.Sprintf("error executing unknown action %s", action),
		})
	}
}
