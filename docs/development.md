# Development

## Run locally

```bash
go run . --port 8080
```

The server listens on `:8080`, serves `/api/v1/demo/{hello,status}`, and
serves the embedded `ui/dist` from `/`.

## Test

```bash
GOWORK=off GOPROXY=direct GOSUMDB=off go -C go test -count=1 -short ./...
```

## Audit

The repo must pass `audit.sh` from core/go before any tag bump:

```bash
bash /Users/snider/Code/core/go/tests/cli/v090-upgrade/audit.sh .
```

Verdict must be `COMPLIANT` (every counter at 0).

## Cutting a release

1. Verify build / vet / test / audit all clean
2. `git tag -a vX.Y.Z -m "vX.Y.Z — <summary>"`
3. `git push <remote> dev vX.Y.Z` for each of homelab / origin / github
4. Open dev → main PR on github, admin-merge
