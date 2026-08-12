---
name: openswim-meet-domain
description: >-
  OpenSwim swim-meet domain language and meet-day roles (starter, timer, judge,
  scorekeeper). Use when modeling meets, heats, timing sync, results, ribbons,
  or fan live views.
---

# OpenSwim meet domain

## When to use

- Designing or implementing meet setup, entries, seeding, timing, scoring, ribbons
- Clarifying entity names or role responsibilities
- Native mobile meet-day UI work

## Instructions

1. Read [`docs/domain-model.md`](../../docs/domain-model.md) for entities and relationships.
2. Skim [`docs/competitors.md`](../../docs/competitors.md) for Swimmingly / Fan parity expectations.
3. Keep meet-day timing **native** (SwiftUI / Compose); see [`docs/agents/constraints.md`](../../docs/agents/constraints.md).
4. Core flow: Meet → MeetEvent → Heat → LaneAssignment → StartPulse / LaneTime → OfficialTime → Result → TeamScore / RibbonJob.
5. Device huddle uses temporary `TimingDeviceSession` (QR join), not permanent org roles.

## Role cheat sheet

| Role | Responsibility |
|------|----------------|
| Starter | Emit start pulse for current heat |
| Timer | Stop lane time after start latch |
| Judge | DQ / stroke-turn (Phase 2) |
| Scorekeeper | Verify times, advance heats, trigger ribbons |
| Fan/parent | Live heat sheet + results (read-mostly) |
