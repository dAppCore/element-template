# AGENTS.md — element-template

A canonical scaffold for a `github.com/dappcore/api/go` service provider.

## What this is

`element-template` demonstrates the smallest viable service-provider shape:
a `DemoProvider` struct implementing `api.RouteGroup` (Name / BasePath /
RegisterRoutes), wired into a `github.com/dappcore/api/go` engine that serves both
the JSON API and a Lit custom-element frontend.

Copy this repo as the starting point for new providers — replace
`DemoProvider` with the real provider type, swap the static `ui/` for the
real frontend, and you have a v0.9.0-shaped service.

## Code shape

- `main.go`             — engine bootstrap (api.New + Register + Serve)
- `provider.go`         — `DemoProvider` implements api.RouteGroup
- `provider_test.go`    — Test triplets per public method
- `provider_example_test.go` — Example<Symbol> coverage
- `static.go`           — `ui/dist` embed + no-cache wrapper

## Dependencies (v0.9.0)

- `github.com/dappcore/go`        — core primitives, Exit, Println
- `github.com/dappcore/api/go`    — API engine + RouteGroup contract
- `github.com/dappcore/go-log/go`    — structured logger

## Audit contract

Every public symbol on `DemoProvider` (Name / BasePath / RegisterRoutes)
has a Test<File>_<Symbol>_{Good,Bad,Ugly} triplet in `provider_test.go`
and an `Example<Symbol>` in `provider_example_test.go`. `audit.sh`
must report `verdict: COMPLIANT` before any tag bump.
