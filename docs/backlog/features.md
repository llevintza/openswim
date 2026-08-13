# Features

Feature catalog by epic. IDs are stable references for tasks (`{Epic}-{Feature}`).

Status legend: `planned` | `mvp` | `later` | `done`

---

## E1 — Platform & Auth

| ID | Feature | Phase | Status | Description |
|----|---------|-------|--------|-------------|
| E1-F1 | Repo & docs scaffold | 0 | mvp | Monorepo layout, license, roadmap, backlog |
| E1-F2 | Go API skeleton | 0 | done | Module, config, `/health`, Postgres connection |
| E1-F3 | OpenAPI contracts package | 0 | planned | Shared OpenAPI specs under `packages/contracts` |
| E1-F4 | Email + password auth | 1 | mvp | Register, login, logout, session/JWT |
| E1-F5 | Magic link auth | 1 | mvp | Passwordless email login |
| E1-F6 | Role-based access control | 1 | mvp | Org/team roles enforced on API |
| E1-F7 | Meet device sessions | 1 | mvp | Temporary starter/timer/judge/scorekeeper grants |
| E1-F8 | Web app shell | 0–1 | planned | Next.js app, auth pages, nav shell |
| E1-F9 | iOS app shell | 0–1 | planned | SwiftUI project, auth, deep links |
| E1-F10 | Android app shell | 0–1 | planned | Compose project, auth, deep links |
| E1-F11 | WebSocket hub | 1 | mvp | Meet + chat realtime channels |

---

## E2 — Organization & Seasons

| ID | Feature | Phase | Status | Description |
|----|---------|-------|--------|-------------|
| E2-F1 | Create organization | 1 | mvp | Name, locale, timezone |
| E2-F2 | Teams under org | 1 | mvp | Create/list teams |
| E2-F3 | Seasons | 1 | mvp | Season name, dates, age cutoff |
| E2-F4 | Invite codes | 1 | mvp | Coach/parent invite links with role |
| E2-F5 | Org settings | 1 | mvp | Basic profile, contact, default pool |
| E2-F6 | Member directory | 1 | mvp | List members by role |

---

## E3 — Roster & Guardians

| ID | Feature | Phase | Status | Description |
|----|---------|-------|--------|-------------|
| E3-F1 | Swimmer CRUD | 1 | mvp | Name, DOB, gender, notes |
| E3-F2 | Season roster | 1 | mvp | Add/remove swimmers for a season/team |
| E3-F3 | Age group computation | 1 | mvp | Age from DOB + season cutoff |
| E3-F4 | Guardian links | 1 | mvp | Link parent users to swimmers |
| E3-F5 | Roster import CSV | 2 | later | Bulk import for season start |
| E3-F6 | Swimmer eligibility filters | 1 | mvp | Filter by age/gender for event entry UI |

---

## E4 — Meet Setup & Entries

| ID | Feature | Phase | Status | Description |
|----|---------|-------|--------|-------------|
| E4-F1 | Meet CRUD | 1 | mvp | Dual meet: host, opponent, date, venue, course |
| E4-F2 | Dual meet event template | 1 | mvp | Standard event list (age × stroke × distance) |
| E4-F3 | Meet entries | 1 | mvp | Enter swimmers into events with seed times |
| E4-F4 | Simple seeding | 1 | mvp | Seed by time → heats/lanes |
| E4-F5 | Heat sheet views | 1 | mvp | Web + mobile heat sheets |
| E4-F6 | Heat sheet PDF | 1 | mvp | Printable heat sheet |
| E4-F7 | Day-of entry/heat edits | 2 | later | Scratch, add, reseed mid-meet |
| E4-F8 | Relays | 2 | later | Relay entries and legs |
| E4-F9 | Exhibition swims | 2 | later | Non-scoring entries |
| E4-F10 | Pool / lane config | 2 | later | Lane count, course per meet |
| E4-F11 | Psych sheets | 2 | later | Unseeded entry lists |

---

## E5 — Live Timing & Scorekeeping

| ID | Feature | Phase | Status | Description |
|----|---------|-------|--------|-------------|
| E5-F1 | Meet huddle QR join | 1 | mvp | Devices scan QR to join with role |
| E5-F2 | Starter start pulse | 1 | mvp | Broadcast official start to timers |
| E5-F3 | Lane timer UI | 1 | mvp | Native stopwatch tied to start pulse |
| E5-F4 | Scorekeeper verify UI | 1 | mvp | iPad-friendly accept/edit times |
| E5-F5 | Heat advance controls | 1 | mvp | Scorekeeper/starter advance heat/event |
| E5-F6 | Multi-timer per lane | 2 | later | Aggregate official time rules |
| E5-F7 | Stroke & turn judge | 2 | later | DQ codes synced to scorekeeper |
| E5-F8 | Offline LAN meet mode | 2 | later | Local hub; cloud sync after |
| E5-F9 | Device nicknames & roster | 1 | mvp | Scorekeeper sees connected devices |

---

## E6 — Results, Rankings & Ribbons

