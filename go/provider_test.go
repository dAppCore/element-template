// SPDX-License-Identifier: EUPL-1.2

package main

import (
	. "dappco.re/go"
	"github.com/gin-gonic/gin"
)

func TestProvider_DemoProvider_Name_Good(t *T) {
	p := &DemoProvider{}
	got := p.Name()

	AssertEqual(t, "demo", got)
	AssertNotContains(t, got, "/")
}

func TestProvider_DemoProvider_Name_Bad(t *T) {
	var p *DemoProvider
	got := p.Name()

	AssertEqual(t, "demo", got)
	AssertNotEmpty(t, got)
}

func TestProvider_DemoProvider_Name_Ugly(t *T) {
	first := (&DemoProvider{}).Name()
	second := (&DemoProvider{}).Name()

	AssertEqual(t, first, second)
	AssertNotEmpty(t, second)
}

func TestProvider_DemoProvider_BasePath_Good(t *T) {
	p := &DemoProvider{}
	got := p.BasePath()

	AssertEqual(t, "/api/v1/demo", got)
	AssertContains(t, got, p.Name())
}

func TestProvider_DemoProvider_BasePath_Bad(t *T) {
	var p *DemoProvider
	got := p.BasePath()

	AssertEqual(t, "/api/v1/demo", got)
	AssertContains(t, got, "demo")
}

func TestProvider_DemoProvider_BasePath_Ugly(t *T) {
	path := (&DemoProvider{}).BasePath()

	AssertTrue(t, HasPrefix(path, "/"))
	AssertFalse(t, HasSuffix(path, "/"))
}

func TestProvider_DemoProvider_RegisterRoutes_Good(t *T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	group := router.Group((&DemoProvider{}).BasePath())

	(&DemoProvider{}).RegisterRoutes(group)

	routes := router.Routes()
	AssertLen(t, routes, 2)
	AssertEqual(t, "GET", routes[0].Method)
	AssertEqual(t, "/api/v1/demo/hello", routes[0].Path)
	AssertEqual(t, "GET", routes[1].Method)
	AssertEqual(t, "/api/v1/demo/status", routes[1].Path)
}

func TestProvider_DemoProvider_RegisterRoutes_Bad(t *T) {
	var group *gin.RouterGroup

	AssertPanics(t, func() {
		(&DemoProvider{}).RegisterRoutes(group)
	})
}

func TestProvider_DemoProvider_RegisterRoutes_Ugly(t *T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	group := router.Group("/:tenant/demo")

	(&DemoProvider{}).RegisterRoutes(group)

	routes := router.Routes()
	AssertLen(t, routes, 2)
	AssertEqual(t, "/:tenant/demo/hello", routes[0].Path)
	AssertEqual(t, "/:tenant/demo/status", routes[1].Path)
}
