# OpenSwim Roadmap

OpenSwim is a free, ad-free platform for small-town swim teams. It combines team operations (TeamSnap-like), meet timing and scoring (Swimmingly / Hy-Tek-like), and parent/fan live results (SwimminglyFan-like).

## Principles

- Free forever for local teams — no ads, no feature paywalls
- Native mobile for meet-day timing reliability (SwiftUI + Compose)
- Paperless dual meets first; league and interop later
- Self-hostable and community-friendly

## Architecture (locked)

| Layer | Choice |
|-------|--------|
| API | Go + Postgres |
| Realtime | WebSockets from Go (meet sync + chat) |
| Web | Next.js (App Router) TypeScript |
| Mobile | Native SwiftUI (iOS/iPad) + Kotlin Compose (Android) |
| Contracts | OpenAPI under `packages/contracts` |

## Phase 0 — Foundations

**Goal:** Repo, docs, and platform skeleton ready for feature work.

- Monorepo layout, license, glossary, competitor map, backlog
- Post-scaffold: Go module, Postgres migrations, auth stubs, OpenAPI, empty clients wired to healthcheck

## Phase 1 — MVP thin slice (balanced)

**Goal:** A summer-league dual meet runs paperless; parents watch live results.

1. Org signup + invite coaches/parents
2. Season + roster (swimmers, age/gender basics)
3. Dual meet, entries, simple seeding → heat sheet
4. Meet-day roles: starter, lane timers, scorekeeper (iPad-friendly)
5. Sync times → verify → live results + rankings / team score
6. Ribbon label PDF export
7. Basic manager ↔ parent team chat (in-app; push hooks ready)

**Success metric:** One local team completes a dual meet without paper heat sheets or manual time cards.

## Phase 2 — Meet-day parity

**Goal:** Swimmingly-depth meet operations.

- Stroke/turn judge DQs; multi-timer per lane + official time rules
- On-the-fly heat/entry changes; exhibition; relays
- Meet progress bar; “3 events before” race alerts
- Best-time / time-improvement ribbon labels
- Offline meet LAN mode with later cloud sync
- Pool configs (yards/meters, lane count)

## Phase 3 — Team ops parity

**Goal:** TeamSnap-depth season coordination.

- Calendar (practices/meets), RSVP/availability, volunteer assignments
- Announcements, polls, light photo sharing
- Multi-team org / league season standings
- Registration forms + waitlists (offline fee marking only)
- Hardened role permissions

## Phase 4 — Athlete evolution & fan polish

**Goal:** Multi-year tracking and richer spectator experience.

- Swimmer timeline (PBs by stroke/distance/course across years)
- Follow swimmers across meets/seasons
- Public meet pages; shareable heat sheets
- Championship formats, records, time standards

## Phase 5 — Interop & ops maturity

**Goal:** League interop and production operations.

- Hy-Tek / SDIF / CL2 import-export where useful
- Scoreboard display mode (TV/HDMI from scorekeeper)
- Self-host + hosted free tier: observability, backup
- Optional fee tracking / Stripe later (still no ads)

## Non-goals (v1–v2)

- Live streaming
- Paid drills / practice plan libraries
- Website builder
- Card payment processing (deferred)
