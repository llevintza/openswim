# Harness map

How each coding harness loads OpenSwim instructions without duplicating content.

## Source of truth

| Path | Role |
|------|------|
| [`AGENTS.md`](../../AGENTS.md) | Canonical entry (all harnesses) |
| [`docs/agents/`](./) | Detailed constraints, conventions, coding notes |
| [`skills/*/SKILL.md`](../../skills/) | Portable task skills |
| [`apps/*/AGENTS.md`](../../apps/) | Package-local overrides (nearest wins) |

## Adapters (pointers only)

| Harness | Adapter | Behavior |
|---------|---------|----------|
| Codex / AGENTS.md-native | (none) | Reads `AGENTS.md` + nested files |
| Claude Code | [`CLAUDE.md`](../../CLAUDE.md) | `@AGENTS.md` import only |
| Cursor | [`.cursor/rules/openswim-core.mdc`](../../.cursor/rules/openswim-core.mdc) | `alwaysApply`; tells agent to follow `AGENTS.md` |
| GitHub Copilot | [`.github/copilot-instructions.md`](../../.github/copilot-instructions.md) | Short “follow AGENTS.md” stub |
| Gemini CLI | [`.gemini/settings.json`](../../.gemini/settings.json) | `context.fileName: AGENTS.md` |

## Policy

1. **Never** paste full conventions into adapters.
2. If a rule must change, edit `AGENTS.md` or `docs/agents/*` once.
3. Skills live under repo-root `skills/`, not `.cursor/skills` or `.claude/skills`.
4. Closest nested `AGENTS.md` wins for conflicts; user chat overrides everything.
5. Do not add `.cursorrules`, duplicate `CLAUDE.md` bodies, or second copies of skill text.

## Skills discovery

Agents should open `skills/*/SKILL.md` when the task matches the skill description. Harness-specific skill install dirs are out of scope for this repo; point contributors at `skills/` instead.
