# Competitor Feature Mapping

How OpenSwim covers capabilities from TeamSnap, Swimmingly, SwimminglyFan, and Hy-Tek-style meet tools. Phase numbers refer to [roadmap.md](./roadmap.md).

## TeamSnap (team ops)

| Competitor capability | OpenSwim approach | Phase |
|----------------------|-------------------|-------|
| Org / club registration | Organization signup, seasons | 1 |
| Rosters & invites | Roster CRUD, invite codes, guardian links | 1 |
| Scheduling & calendar sync | ScheduleEvent; external calendar later | 3 |
| Team chat / alerts | Conversation + Message; push hooks | 1 / 3 |
| RSVP / availability | RSVP on ScheduleEvent | 3 |
| Volunteer assignments | VolunteerSlot | 3 |
| Payments / registration fees | Offline “mark paid”; Stripe optional | 3 / 5 |
| Practice plans / drills library | Out of scope (no paid content library) | — |
| Live streaming / highlights | Out of scope for v1–v2 | — |
| Website builder | Out of scope | — |
| Ads / paid tiers | Never — free, ad-free | — |

## Swimmingly (meet ops + Clubhouse)

| Competitor capability | OpenSwim approach | Phase |
|----------------------|-------------------|-------|
| Team / season Clubhouse | Org + Season + web admin | 1 |
| Registration & meet entries | MeetEntry, entry workflows | 1–2 |
| Heat sheets | Seed → Heat / LaneAssignment; web + mobile | 1 |
| Semi-automatic phone timing | Native starter + timers + WebSocket sync | 1–2 |
| Scorekeeper (iPad) | Native scorekeeper verify / edit UI | 1 |
| Stroke & turn judges / DQ | Judge role + DQ on Result | 2 |
| Live scoring / team scores | OfficialTime → Result → TeamScore | 1–2 |
| Ribbon / award labels | RibbonJob → PDF labels | 1–2 |
| Time improvement / PB labels | PB detection + label types | 2 |
| Day-of heat/entry changes | Edit heats/entries during meet | 2 |
| Relays | RelayEntry + RelayLeg | 2 |
| Device huddle / QR join | TimingDeviceSession via QR | 1 |

## SwimminglyFan (parents / fans)

| Competitor capability | OpenSwim approach | Phase |
|----------------------|-------------------|-------|
| Heat sheets on phone | Parent/fan meet views (web + native) | 1 |
| Live event / heat tracker | Realtime meet progress | 1–2 |
| Live results & team scores | Realtime Result / TeamScore | 1 |
| Follow my swimmers | Follow entity + filtered views | 4 |
| Push before race (“3 events before”) | Notification hooks → push | 2 |
| Personalized improvements | PB / time delta on results | 2 / 4 |

## Hy-Tek Meet Manager (reference depth)

| Competitor capability | OpenSwim approach | Phase |
|----------------------|-------------------|-------|
| Psych / heat sheets | Reports + PDF/web | 1–2 |
| Seeding methods | Simple seed MVP; richer later | 1 / 2 |
| Award labels | RibbonJob | 1–2 |
| Multi-team scoring rules | Dual first; multi/league later | 1 / 4 |
| Timing system interfaces | Phone-first; hardware later if needed | 1 / 5 |
| Scoreboard output | Display mode from scorekeeper | 5 |
| SDIF / CL2 / import-export | Interop epic | 5 |
| Records & time standards | Record / TimeStandard | 4 |

## Combined product surfaces

| Surface | Primary users | Apps |
|---------|---------------|------|
| Clubhouse / admin | Org admins, coaches | Web (+ mobile management) |
| Meet operations | Starter, timers, judges, scorekeeper | Native iOS/Android (iPad scorekeeper) |
| Fan / parent | Guardians, spectators | Native + web |
| Athlete history | Coaches, parents | Web + native |

OpenSwim’s differentiator vs the above: **one free stack** covering team ops + meet timing + live fan results without ads or subscription gates.
