# Local onboarding

This guide gets contributors ready to work on OpenSwim **locally**. Prefer **Dev Containers** so language toolchains (Go, later Node, etc.) stay inside Docker — not on the host.

## Goals

1. Clone the repo and understand purpose, roadmap, and backlog.
2. Open the Dev Container for API (and future shared) work.
3. Follow [CONTRIBUTING.md](../CONTRIBUTING.md) for branches and pull requests.

## Host prerequisites

| Tool | Notes |
|------|--------|
| Git | Clone and PR workflow |
| GitHub account | Forks / PRs against [openswim](https://github.com/llevintza/openswim) |
| Editor | Cursor or VS Code with the **Dev Containers** feature/extension |
| Docker Desktop (or compatible engine) | Required to run [`.devcontainer/`](../.devcontainer/) |

Do **not** install Go, Node, or Python on the host for OpenSwim day-to-day work. Those run inside the container (or native IDE toolchains for iOS/Android only).

### Optional: coding agents

If you use Cursor, Claude Code, Copilot CLI, Codex, or similar, read [AGENTS.md](../AGENTS.md) and [docs/agents/](agents/). Portable skills live under [`skills/`](../skills/). Assume shell commands for the API run **inside** the Dev Container unless noted.

## Tooling versions (inside the Dev Container)

| Tool | Target |
|------|--------|
| Go | **1.22+** (image: `devcontainers/go:1.22`) |
| Postgres | **16** (Compose service `postgres`) |

Web/iOS/Android toolchains will be added to the Dev Container (or profiles) when those apps boot. Until then:

| Surface | Notes |
|---------|--------|
| Web (`apps/web`) | Not bootstrapped; Node **22 LTS** preferred when added to the container |
| iOS (`apps/ios`) | Still needs macOS + Xcode **16+** on the host for Simulator/device |
| Android (`apps/android`) | Android Studio / SDK on the host until a container story exists |

Meet timing stays **native** SwiftUI / Kotlin Compose.

## Clone and open

```bash
git clone https://github.com/llevintza/openswim.git
cd openswim
```

In Cursor/VS Code: **Dev Containers: Reopen in Container**. Compose starts `dev` (Go) and `postgres`.

Suggested reading order:

1. [README.md](../README.md) — purpose and idea  
2. This onboarding guide  
3. [docs/roadmap.md](roadmap.md) — phases  
4. [docs/backlog/](backlog/) — epics, features, tasks  
5. [docs/domain-model.md](domain-model.md) — entity language  
6. [docs/adr/](adr/) — architecture decisions  
7. [CONTRIBUTING.md](../CONTRIBUTING.md) — how to send changes  

## Run the API

Inside the Dev Container:

```bash
cd apps/api
go run ./cmd/api
curl -s localhost:8080/health
```

Details: [apps/api/README.md](../apps/api/README.md).

## What works today

| Activity | Status |
|----------|--------|
| Read / edit docs and backlog | Yes |
| Agent metadata (`AGENTS.md`, `docs/agents/`, `skills/`) | Yes |
| Dev Container (Go + Postgres) | Yes |
| `go run ./cmd/api` + `GET /health` | Yes (in Dev Container) |
| `/livez`, `/readyz`, `/metrics` | Stubbed — see [ADR 0002](adr/0002-api-observability.md) |
| `npm run dev` / Xcode / Gradle | Not yet |

## Contribution path

1. Create a feature branch from `main` (e.g. `feature/short-description`).
2. Make focused changes; do not commit secrets (`.env`, keys, credentials).
3. Open a PR against `main` using the GitHub PR template.
4. Expect review from CODEOWNERS (`@llevintza`).

Details: [CONTRIBUTING.md](../CONTRIBUTING.md).

## Troubleshooting

- **Cannot push `.github/workflows/`** — current GitHub OAuth credentials may lack the `workflow` scope. See [docs/agents/known-issues.md](agents/known-issues.md).
- **Dev Container build fails** — confirm Docker is running; rebuild the container from the command palette.
- **API cannot reach Postgres** — use host `postgres` in `DATABASE_URL` (Compose service name), not `localhost`, from inside `dev`.

## Related docs

- [Architecture decisions](adr/)  
- [Competitor mapping](competitors.md)  
- [Security policy](../SECURITY.md)  
- [Code of Conduct](../CODE_OF_CONDUCT.md)  
