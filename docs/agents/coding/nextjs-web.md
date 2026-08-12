# Next.js web coding notes

Applies when working under `apps/web`.

## Stack

- Next.js App Router + TypeScript (not bootstrapped until requested).
- Consumes `apps/api` REST + WebSocket.
- Clubhouse admin, fan meet views, chat UI.

## Conventions (when code exists)

- Prefer server components where they fit; client components for live meet/chat.
- Shared types from OpenAPI codegen under `packages/contracts` when available.
- No ad slots, paywalls, or upsell banners.
- Accessible forms for roster/meet admin; mobile-friendly fan views.

## See also

- Root [`AGENTS.md`](../../../AGENTS.md)
- [`../constraints.md`](../constraints.md)
