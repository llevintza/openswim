# Contributing to OpenSwim

Thanks for helping build a free swim-team platform for small towns.

## Before you start

1. Read the [README](README.md) (purpose and idea).
2. Complete [local onboarding](docs/onboarding.md) (tools and orientation).
3. Skim the [roadmap](docs/roadmap.md) and [backlog](docs/backlog/epics.md) so changes fit a phase.
4. Follow the [Code of Conduct](CODE_OF_CONDUCT.md).

Coding agents should follow [AGENTS.md](AGENTS.md) (harness-agnostic). Do not duplicate project rules into tool-specific files.

## How we work

- **License:** MIT ([LICENSE](LICENSE)).
- **Workflow:** [GitHub Flow](https://docs.github.com/en/get-started/using-github/github-flow) — branch from `main`, open a PR into `main`, merge, delete the branch.
- **Branching:** `feature/<slug>`, `fix/<slug>`, `docs/<slug>`, or `chore/<slug>` (lowercase kebab-case; prefer backlog ids in the slug). Do not use `username/...` names.
- **Commits:** [Conventional Commits](https://www.conventionalcommits.org/) — `type(optional-scope): subject` (e.g. `feat(api): add health endpoint`).
- **PRs:** open against `main`; fill out the PR template (summary + test plan).
- **Scope:** keep changes focused; prefer updating backlog docs when changing product scope.
- **Secrets:** never commit `.env`, keys, or credentials.
- **CI:** GitHub Actions workflows are not in the repo yet (see [known issues](docs/agents/known-issues.md)).

Agent-oriented detail: [docs/agents/conventions.md](docs/agents/conventions.md).

## What to work on

Good starting points while runtimes are unbootstrapped:

- Documentation clarity (onboarding, domain, backlog)
- Roadmap / task wording
- Repo hygiene (with maintainer agreement)

When apps exist, prefer tasks marked `mvp` in [docs/backlog/features.md](docs/backlog/features.md) and checklists in [docs/backlog/tasks.md](docs/backlog/tasks.md).

## Review

[CODEOWNERS](.github/CODEOWNERS) requests review from `@llevintza`. Be ready to iterate on feedback.

## Security

Report vulnerabilities privately per [SECURITY.md](SECURITY.md)—do not open a public issue for exploitable bugs.
