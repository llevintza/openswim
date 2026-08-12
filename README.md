# OpenSwim

Free, ad-free swim team platform for small-town teams: roster and season ops, paperless meet timing, live parent results, ribbon labels, and team chat — without subscriptions or ads.

## Why

Apps like TeamSnap, Swimmingly, and SwimminglyFan cover pieces of the problem but often gate features behind paywalls, ads, or limits. OpenSwim aims for the combined functionality as an open, free stack.

## Architecture

| Layer | Stack |
|-------|--------|
| API | Go + Postgres + WebSockets |
| Web | Next.js (App Router) TypeScript |
| iOS / iPad | Native SwiftUI |
| Android | Native Kotlin Compose |
| Contracts | OpenAPI in `packages/contracts` |

Native mobile is intentional: meet-day timers need reliable clock sync and foreground timing behavior.

## Repository layout

```
apps/api          Go API (scaffold)
apps/web          Next.js web (scaffold)
apps/ios          SwiftUI (scaffold)
apps/android      Compose (scaffold)
packages/contracts
packages/design-tokens
docs/             Roadmap, domain model, backlog, agent metadata
skills/           Portable Agent Skills (harness-agnostic)
scripts/          Future codegen / helpers
deploy/           Future self-host configs
```

## Docs

- [Roadmap](docs/roadmap.md) — phases 0–5
- [Domain model](docs/domain-model.md)
- [Competitor mapping](docs/competitors.md)
- [Epics](docs/backlog/epics.md) · [Features](docs/backlog/features.md) · [Tasks](docs/backlog/tasks.md)

## For coding agents

Harness-agnostic entrypoint: **[AGENTS.md](AGENTS.md)**. Details and anti-duplication policy live in [docs/agents/](docs/agents/) ([harness map](docs/agents/harness-map.md), [known issues](docs/agents/known-issues.md)). Portable skills: [skills/](skills/). Thin adapters only: `CLAUDE.md`, `.cursor/rules/`, `.github/copilot-instructions.md`, `.gemini/settings.json`.

## Current status

Phase 0 scaffold: documentation and empty app directories only. No runtimes bootstrapped yet.

## License

MIT — see [LICENSE](LICENSE).