| ID | Feature | Phase | Status | Description |
|----|---------|-------|--------|-------------|
| E6-F1 | Place calculation | 1 | mvp | Places from official times (DQ last) |
| E6-F2 | Dual meet team scoring | 1 | mvp | Configurable point table |
| E6-F3 | Live rankings | 1 | mvp | Event results + running team score |
| E6-F4 | Ribbon label PDF | 1 | mvp | Place ribbons for printing |
| E6-F5 | PB / time-improvement labels | 2 | later | Labels when beating prior best |
| E6-F6 | Participation labels | 2 | later | Optional labels for all swimmers |
| E6-F7 | Meet results export | 2 | later | PDF/CSV results packet |
| E6-F8 | Official vs unofficial toggle | 2 | later | Scorekeeper publish control |

---

## E7 — Parent/Fan Live Experience

| ID | Feature | Phase | Status | Description |
|----|---------|-------|--------|-------------|
| E7-F1 | Live heat sheet (fan) | 1 | mvp | Read-only heat sheet during meet |
| E7-F2 | Live results (fan) | 1 | mvp | Results as scorekeeper verifies |
| E7-F3 | Live team scores (fan) | 1 | mvp | Dual meet scoreboard view |
| E7-F4 | Meet progress bar | 2 | later | Current event/heat indicator |
| E7-F5 | Follow swimmers | 4 | later | Personalized entries/results |
| E7-F6 | Race-soon push alerts | 2 | later | Notify N events before swim |
| E7-F7 | Public meet web page | 4 | later | Shareable link without full account |
| E7-F8 | Time improvement on results | 2 | later | Delta vs seed/PB on fan UI |

---

## E8 — Messaging & Notifications

| ID | Feature | Phase | Status | Description |
|----|---------|-------|--------|-------------|
| E8-F1 | Team channel chat | 1 | mvp | Manager/coach ↔ parents channel |
| E8-F2 | Unread indicators | 1 | mvp | Per-conversation unread |
| E8-F3 | In-app notifications | 1 | mvp | Meet + chat events in-app |
| E8-F4 | Push notification plumbing | 1 | mvp | Device token registration hooks |
| E8-F5 | Announcements | 3 | later | Org-wide broadcast posts |
| E8-F6 | Polls | 3 | later | Simple team polls |
| E8-F7 | Direct messages | 3 | later | Manager ↔ parent DM |

---

## E9 — Scheduling & Volunteers

| ID | Feature | Phase | Status | Description |
|----|---------|-------|--------|-------------|
| E9-F1 | Practice/meet calendar | 3 | later | ScheduleEvent CRUD |
| E9-F2 | RSVP / availability | 3 | later | Going / not / maybe |
| E9-F3 | Volunteer slots | 3 | later | Assign timers, ribbons, etc. |
| E9-F4 | External calendar feed | 3 | later | iCal per team |
| E9-F5 | Photo sharing (light) | 3 | later | Album attached to events |

---

## E10 — Swimmer Development History

| ID | Feature | Phase | Status | Description |
|----|---------|-------|--------|-------------|
| E10-F1 | Career results timeline | 4 | later | All results across seasons |
| E10-F2 | Personal bests by course | 4 | later | Materialized PBs |
| E10-F3 | Year-over-year charts | 4 | later | Simple progress views |
| E10-F4 | Time standards badges | 4 | later | Tag results vs standards |

---

## E11 — League & Multi-team Meets

| ID | Feature | Phase | Status | Description |
|----|---------|-------|--------|-------------|
| E11-F1 | Multi-team meet | 3–4 | later | 3+ teams at one meet |
| E11-F2 | League standings | 3 | later | Aggregate dual/multi scores |
| E11-F3 | Championship format | 4 | later | Prelims/finals or scored finals |
| E11-F4 | Records board | 4 | later | Pool/team/league records |

---

## E12 — Interop, Export & Self-host

| ID | Feature | Phase | Status | Description |
|----|---------|-------|--------|-------------|
| E12-F1 | Results CSV/PDF pack | 5 | later | Portable meet archive |
| E12-F2 | Hy-Tek / SDIF / CL2 | 5 | later | Import/export where useful |
| E12-F3 | Scoreboard display mode | 5 | later | Full-screen TV from scorekeeper |
| E12-F4 | Docker Compose deploy | 5 | later | One-command self-host |
| E12-F5 | Observability & backups | 5 | later | Structured logs, Prometheus metrics, `/livez`/`/readyz`, DB backup docs (see ADR 0002) |
| E12-F6 | Offline fee tracking | 5 | later | Mark registration paid; optional Stripe |

---

## MVP feature set (Phase 1)

Must ship: `E1-F4`–`E1-F7`, `E1-F11`, `E2-F1`–`E2-F6`, `E3-F1`–`E3-F4`, `E3-F6`, `E4-F1`–`E4-F6`, `E5-F1`–`E5-F5`, `E5-F9`, `E6-F1`–`E6-F4`, `E7-F1`–`E7-F3`, `E8-F1`–`E8-F4`.
