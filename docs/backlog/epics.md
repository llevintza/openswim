# Epics

Index of product epics. Features live in [features.md](./features.md); Phase 0–1 tasks in [tasks.md](./tasks.md).

| ID | Epic | Primary phase | Summary |
|----|------|---------------|---------|
| E1 | Platform & Auth | 0–1 | Identity, sessions, roles, API health, contracts |
| E2 | Organization & Seasons | 1 | Orgs, teams, seasons, invites |
| E3 | Roster & Guardians | 1 | Swimmers, roster, parent links, age groups |
| E4 | Meet Setup & Entries | 1–2 | Meets, events, entries, seeding, heat sheets |
| E5 | Live Timing & Scorekeeping | 1–2 | Device huddle, starter, timers, scorekeeper |
| E6 | Results, Rankings & Ribbons | 1–2 | Places, team scores, ribbon PDFs |
| E7 | Parent/Fan Live Experience | 1–4 | Live heat sheets, results, follows, alerts |
| E8 | Messaging & Notifications | 1–3 | Team chat, announcements, push |
| E9 | Scheduling & Volunteers | 3 | Calendar, RSVP, volunteer slots |
| E10 | Swimmer Development History | 4 | Multi-year PBs, timelines, standards |
| E11 | League & Multi-team Meets | 3–4 | Multi-team meets, standings, championships |
| E12 | Interop, Export & Self-host | 5 | Hy-Tek formats, scoreboard, ops |

## Dependency sketch

```text
E1 ──► E2 ──► E3 ──► E4 ──► E5 ──► E6
                      │              │
                      └──────────────┴──► E7
E1 ──► E8
E2 ──► E9
E6 ──► E10
E4 ──► E11
E6 ──► E12
```
