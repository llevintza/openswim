# GitHub metadata

Community and GitHub-specific files for OpenSwim live here.

| Path | Purpose |
|------|---------|
| [CODEOWNERS](CODEOWNERS) | Review ownership (`@llevintza`) |
| [PULL_REQUEST_TEMPLATE.md](PULL_REQUEST_TEMPLATE.md) | Default PR body |
| [copilot-instructions.md](copilot-instructions.md) | Thin Copilot adapter → root `AGENTS.md` |

## CI workflows

Files under `workflows/` are **not** in the repo yet. Pushing `.github/workflows/`
requires a GitHub credential with the `workflow` scope (see
[docs/agents/known-issues.md](../docs/agents/known-issues.md)).

Until then, keep CI out of this directory and document local checks in
[CONTRIBUTING.md](../CONTRIBUTING.md) and [docs/onboarding.md](../docs/onboarding.md).
