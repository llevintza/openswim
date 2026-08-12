# AGENTS.md

Harness-agnostic instructions for coding agents working on **OpenSwim**.

Detailed rules live under [`docs/agents/`](docs/agents/). Do **not** duplicate them into Cursor/Claude/Copilot-specific files — those are thin pointers only. See [`docs/agents/harness-map.md`](docs/agents/harness-map.md).

## Project

OpenSwim is a free, ad-free swim team platform for small-town teams: org/roster ops, paperless meet timing, live parent results, ribbon labels, and team chat.

## Architecture (locked)

| Layer | Stack |
|-------|--------|
| API | Go + Postgres + WebSockets |
| Web | Next.js (App Router) TypeScript |
| iOS / iPad | Native SwiftUI |
| Android | Native Kotlin Compose |
| Contracts | OpenAPI in `packages/contracts` |

Meet-day timing stays **native**. Do not introduce React Native, Flutter, or Expo for starter/timer/scorekeeper flows.

## Always read first

1. [`docs/agents/constraints.md`](docs/agents/constraints.md) — hard product/tech constraints  
2. [`docs/agents/conventions.md`](docs/agents/conventions.md) — git, docs, monorepo discipline  
3. [`docs/roadmap.md`](docs/roadmap.md) + [`docs/backlog/`](docs/backlog/) — scope and task IDs  
4. Nearest nested `AGENTS.md` when editing under `apps/*`

## Repo map

- `apps/api` — Go API  
- `apps/web` — Next.js  
- `apps/ios` — SwiftUI  
- `apps/android` — Compose  
- `packages/contracts` — OpenAPI  
- `docs/` — roadmap, domain, backlog, agent metadata  
- `skills/` — portable Agent Skills (`*/SKILL.md`)

## Skills

Load project skills from repo-root [`skills/`](skills/) (not harness-private skill dirs):

- `skills/openswim-backlog` — work from epics/features/tasks  
- `skills/openswim-meet-domain` — domain + meet-day concepts  
- `skills/update-known-issues` — append execution problems to agent metadata  

## Current phase

Phase 0 scaffold is in place. Do **not** bootstrap Go modules, Next.js, Xcode, or Android projects unless the user explicitly asks.

## Commits and PRs

- Feature branches; open PRs against `main`  
- No force-push to `main`  
- Commit only when the user asks (this repo’s agent conventions)  
- Prefer HEREDOC commit messages focused on why  

## Known issues

Living log: [`docs/agents/known-issues.md`](docs/agents/known-issues.md). Append any harness/tooling failures encountered during work.
