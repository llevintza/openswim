# Tasks

Detailed checklist for **Phase 0** and **Phase 1**. Later-phase work is catalogued as features in [features.md](./features.md); break into tasks when those phases start.

Task IDs: `{Feature}-T{n}` (e.g. `E5-F2-T1`).

---

## Phase 0 — Foundations

### E1-F1 Repo & docs scaffold

- [x] E1-F1-T1 Create monorepo directory tree (`apps/`, `packages/`, `docs/`, `scripts/`, `deploy/`, `.github/`)
- [x] E1-F1-T2 Add root `README.md`, `LICENSE` (MIT), `.gitignore`
- [x] E1-F1-T3 Write `docs/roadmap.md`, `docs/domain-model.md`, `docs/competitors.md`
- [x] E1-F1-T4 Write `docs/backlog/epics.md`, `features.md`, `tasks.md`
- [x] E1-F1-T5 Add stub READMEs + `.gitkeep` under app/package dirs

### E1-F2 Go API skeleton *(post-scaffold)*

- [x] E1-F2-T1 Initialize Go module under `apps/api`
- [x] E1-F2-T2 Config via env (`DATABASE_URL`, `PORT`, `JWT_SECRET`)
- [x] E1-F2-T3 HTTP server with `GET /health`
- [x] E1-F2-T4 Postgres connection + migration runner stub
- [x] E1-F2-T5 Dockerfile stub for API

### E1-F3 OpenAPI contracts

- [ ] E1-F3-T1 Add OpenAPI 3 skeleton in `packages/contracts`
- [ ] E1-F3-T2 Document `/health` and auth placeholder paths
- [ ] E1-F3-T3 Script stub to validate/lint OpenAPI (`scripts/`)

### E1-F8 / E1-F9 / E1-F10 Client shells *(post-scaffold)*

- [ ] E1-F8-T1 Bootstrap Next.js App Router in `apps/web`
- [ ] E1-F8-T2 Healthcheck page calling API `/health`
- [ ] E1-F9-T1 Create Xcode SwiftUI project in `apps/ios`
- [ ] E1-F9-T2 Healthcheck screen calling API
- [ ] E1-F10-T1 Create Android Compose project in `apps/android`
- [ ] E1-F10-T2 Healthcheck screen calling API

---

## Phase 1 — MVP thin slice

### E1-F4 Email + password auth

- [ ] E1-F4-T1 `users` table migration
- [ ] E1-F4-T2 Register endpoint (email, password hash)
- [ ] E1-F4-T3 Login endpoint → session/JWT
- [ ] E1-F4-T4 Logout / token revoke
- [ ] E1-F4-T5 Web login/register forms
- [ ] E1-F4-T6 iOS login/register screens
- [ ] E1-F4-T7 Android login/register screens

### E1-F5 Magic link auth

- [ ] E1-F5-T1 Generate one-time token + email send stub
- [ ] E1-F5-T2 Consume magic link → session
- [ ] E1-F5-T3 Web “email me a link” flow

### E1-F6 Role-based access control

- [ ] E1-F6-T1 `memberships` table (user, org/team, role)
- [ ] E1-F6-T2 Middleware: require auth + role checks
- [ ] E1-F6-T3 Seed roles: `org_admin`, `coach`, `parent`

### E1-F7 Meet device sessions

- [ ] E1-F7-T1 `timing_device_sessions` table
- [ ] E1-F7-T2 Issue short-lived meet join token (role + optional lane)
- [ ] E1-F7-T3 API to list/revoke active devices for a meet

### E1-F11 WebSocket hub

- [ ] E1-F11-T1 WS endpoint with auth (user or device session)
- [ ] E1-F11-T2 Channel model: `meet:{id}`, `team:{id}:chat`
- [ ] E1-F11-T3 Heartbeat + reconnect guidance in contracts

### E2-F1 Create organization

- [ ] E2-F1-T1 `organizations` migration
- [ ] E2-F1-T2 `POST /orgs` (creator becomes `org_admin`)
- [ ] E2-F1-T3 Web “Create team/org” flow

### E2-F2 Teams under org

- [ ] E2-F2-T1 `teams` migration
- [ ] E2-F2-T2 CRUD endpoints scoped to org
- [ ] E2-F2-T3 Web team list/create

### E2-F3 Seasons

