# Go API coding notes

Applies when working under `apps/api`.

## Stack

- Go module root: `apps/api` (`github.com/llevintza/openswim/apps/api`).
- Postgres for persistence; SQL migrations under `apps/api/migrations/` via `golang-migrate`.
- HTTP JSON API + WebSocket hub in-process for MVP.
- Config via environment variables (`DATABASE_URL`, `PORT`, `JWT_SECRET`).
- Develop via the root [Dev Container](../../../.devcontainer/) (host should not need a local Go install).

## Conventions

- Standard Go layout: `cmd/`, `internal/`, keep public surface small.
- Context on all request-scoped work; no bare `panic` for expected errors.
- Validate authn/authz at HTTP boundary using membership + meet device sessions.
- Prefer explicit SQL/migrations over heavy ORMs unless the team standardizes later.
- Generate or hand-sync handlers against `packages/contracts` OpenAPI.

## HTTP surface (Phase 0)

| Path | Role |
|------|------|
| `GET /health` | Product/client health (DB ping) |
| `GET /livez` | Container liveness (stub until E12-F5-T3) |
| `GET /readyz` | Container readiness (stub until E12-F5-T3) |
| `GET /metrics` | Prometheus scrape (stub until E12-F5-T2) |

Logging baseline is JSON `slog` to stdout. Full logging, metrics, and probes: [ADR 0002](../../adr/0002-api-observability.md). Do not add `prometheus/client_golang` until E12-F5-T2.

## Meet realtime

- Channels such as `meet:{id}` and `team:{id}:chat`.
- Starter start pulses and lane times are latency-sensitive; keep payloads small and timestamps clear (server time + device time when useful).

## See also

- [`../../domain-model.md`](../../domain-model.md)
- Root [`AGENTS.md`](../../../AGENTS.md)
