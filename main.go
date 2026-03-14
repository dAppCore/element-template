// SPDX-License-Identifier: EUPL-1.2

// core-element-template demonstrates the service provider pattern.
// It serves a Lit custom element backed by a Go API using core/go-api.
//
// Usage:
//
//	go run . [--port 8080]
package main

import (
	"context"
	"embed"
	"flag"
	"io/fs"
	"os"
	"os/signal"
	"syscall"

	api "forge.lthn.ai/core/api"
	"forge.lthn.ai/core/go-log"
	"github.com/gin-gonic/gin"
)

//go:embed all:ui/dist
var uiAssets embed.FS

func main() {
	port := flag.String("port", "8080", "HTTP server port")
	flag.Parse()

	logger := log.New("element-template")

	// Create API engine with middleware
	engine, err := api.New(
		api.WithCORS(),
	)
	if err != nil {
		logger.Error("failed to create API engine", "error", err)
		os.Exit(1)
	}

	// Register the demo provider
	engine.Register(&DemoProvider{})

	// Serve the Lit custom element UI as static files
	staticFS, err := fs.Sub(uiAssets, "ui/dist")
	if err != nil {
		logger.Error("failed to load UI assets", "error", err)
		os.Exit(1)
	}
	engine.Router().NoRoute(gin.WrapH(
		noCache(staticFS),
	))

	// Start server
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	addr := ":" + *port
	logger.Info("starting server", "addr", addr)

	go func() {
		if err := engine.Serve(ctx, addr); err != nil {
			logger.Error("server error", "error", err)
		}
	}()

	<-ctx.Done()
	logger.Info("shutting down")
}
