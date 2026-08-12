# OpenSwim

**OpenSwim** is a free, ad-free platform for small-town swim teams. It combines team operations, paperless meet timing, live parent results, ribbon labels, and manager↔parent chat in one open stack—without subscriptions, ads, or artificial feature gates.

## The idea

Local summer leagues and rec teams often juggle several paid or limited apps (TeamSnap for roster/chat, Swimmingly for meet timing, SwimminglyFan for live results). Towns either live with missing features, pay for multiple products, or both.

OpenSwim’s goal is the **combined functionality** of those tools as a single free project that towns can use and contributors can improve:

- **Coaches / managers** — orgs, seasons, rosters, meet setup, scoring
- **Meet volunteers** — starter, lane timers, scorekeeper (iPad-friendly), judges
- **Parents / fans** — heat sheets, live results, rankings, team chat

## What it will do

- Register teams and swimmers; track seasons across years
- Run dual (and later multi-team) meets with phone-based timing synced to a scorekeeper
- Publish live results and rankings; export ribbon labels for printing
- Keep managers and parents aligned with team chat and (later) schedules / volunteers

See the [roadmap](docs/roadmap.md) and [backlog](docs/backlog/epics.md) for phased delivery. Current MVP aim: one paperless dual meet with live parent results and basic chat.

## Architecture

| Layer | Stack |
|-------|--------|
| API | Go + Postgres + WebSockets |
| Web | Next.js (App Router) TypeScript |
| iOS / iPad | Native SwiftUI |
| Android | Native Kotlin Compose |
| Contracts | OpenAPI in `packages/contracts` |

Native mobile is intentional: meet-day timers need reliable clock sync and foreground timing behavior.

## Status

**Phase 0 — scaffold.** Documentation, monorepo layout, and agent metadata are in place. Application runtimes (Go module, Next.js, Xcode, Android Gradle) are **not** bootstrapped yet. You can contribute to docs and structure today; local “run the app” steps arrive after Phase 0 boot.

## Get started

**New contributors:** follow the **[local onboarding guide](docs/onboarding.md)** (prerequisites, tools, and how to orient in the repo).

Also useful:

- [Contributing](CONTRIBUTING.md) · [Contributors](CONTRIBUTORS.md) · [Code of Conduct](CODE_OF_CONDUCT.md) · [Security](SECURITY.md)
- [Roadmap](docs/roadmap.md) · [Domain model](docs/domain-model.md) · [Competitor mapping](docs/competitors.md)
- [Epics](docs/backlog/epics.md) · [Features](docs/backlog/features.md) · [Tasks](docs/backlog/tasks.md)

## Repository layout

```
apps/api              Go API (scaffold)
apps/web              Next.js web (scaffold)
apps/ios              SwiftUI (scaffold)
apps/android          Compose (scaffold)
packages/contracts    Shared OpenAPI (scaffold)
packages/design-tokens
docs/                 Roadmap, backlog, onboarding, agent metadata
skills/               Portable Agent Skills (harness-agnostic)
.github/              CODEOWNERS, PR template, Copilot adapter (no CI workflows yet)
scripts/              Future codegen / helpers
deploy/               Future self-host configs
```

## For coding agents

Harness-agnostic entrypoint: **[AGENTS.md](AGENTS.md)**. Details: [docs/agents/](docs/agents/). Skills: [skills/](skills/). Adapters (`CLAUDE.md`, `.cursor/rules/`, etc.) are pointers only—do not duplicate project rules into them.

## License

MIT — see [LICENSE](LICENSE).
