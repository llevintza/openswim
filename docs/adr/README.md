# Architecture Decision Records

ADRs capture significant technical decisions for OpenSwim: context, the choice made, consequences, and (when useful) implementation instructions.

## Index

| ADR | Title | Status |
|-----|-------|--------|
| [0001](0001-record-architecture-decisions.md) | Record architecture decisions | Accepted |
| [0002](0002-api-observability.md) | API observability (logs, metrics, probes) | Accepted (impl later) |

## How to add an ADR

1. Copy the template below into `docs/adr/NNNN-short-title.md` (next free number, zero-padded to 4 digits).
2. Fill Context, Decision, Consequences; add Implementation when operators or future implementers need a runbook.
3. Set Status to `Proposed` or `Accepted`.
4. Link it from this index and from any backlog tasks that implement it.

### Template

```markdown
# NNNN. Title

- Status: Proposed | Accepted | Superseded by ADR-XXXX
- Date: YYYY-MM-DD

## Context

What problem or force requires a decision?

## Decision

What we will do.

## Consequences

Trade-offs, follow-ups, and non-goals.

## Implementation (optional)

Concrete steps, env vars, endpoints, acceptance checks.
```
