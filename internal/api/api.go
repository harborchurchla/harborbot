package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	result, err := a.ActionService.ExecuteByID(action, ctx.Request.URL.Query())
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": fmt.Sprintf("error while executing action %s: %s", action, err.Error()),
		})
		return
	}

	ctx.JSON(200, result)
}
