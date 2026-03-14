# Core Element Template

Starter template for building custom HTML elements backed by a Go API. Part of the [Core ecosystem](https://core.help).

## Quick Start

```bash
# Clone and rename
git clone https://forge.lthn.ai/core/element-template.git my-element
cd my-element

# Install UI dependencies and build
cd ui && npm install && npm run build && cd ..

# Run
go run . --port 8080
```

Open `http://localhost:8080` — you'll see the `<core-demo-element>` fetching data from the Go API.

## What's Included

| Component | Technology | Purpose |
|-----------|-----------|---------|
| Go backend | core/go-api (Gin) | REST API with CORS, middleware |
| Custom element | Lit 3 | Self-contained web component |
| Build config | `.core/build.yaml` | Cross-platform binary builds |

## Making It Yours

1. Update `go.mod` module path
2. Rename `DemoProvider` in `provider.go` — implement your API
3. Rename `CoreDemoElement` in `ui/src/index.ts` — implement your UI
4. Update the element tag in `ui/index.html`

## Service Provider Pattern

The `DemoProvider` implements `api.RouteGroup`:

```go
func (p *DemoProvider) Name() string { return "demo" }
func (p *DemoProvider) BasePath() string { return "/api/v1/demo" }
func (p *DemoProvider) RegisterRoutes(rg *gin.RouterGroup) {
    rg.GET("/hello", p.hello)
}
```

Register it with `engine.Register(&DemoProvider{})` and it gets middleware, CORS, and OpenAPI for free. The same provider can plug into core/ide's registry.

## Licence

EUPL-1.2