- [ ] E2-F3-T1 `seasons` migration (name, start, end, age_cutoff)
- [ ] E2-F3-T2 CRUD + set active season
- [ ] E2-F3-T3 Web season picker

### E2-F4 Invite codes

- [ ] E2-F4-T1 `invites` table (code, role, expiry, max uses)
- [ ] E2-F4-T2 Redeem invite → membership
- [ ] E2-F4-T3 Web/mobile redeem screen

### E2-F5 Org settings

- [ ] E2-F5-T1 Patch org profile (name, contact, timezone)
- [ ] E2-F5-T2 Default venue/pool fields (basic)

### E2-F6 Member directory

- [ ] E2-F6-T1 List members with roles
- [ ] E2-F6-T2 Remove / change role (admin only)

### E3-F1 Swimmer CRUD

- [ ] E3-F1-T1 `swimmers` migration
- [ ] E3-F1-T2 CRUD API
- [ ] E3-F1-T3 Web swimmer form
- [ ] E3-F1-T4 Mobile swimmer list (coach)

### E3-F2 Season roster

- [ ] E3-F2-T1 `season_roster_entries` migration
- [ ] E3-F2-T2 Add/remove swimmer on season/team
- [ ] E3-F2-T3 Roster UI filtered by season

### E3-F3 Age group computation

- [ ] E3-F3-T1 Shared age-group helper (cutoff + DOB)
- [ ] E3-F3-T2 Expose computed age group on roster API
- [ ] E3-F3-T3 Unit tests for cutoff edge cases

### E3-F4 Guardian links

- [ ] E3-F4-T1 `guardian_links` migration
- [ ] E3-F4-T2 Link/unlink parent ↔ swimmer
- [ ] E3-F4-T3 Parent “my swimmers” API

### E3-F6 Swimmer eligibility filters

- [ ] E3-F6-T1 Filter roster by age/gender for entry picker
- [ ] E3-F6-T2 Wire into meet entry UI

### E4-F1 Meet CRUD

- [ ] E4-F1-T1 `meets` migration (host, opponent, datetime, venue, course, status)
- [ ] E4-F1-T2 CRUD API
- [ ] E4-F1-T3 Web create dual meet form

### E4-F2 Dual meet event template

- [ ] E4-F2-T1 `event_definitions` + seed standard summer dual template
- [ ] E4-F2-T2 Instantiate `meet_events` from template on meet create
- [ ] E4-F2-T3 Allow reorder / enable-disable events

### E4-F3 Meet entries

- [ ] E4-F3-T1 `meet_entries` migration (seed time nullable)
- [ ] E4-F3-T2 Enter/withdraw swimmer for meet event
- [ ] E4-F3-T3 Entry UI for coaches (both teams if invited)

### E4-F4 Simple seeding

- [ ] E4-F4-T1 Seed algorithm: sort by seed time, assign heats/lanes
- [ ] E4-F4-T2 Persist `heats` + `lane_assignments`
- [ ] E4-F4-T3 Re-seed before meet start (replace assignments)

### E4-F5 Heat sheet views

- [ ] E4-F5-T1 API: heat sheet by meet (events → heats → lanes)
- [ ] E4-F5-T2 Web heat sheet page
- [ ] E4-F5-T3 iOS heat sheet screen
- [ ] E4-F5-T4 Android heat sheet screen

### E4-F6 Heat sheet PDF

- [ ] E4-F6-T1 Generate printable PDF from heat sheet data
- [ ] E4-F6-T2 Download from web admin

### E5-F1 Meet huddle QR join

- [ ] E5-F1-T1 Scorekeeper generates join QR (meet id + token)
- [ ] E5-F1-T2 Native scan → select role (starter/timer/…) + lane if timer
- [ ] E5-F1-T3 Register `TimingDeviceSession`; subscribe to meet WS channel

### E5-F2 Starter start pulse

- [ ] E5-F2-T1 Starter UI: current heat, Start button
- [ ] E5-F2-T2 Broadcast `StartPulse` over WS with server timestamp
- [ ] E5-F2-T3 Timers latch start; show running clock

### E5-F3 Lane timer UI

- [ ] E5-F3-T1 Native timer screen (lane, heat, swimmer name)
- [ ] E5-F3-T2 Stop → send `LaneTime` to API/WS
- [ ] E5-F3-T3 Prevent double-submit; allow recall before verify

### E5-F4 Scorekeeper verify UI

