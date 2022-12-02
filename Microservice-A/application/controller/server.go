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
	Running   bool   `json:"running"`
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
	c.router.POST("/command", func(ctx_res *gin.Context) {
		ctx_res.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx_res.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		ctx_res.Writer.Header().Set("Content-Type", "application/json")

		err := ctx_res.ShouldBindJSON(&command)
		if err != nil {
			ctx_res.JSON(http.StatusBadRequest, "Internal error")
			return
		}

		// changed scheduler
		if command.Running {
			lead.SetScheduler(command.Scheduler)
			lead.SetGetLimit(command.GetLimit)
		}

		if !command.Running {
			lead.RemoveScheduler()
		}

		// return response
		ctx_res.JSON(http.StatusOK, command)
	})
}

func (c *config) RunWebServer(addr string) {
	c.router.Run(addr)
}
