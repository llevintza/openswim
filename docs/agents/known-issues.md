# Known issues

Living log of harness, git, and tooling problems. Newest first.

Agents: when you hit a failure during execution, append an entry using the template in [`skills/update-known-issues/SKILL.md`](../../skills/update-known-issues/SKILL.md).

---

## 2026-08-12 — No Docker / Go on agent host during E1-F2

| Field | Value |
|-------|--------|
| Harness / tool | Cursor agent shell on macOS worktree |
| Symptom | `docker` and `go` missing from PATH while implementing E1-F2; Dev Container could not be started from the agent |
| Impact | Could not verify `go run` against Compose Postgres or build the Dev Container image in-session |
| Workaround | Install Go via Homebrew for agent-side `go test`/`go mod tidy` only; humans use [`.devcontainer/`](../../.devcontainer/). Re-verify API + `/health` inside the Dev Container on a machine with Docker |
| Status | Open (environment); Dev Container is the supported path |

---

## 2026-08-12 — GitHub shows `.github/README.md` instead of root README

| Field | Value |
|-------|--------|
| Harness / tool | GitHub repository home page |
| Symptom | Repo homepage rendered the short `.github/README.md` (“GitHub metadata”) instead of the product splash in root `README.md` |
| Impact | Figlet/logo README appeared broken or missing on github.com/llevintza/openswim |
| Workaround | Do **not** keep a `.github/README.md`. Put GitHub-folder docs in `.github/ABOUT.md` (or similar). GitHub’s precedence is `.github/` → root → `docs/` |
| Status | Resolved (PR #4) |

---

## 2026-08-12 — GitHub OAuth rejects `.github/workflows/` pushes

| Field | Value |
|-------|--------|
| Harness / tool | git push via GitHub HTTPS OAuth App credentials |
| Symptom | `remote rejected] main -> main (refusing to allow an OAuth App to create or update workflow ... without workflow scope)` |
| Impact | Cannot push any file under `.github/workflows/` (including `.gitkeep`) with the current token |
| Workaround | Keep CI workflows out of the tree until a credential with the `workflow` scope is available; use `.github/` for Copilot instructions and other non-workflow metadata only |
| Status | Open |

---

## Template

```markdown
## YYYY-MM-DD — Short title

| Field | Value |
|-------|--------|
| Harness / tool | |
| Symptom | |
| Impact | |
| Workaround | |
| Status | Open / Resolved |
```
