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

## Git (GitHub Flow + Conventional Commits)

Mandatory for agents and humans. Workflow: **GitHub Flow** — short-lived branches from `main`, open a PR into `main`, merge, delete the branch. No long-lived release or environment branches.

### Branches

Create from up-to-date `main`. Name with a type prefix and lowercase kebab-case slug:

| Pattern | Use |
|---------|-----|
| `feature/<slug>` | New capability / backlog feature work |
| `fix/<slug>` | Bug fixes |
| `docs/<slug>` | Docs / agent metadata only |
| `chore/<slug>` | Tooling, repo hygiene, non-product chores |

- Prefer a backlog id in the slug when known: `feature/e1-f2-go-api-skeleton`.
- **Do not** use `username/...`, bare ticket numbers, or harness/Orca default names (`llevintza/...`) for new work.
- If the current branch does not match, create or rename to a compliant name (or ask the user) **before** the first commit or PR on that work.

Examples:

- Good: `feature/e1-f2-go-api-skeleton`, `fix/api-health-503`, `docs/conventional-git-naming`
- Bad: `llevintza/e1-f2-go-api-skeleton`, `my-branch`, `E1-F2`, `patch`

### Commits

Use [Conventional Commits](https://www.conventionalcommits.org/). Commit **only** when the user asks. Pass the message via HEREDOC.

Format:

```text
type(optional-scope): subject

Optional body focused on why.

Optional footers:
Refs: E1-F2
Closes #19
```

| Rule | Detail |
|------|--------|
| Types | `feat`, `fix`, `docs`, `chore`, `refactor`, `test`, `ci`, `perf`, `style`, `build` |
| Scope | Optional; prefer when clear (`api`, `web`, `ios`, `android`, `docs`, `agents`) |
| Subject | Imperative, lowercase start, no trailing period, ~72 characters |
| Body | Why-focused; include backlog/issue footers when relevant |

Examples:

```text
feat(api): add client health endpoint with database ping

docs(agents): enforce conventional commits and GitHub Flow branches

fix(api): return 503 from /health when postgres is down
```

Anti-patterns: freeform subjects without a type (`Bootstrap the Go API...`), scopes as the only signal, or amending/force-pushing/`--no-verify` except where existing safety rules explicitly allow.

- Never update git config; never skip hooks unless the user asks.
- Avoid `git commit --amend` unless the usual amend safety conditions apply.
- Never force-push to `main`.

## Pull requests

- Use `gh pr create` with a Summary and Test plan; base branch is `main`.
- Include how to verify agent adapters remain stubs (no duplicated rule bodies).

## Backlog work

- Prefer task IDs from `docs/backlog/tasks.md` (e.g. `E5-F2-T1`).
- Mark progress in docs only when the user wants backlog status updated; otherwise implement code/docs as requested.

## Communication with humans

- Be concise; lead with the outcome.
- Cite files with standard paths; use code-citation format when referencing existing code.
