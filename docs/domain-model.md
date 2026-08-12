# Domain Model

Sketch of core entities for OpenSwim. Names are logical; physical schema lands with Phase 0 API work.

## Organizations & seasons

| Entity | Purpose |
|--------|---------|
| **Organization** | Town/club that owns teams (e.g. “Riverside Rec Swim”). |
| **Team** | Competitive unit within an org (often one per age program). |
| **Season** | Named period with age cutoff date (e.g. Summer 2026). |
| **PoolConfig** | Course (SCY/SCM/LCM), lane count, timing notes for a venue. |

## People & access

| Entity | Purpose |
|--------|---------|
| **User** | Login identity (email). |
| **Membership** | User ↔ org/team with **Role**. |
| **Role** | `org_admin`, `coach` / `manager`, `parent`, `swimmer` (account), meet roles: `starter`, `timer`, `judge`, `scorekeeper`. |
| **Swimmer** | Athlete profile: name, DOB, gender, preferred events; survives across seasons. |
| **SeasonRosterEntry** | Swimmer on a team for a season (age group computed from cutoff). |
| **GuardianLink** | Parent/guardian User ↔ Swimmer. |

Meet roles may be temporary **TimingDeviceSession** grants for a specific meet (QR join), not permanent memberships.

## Meet structure

| Entity | Purpose |
|--------|---------|
| **EventDefinition** | Stroke + distance + age group + gender (+ relay flag); reusable templates. |
| **Meet** | Scheduled competition: host team, opponent(s), venue, date, course, status. |
| **MeetEvent** | Instance of an EventDefinition in a meet (order, scoring rules). |
| **MeetEntry** | Swimmer (or relay) entered in a MeetEvent with seed time. |
| **Heat** | Numbered heat within a MeetEvent. |
| **LaneAssignment** | Swimmer/relay in a heat + lane. |
| **RelayLeg** | Ordered swimmers on a relay entry. |

## Timing & results

| Entity | Purpose |
|--------|---------|
| **TimingDeviceSession** | Device joined to a meet huddle with role + lane (for timers). |
| **StartPulse** | Starter’s official start timestamp (and optional reaction) for a heat. |
| **LaneTime** | Raw stop from a timer device (lane, heat, timestamp, timer index). |
| **OfficialTime** | Computed/accepted time after scorekeeper rules (median, pad, etc.). |
| **Result** | Final place, time, DQ flag, exhibition flag for an entry. |
| **TeamScore** | Running / final dual (or multi) meet points by team. |
| **RibbonJob** | Request to generate place / PB / participation labels (PDF). |

## Communication & schedule

| Entity | Purpose |
|--------|---------|
| **Conversation** | Team channel or DM thread. |
| **Message** | Chat message with author, body, timestamps. |
| **ScheduleEvent** | Practice, meet, or other calendar item. |
| **RSVP** | Availability response for a ScheduleEvent. |
| **VolunteerSlot** | Named duty (timer lane 3, ribbons, etc.) with assignee. |

## Fan / history (later phases)

| Entity | Purpose |
|--------|---------|
| **Follow** | User follows a Swimmer for alerts and personalized views. |
| **PersonalBest** | Best Result per stroke/distance/course (materialized). |
| **Record** | Pool / team / league record holders. |
| **TimeStandard** | Named cut times for ranking badges. |

## Key relationships

```text
Organization
  └── Team(s)
        └── SeasonRosterEntry(s) ── Swimmer ── GuardianLink ── User
  └── Season(s)
  └── Meet(s)
        ├── MeetEvent(s) ── EventDefinition
        │     ├── MeetEntry(s)
        │     └── Heat(s) ── LaneAssignment(s)
        ├── TimingDeviceSession(s)
        ├── Result(s) / TeamScore
        └── RibbonJob(s)
```

## Age groups

Age for a season is derived from **Swimmer.DOB** and **Season.ageCutoffDate** (league-configurable). Event eligibility filters use that computed age group.
