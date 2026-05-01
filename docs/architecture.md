# Architecture

## Layers

```
main.go              ── engine bootstrap, signal handling, log setup
   │
   ▼
api.Engine           (dappco.re/go/api) ── route registry, middleware, serve loop
   │
   ▼
DemoProvider         (provider.go) ── implements api.RouteGroup
   │
   ▼
gin.RouterGroup      ── route handlers (hello, status)
```

## Provider contract

A service provider implements three methods:

- `Name() string`         — identifier used in logs and admin views
- `BasePath() string`     — URL prefix (e.g. `/api/v1/demo`)
- `RegisterRoutes(*gin.RouterGroup)` — wire handlers under the prefix

`DemoProvider` is the reference implementation. Real providers swap the
struct, methods, and handler bodies; the contract is identical.

## Frontend

`static.go` embeds `ui/dist` (the Lit build output) and serves it through a
no-cache wrapper so live reload during development behaves correctly. In
production the no-cache wrapper can be swapped for a far-future-cache wrapper
once the build pipeline emits hashed asset filenames.
