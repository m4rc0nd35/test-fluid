package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m4rc0nd35/test-fluid/application/adapter"
	"github.com/m4rc0nd35/test-fluid/application/toolkit"
)

type config struct {
	router *gin.Engine
}

type Command struct {
	Pause     bool   `json:"pause"`
	Scheduler string `json:"cron"`
	GetLimit  int    `json:"getLimit"`
}

func NewWebServer() *config {
	return &config{
		gin.New(), // Init contex gin
	}
}

func (c *config) Webserver(lead adapter.LeadAdapter) {
	defer toolkit.Recover("WebServer A")

	var command Command

	// c.router.Use(gin.Logger())
	c.router.POST("/setting", func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		ctx.Writer.Header().Set("Content-Type", "application/json")

		err := ctx.ShouldBindJSON(&command)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Internal error")
			return
		}

		lead.SetScheduler(command.Scheduler)
		lead.SetGetLimit(command.GetLimit)
		lead.Pause(command.Pause)

		// return response
		ctx.JSON(http.StatusOK, command)
	})
}

func (c *config) RunWebServer(addr string) {
	c.router.Run(addr)
}
