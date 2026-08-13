# 0001. Record architecture decisions

- Status: Accepted
- Date: 2026-08-12

## Context

OpenSwim will make stack and ops choices (API layout, observability, deploy) that future contributors and coding agents must rediscover without a durable record.

## Decision

Store Architecture Decision Records under [`docs/adr/`](./) using sequential `NNNN-title.md` files and the process in [`README.md`](./README.md).

## Consequences

- Decisions are reviewable in PRs like code.
- Backlog tasks can point at ADRs instead of re-specifying design in every ticket.
- Superseding an ADR is explicit (new ADR + status update), not silent rewrite.
