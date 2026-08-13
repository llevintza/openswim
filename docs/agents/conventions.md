# Conventions

## Monorepo edits

- Change only files needed for the task.
- Match existing doc/code style; prefer small, clear markdown and code.
- Nested `apps/*/AGENTS.md` may refine stack notes; do not contradict root constraints.
- Empty dirs use `.gitkeep` until real sources exist.

## Documentation

- Product roadmap/backlog: `docs/roadmap.md`, `docs/backlog/*`.
- Domain language: `docs/domain-model.md`.
- Competitor mapping: `docs/competitors.md`.
- Architecture decisions: `docs/adr/`.
- Agent guidance: this tree (`docs/agents/`).
- When you hit tooling failures, append [`known-issues.md`](./known-issues.md).

## Dev Containers

- Supported local workflow is the root [`.devcontainer/`](../../.devcontainer/) (Go + Postgres today).
- Prefer running `go` / API commands inside the Dev Container shell; do not assume host-installed language toolchains.

## Git

- Branch from `main` for features (`feature/...`).
- Open PRs against `main`; do not force-push `main`.
- Commit only when the user asks.
- Commit messages: short why-focused summary; use HEREDOC in non-interactive shells.
- Never update git config; never skip hooks unless the user asks.
- Avoid `git commit --amend` unless the usual amend safety conditions apply.

## Pull requests

- Use `gh pr create` with a Summary and Test plan.
- Include how to verify agent adapters remain stubs (no duplicated rule bodies).

## Backlog work

- Prefer task IDs from `docs/backlog/tasks.md` (e.g. `E5-F2-T1`).
- Mark progress in docs only when the user wants backlog status updated; otherwise implement code/docs as requested.

## Communication with humans

- Be concise; lead with the outcome.
- Cite files with standard paths; use code-citation format when referencing existing code.
