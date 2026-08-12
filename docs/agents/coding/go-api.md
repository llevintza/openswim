# Go API coding notes

Applies when working under `apps/api`.

## Stack

- Go module root: `apps/api` (not initialized until Phase 0 boot is requested).
- Postgres for persistence; migrations checked into the API tree when added.
- HTTP JSON API + WebSocket hub in-process for MVP.
- Config via environment variables (`DATABASE_URL`, `PORT`, auth secrets).

## Conventions (when code exists)

- Standard Go layout: `cmd/`, `internal/`, keep public surface small.
- Context on all request-scoped work; no bare `panic` for expected errors.
- Validate authn/authz at HTTP boundary using membership + meet device sessions.
- Prefer explicit SQL/migrations over heavy ORMs unless the team standardizes later.
- Generate or hand-sync handlers against `packages/contracts` OpenAPI.

## Meet realtime

- Channels such as `meet:{id}` and `team:{id}:chat`.
- Starter start pulses and lane times are latency-sensitive; keep payloads small and timestamps clear (server time + device time when useful).

## See also

- [`../../domain-model.md`](../../domain-model.md)
- Root [`AGENTS.md`](../../../AGENTS.md)
