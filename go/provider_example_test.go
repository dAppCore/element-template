// SPDX-License-Identifier: EUPL-1.2

package main

import (
	core "dappco.re/go"
)

// ExampleDemoProvider_Name shows how to retrieve a provider's identifier.
//
// Output: demo
func ExampleDemoProvider_Name() {
	p := &DemoProvider{}
	core.Println(p.Name())
	// Output: demo
}

// ExampleDemoProvider_BasePath shows the URL prefix for a provider's routes.
//
// Output: /api/v1/demo
func ExampleDemoProvider_BasePath() {
	p := &DemoProvider{}
	core.Println(p.BasePath())
	// Output: /api/v1/demo
}

// ExampleDemoProvider_RegisterRoutes shows DemoProvider satisfies the
// api.RouteGroup contract by exposing Name / BasePath / RegisterRoutes.
// Construction-only example — no HTTP traffic exchanged.
func ExampleDemoProvider_RegisterRoutes() {
	p := &DemoProvider{}
	core.Println(p.Name(), p.BasePath())
	// Output: demo /api/v1/demo
}
