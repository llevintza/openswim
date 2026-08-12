# Agent metadata (`docs/agents`)

Single source of truth for OpenSwim coding-agent guidance, shared across Cursor, Claude Code, Copilot CLI, Codex, Gemini CLI, and others.

## How to use

1. Start at repo-root [`AGENTS.md`](../../AGENTS.md).
2. Follow links here for constraints, conventions, and stack coding notes.
3. Use nested `apps/*/AGENTS.md` when working inside one app.
4. Load workflows from [`skills/`](../../skills/) (`*/SKILL.md`).
5. Log tooling failures in [`known-issues.md`](./known-issues.md).

## Contents

| File | Purpose |
|------|---------|
| [harness-map.md](./harness-map.md) | Per-harness load paths; anti-duplication policy |
| [constraints.md](./constraints.md) | Hard product and technical constraints |
| [conventions.md](./conventions.md) | Git/PR, docs, monorepo edit discipline |
| [coding/](./coding/) | Stack-specific conventions |
| [known-issues.md](./known-issues.md) | Execution/tooling issue log |

## Anti-duplication

Do not copy this content into `.cursor/rules`, `CLAUDE.md`, or `.github/copilot-instructions.md`. Those files must remain short pointers so harnesses do not clash or drift.
