# Constraints

Hard rules for OpenSwim. Do not violate unless the user explicitly overrides them in chat.

## Product

- **Free forever** for local teams: no ads, no feature paywalls, no upsell UI.
- Target users: small-town / summer-league swim teams, coaches, parents, meet volunteers.
- Combined scope: TeamSnap-like team ops + Swimmingly-like meet timing + SwimminglyFan-like live results.
- **Non-goals (v1–v2):** live streaming, paid drills libraries, website builder, card payments (offline “mark paid” only until a later optional phase).

## Architecture

- API: **Go + Postgres**; realtime via **WebSockets** from the Go service.
- Web: **Next.js** App Router (TypeScript).
- Mobile: **native SwiftUI** (iOS/iPad) and **native Kotlin Compose** (Android).
- Shared contracts: OpenAPI under `packages/contracts`.
- Meet-day roles (starter, timer, judge, scorekeeper) **must remain native**. Do not migrate them to React Native, Flutter, Expo, or WebView shells.

## Scope discipline

- Follow [`docs/roadmap.md`](../roadmap.md) phases and [`docs/backlog/`](../backlog/) IDs.
- Phase 0 scaffold is docs + empty dirs; **do not** bootstrap app frameworks unless asked.
- Prefer updating backlog docs when product scope changes; do not invent parallel tracking systems.
- Keep PRs focused; avoid drive-by refactors unrelated to the task.

## Security / privacy

- No secrets in git (`.env`, keys, credentials).
- Treat swimmer DOB and guardian data as sensitive PII in designs and logs.
- Temporary meet device sessions are not permanent org memberships.

## Licensing

- MIT ([`LICENSE`](../../LICENSE)). Do not add conflicting license files without an explicit decision.