- [ ] E5-F4-T1 Heat grid: lanes, raw times, flags
- [ ] E5-F4-T2 Accept / edit / empty / DQ stub flag
- [ ] E5-F4-T3 Persist `OfficialTime` + draft `Result`
- [ ] E5-F4-T4 iPad landscape layout

### E5-F5 Heat advance controls

- [ ] E5-F5-T1 Advance to next heat/event
- [ ] E5-F5-T2 Broadcast current heat to all devices + fans
- [ ] E5-F5-T3 Guard: cannot advance if times pending (override allowed)

### E5-F9 Device nicknames & roster

- [ ] E5-F9-T1 Device sets nickname on join
- [ ] E5-F9-T2 Scorekeeper device list (role, lane, connected)

### E6-F1 Place calculation

- [ ] E6-F1-T1 Compute places from official times within meet event
- [ ] E6-F1-T2 Tie-break rules (document + implement)
- [ ] E6-F1-T3 Persist places on `results`

### E6-F2 Dual meet team scoring

- [ ] E6-F2-T1 Default point table (e.g. 5-3-1 individual; relay table)
- [ ] E6-F2-T2 Recompute `team_scores` when results change
- [ ] E6-F2-T3 Expose running score API

### E6-F3 Live rankings

- [ ] E6-F3-T1 Publish results to meet WS channel on verify
- [ ] E6-F3-T2 Fan/coach live results UI (web + mobile)

### E6-F4 Ribbon label PDF

- [ ] E6-F4-T1 `ribbon_jobs` record + generate place labels PDF
- [ ] E6-F4-T2 Filter by event / place depth
- [ ] E6-F4-T3 Download from scorekeeper or web

### E7-F1 Live heat sheet (fan)

- [ ] E7-F1-T1 Public-or-auth fan meet endpoint for heat sheet
- [ ] E7-F1-T2 Web fan view
- [ ] E7-F1-T3 Native fan heat sheet (parent mode)

### E7-F2 Live results (fan)

- [ ] E7-F2-T1 Subscribe to meet WS results events
- [ ] E7-F2-T2 Render event results as they finalize

### E7-F3 Live team scores (fan)

- [ ] E7-F3-T1 Show running dual score on fan home for meet
- [ ] E7-F3-T2 Update via WS

### E8-F1 Team channel chat

- [ ] E8-F1-T1 `conversations` + `messages` migrations
- [ ] E8-F1-T2 Auto team channel per team
- [ ] E8-F1-T3 Post/list messages API
- [ ] E8-F1-T4 WS broadcast on new message
- [ ] E8-F1-T5 Web chat UI
- [ ] E8-F1-T6 iOS chat UI
- [ ] E8-F1-T7 Android chat UI

### E8-F2 Unread indicators

- [ ] E8-F2-T1 Track last_read per user/conversation
- [ ] E8-F2-T2 Badge counts on clients

### E8-F3 In-app notifications

- [ ] E8-F3-T1 Notification records for chat + meet publish
- [ ] E8-F3-T2 In-app notification list UI

### E8-F4 Push notification plumbing

- [ ] E8-F4-T1 Register device push tokens API
- [ ] E8-F4-T2 Provider stub (APNs/FCM) behind interface
- [ ] E8-F4-T3 Document enablement for Phase 2 race alerts

---

## Phase 2+ tasking

When starting Phase 2, expand these features into tasks in a new section or file:

- E4-F7–F11, E5-F6–F8, E6-F5–F8, E7-F4, E7-F6, E7-F8

Phase 3: E8-F5–F7, E9-\*, E11-F1–F2  
Phase 4: E7-F5, E7-F7, E10-\*, E11-F3–F4  

### Phase 5 — Interop & ops (partial)

Implementation timing is flexible; follow [ADR 0002](../adr/0002-api-observability.md) for T1–T3.

### E12-F5 Observability & backups

- [ ] E12-F5-T1 Structured logging per ADR 0002 (`LOG_LEVEL`, fields, request/correlation IDs)
- [ ] E12-F5-T2 Prometheus `/metrics` per ADR 0002 (HTTP + DB pool; scrape docs)
- [ ] E12-F5-T3 Implement `/livez` + `/readyz`; replace stubs; HEALTHCHECK / probe examples per ADR 0002
- [ ] E12-F5-T4 DB backup docs / runbook

Remaining Phase 5 features to expand when started: E12-F1–F4, E12-F6
