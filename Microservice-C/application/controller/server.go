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

func NewWebServer() *config {
	return &config{
		gin.New(), // Init contex gin
	}
}

func (c *config) DataLoggerOneWS(lead adapter.DataLoggerDomain) {
	defer toolkit.Recover("WebServer C")

	// c.router.Use(gin.Logger())
	c.router.GET("/datalogger/:uuid", func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		ctx.Writer.Header().Set("Content-Type", "application/json")

		uuid, err := ctx.Params.Get("uuid")
		if !err {
			ctx.JSON(http.StatusBadRequest, "Internal error")
			return
		}

		// return response
		ctx.JSON(http.StatusOK, lead.FindDataLoggerById(uuid))
	})
}

func (c *config) DataLoggerStatsWS(lead adapter.DataLoggerDomain) {
	defer toolkit.Recover("WebServer C")

	// c.router.Use(gin.Logger())
	c.router.GET("/datalogger/stats", func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		ctx.Writer.Header().Set("Content-Type", "application/json")

		// return response
		ctx.JSON(http.StatusOK, lead.DataLoggerStats())
	})
}

func (c *config) LeadOneWS(lead adapter.LeadDomain) {
	defer toolkit.Recover("WebServer C")

	// c.router.Use(gin.Logger())
	c.router.GET("/lead/:uuid", func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		ctx.Writer.Header().Set("Content-Type", "application/json")

		uuid, err := ctx.Params.Get("uuid")
		if !err {
			ctx.JSON(http.StatusBadRequest, "Internal error")
			return
		}

		// return response
		ctx.JSON(http.StatusOK, lead.FindOneLead(uuid))
	})
}

func (c *config) LeadAllWS(lead adapter.LeadDomain) {
	defer toolkit.Recover("WebServer C")

	// c.router.Use(gin.Logger())
	c.router.GET("/lead/all", func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		ctx.Writer.Header().Set("Content-Type", "application/json")

		// return response
		ctx.JSON(http.StatusOK, lead.FindAllLead())
	})
}

func (c *config) RunWebServer(addr string) {
	c.router.Run(addr)
}
