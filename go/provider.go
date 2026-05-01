// SPDX-License-Identifier: EUPL-1.2

package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// DemoProvider is an example service provider that exposes a REST API.
// Replace this with your own provider implementation.
type DemoProvider struct{}

// Name implements api.RouteGroup.
func (p *DemoProvider) Name() string { return "demo" }

// BasePath implements api.RouteGroup.
func (p *DemoProvider) BasePath() string { return "/api/v1/demo" }

// RegisterRoutes implements api.RouteGroup.
func (p *DemoProvider) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/hello", p.hello)
	rg.GET("/status", p.status)
}

func (p *DemoProvider) hello(c *gin.Context) {
	name := c.DefaultQuery("name", "World")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, " + name + "!",
	})
}

func (p *DemoProvider) status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "running",
		"uptime": time.Since(startTime).String(),
	})
}

var startTime = time.Now()
