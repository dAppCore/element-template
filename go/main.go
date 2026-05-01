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
	"flag"
	"os/signal"
	"syscall"

	core "dappco.re/go"
	api "dappco.re/go/api"
	"dappco.re/go/log"
)

func main() {
	port := flag.String("port", "8080", "HTTP server port")
	flag.Parse()

	logger := log.New(log.Options{Level: log.LevelInfo})
	addr := ":" + *port

	// Create API engine with middleware
	engine, err := api.New(
		api.WithAddr(addr),
		api.WithCORS(),
		api.WithStatic("/", "ui/dist"),
	)
	if err != nil {
		logger.Error("failed to create API engine", "error", err)
		core.Exit(1)
	}

	// Register the demo provider
	engine.Register(&DemoProvider{})

	// Start server
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger.Info("starting server", "addr", addr)

	go func() {
		if err := engine.Serve(ctx); err != nil {
			logger.Error("server error", "error", err)
		}
	}()

	<-ctx.Done()
	logger.Info("shutting down")
}
