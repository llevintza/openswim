# OpenSwim API

Go HTTP + WebSocket service backed by Postgres.

## Stack

- Go module: `github.com/llevintza/openswim/apps/api`
- Postgres via `pgx` pool
- Migrations: `golang-migrate` (`migrations/`; empty until first schema)
- Config: `DATABASE_URL`, `PORT` (default `8080`), `JWT_SECRET`

## Develop in a Dev Container

Do **not** install Go or Postgres on the host. Open the repo in Cursor/VS Code with Dev Containers (see [`.devcontainer/`](../../.devcontainer/) and [onboarding](../../docs/onboarding.md)).

Inside the container:

```bash
cd apps/api
go run ./cmd/api
```

Env is pre-set to the Compose Postgres service. Then:

```bash
curl -s localhost:8080/health
```

## Endpoints

| Path | Status | Notes |
|------|--------|-------|
| `GET /health` | Implemented | Client-facing; pings DB |
| `GET /livez` | Stub `501` | Reserved — [ADR 0002](../../docs/adr/0002-api-observability.md), `E12-F5-T3` |
| `GET /readyz` | Stub `501` | Reserved — ADR 0002, `E12-F5-T3` |
| `GET /metrics` | Stub `501` | Reserved Prometheus — ADR 0002, `E12-F5-T2` |

## Production image

```bash
docker build -t openswim-api -f apps/api/Dockerfile apps/api
```

Do not add a Docker `HEALTHCHECK` against `/livez` until `E12-F5-T3` lands.

## Layout

```text
cmd/api/           process entrypoint
internal/config/   env loading
internal/db/       pgx pool
internal/migrate/  migration runner stub
internal/httpserver/
internal/logging/  JSON slog baseline
internal/metrics/  /metrics stub
internal/health/   /livez /readyz stubs
migrations/        SQL migrations (none yet)
```
