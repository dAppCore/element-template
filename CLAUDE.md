# CLAUDE.md

This file provides guidance to Claude Code when working with this repository.

## Project Overview

Starter template for building custom HTML elements backed by a Go API using the Core ecosystem. Clone this to create a new service provider that can plug into core/ide.

## Build & Development

```bash
# Build UI (Lit custom element)
cd ui && npm install && npm run build

# Run development server
go run . --port 8080

# Build binary
core build

# Quality assurance
core go qa
```

## Architecture

- **`main.go`** — Creates `api.Engine`, registers providers, serves embedded UI
- **`provider.go`** — Example `DemoProvider` implementing `api.RouteGroup`
- **`static.go`** — Static file serving helper
- **`ui/`** — Lit custom element that talks to the Go API via fetch

## Creating Your Own Provider

1. Rename `DemoProvider` in `provider.go`
2. Update `Name()`, `BasePath()`, and routes in `RegisterRoutes()`
3. Update the Lit element in `ui/src/index.ts` to match your API
4. Update the custom element tag in `ui/index.html`

## Conventions

- UK English in all user-facing strings
- EUPL-1.2 licence
- Conventional commits
