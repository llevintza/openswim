# 0002. API observability (logs, metrics, probes)

- Status: Accepted (decision recorded; **implementation tracked by E12-F5-T1–T3**)
- Date: 2026-08-12

## Context

The OpenSwim API runs as a container in self-host and hosted environments. Operators need:

1. **Logs** that are machine-parseable in container stdout aggregators.
2. **Metrics** for RED-style HTTP signals and database pool health, scrapable without a vendor APM.
3. **Orchestration probes** distinct from the product **`GET /health`** used by web/mobile clients.

Phase 0 scaffolds reserved routes and a JSON `slog` baseline in `apps/api`. Full behavior is deferred so the skeleton stays thin, but the contract must be fixed now.

## Decision

| Concern | Choice |
|---------|--------|
| Logging | Stdlib `log/slog` with **JSON** handler to **stdout** |
| Metrics | **Prometheus** text exposition at **`GET /metrics`** (`prometheus/client_golang` when implemented) |
| Liveness | **`GET /livez`** — process is up; **no** dependency checks |
| Readiness | **`GET /readyz`** — process can take traffic; **includes** Postgres ping (and later critical deps) |
| Client health | **`GET /health`** — product-facing; may include DB status for app shells; **not** a kube probe target |

Non-goals for the first implementation pass:

- No required vendor APM (Datadog, New Relic, etc.).
- No OpenTelemetry export requirement (may be revisited later).
- No separate metrics listen address unless scrape isolation becomes necessary; default is same HTTP server as the API.

## Consequences

- Stubs today return `501` with ticket/ADR pointers so reserved paths are not `404`.
- Adding Prometheus is an explicit dependency change under E12-F5-T2.
- Deploy docs (`deploy/`, Compose, K8s) must use `/livez` and `/readyz`, never stub responses, once T3 lands.
- Log volume stays manageable by defaulting to `INFO` and avoiding PII (DOB, tokens) in fields.

## Implementation instructions

Follow these when executing **E12-F5-T1**, **E12-F5-T2**, and **E12-F5-T3**. Replace packages under `apps/api/internal/logging`, `internal/metrics`, and `internal/health` rather than inventing parallel paths.

### Environment

| Variable | Default | Purpose |
|----------|---------|---------|
| `LOG_LEVEL` | `INFO` | `DEBUG` \| `INFO` \| `WARN` \| `ERROR` for slog |
| `PORT` | `8080` | API + probes + `/metrics` (same listener) |
| `DATABASE_URL` | (required) | Used by `/readyz` and `/health` |

Optional later: `OTEL_*` only if a follow-up ADR adds tracing.

### Logging (E12-F5-T1)

1. Honor `LOG_LEVEL` in `internal/logging.Setup`.
2. Emit JSON to stdout only (no file sinks in-process).
3. HTTP middleware: method, path, status, duration_ms, and a **request/correlation id** (generate UUID if incoming `X-Request-ID` absent; echo it on the response).
4. Never log `JWT_SECRET`, raw `Authorization`, or swimmer DOB/guardian PII.
5. Keep startup/shutdown lines at Info.

**Acceptance:** running API under Docker shows JSON lines; a request with `X-Request-ID: abc` appears as field `request_id=abc` on the access log line.

### Metrics (E12-F5-T2)

1. Add `github.com/prometheus/client_golang`.
2. Replace the `/metrics` stub with `promhttp.Handler()` (or equivalent).
3. Minimum series:
   - HTTP: request count and latency histogram by method + route template + status class.
   - Process: default Go collectors are acceptable.
   - DB: pgx pool gauges (acquired / idle / max) if readily available; otherwise document follow-up.
4. Metric namespaced with prefix `openswim_api_` where custom.
5. Document scrape in `deploy/` or ADR appendix:

```yaml
# Example Prometheus scrape fragment
scrape_configs:
  - job_name: openswim-api
    metrics_path: /metrics
    static_configs:
      - targets: ["api:8080"]
```

**Acceptance:** `curl -s localhost:8080/metrics` returns Prometheus text format including at least one HTTP counter after traffic; no `501` on `/metrics`.

### Probes (E12-F5-T3)

| Path | Success | Failure | Checks |
|------|---------|---------|--------|
| `GET /livez` | `200` `{"status":"ok"}` | process dead → no response | none |
| `GET /readyz` | `200` `{"status":"ok"}` | `503` `{"status":"not_ready"}` | Postgres `Ping` (short timeout) |

1. Remove stub `501` handlers in `internal/health`.
2. Do **not** change `/health` product semantics without an OpenAPI/`E1-F3` update.
3. Production Dockerfile / Compose / K8s examples:

```dockerfile
# Only after this task ships — not while stubs return 501
HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD ["/api", "healthcheck", "livez"]
```

If the binary has no subcommand helper, use a distro with `wget`/`curl` in a non-distroless variant for HEALTHCHECK, or rely on orchestrator HTTP probes only:

```yaml
livenessProbe:
  httpGet:
    path: /livez
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 10
readinessProbe:
  httpGet:
    path: /readyz
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 5
```

**Acceptance:** with Postgres up, `/livez` and `/readyz` return 200; with Postgres stopped, `/livez` stays 200 and `/readyz` returns 503; stubs gone.

### Scaffold map (current tree)

| Path | Role until E12-F5 |
|------|-------------------|
| `internal/logging` | JSON slog baseline |
| `internal/metrics` | `/metrics` → 501 + ticket pointer |
| `internal/health` | `/livez`, `/readyz` → 501 + ticket pointer |
| `GET /health` | Implemented client health (DB ping) |

### Related backlog

- `E12-F5-T1` Structured logging
- `E12-F5-T2` Prometheus metrics
- `E12-F5-T3` Container probes
- `E12-F5-T4` DB backup docs (out of scope for this ADR’s runtime surface)
