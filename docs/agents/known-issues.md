# Known issues

Living log of harness, git, and tooling problems. Newest first.

Agents: when you hit a failure during execution, append an entry using the template in [`skills/update-known-issues/SKILL.md`](../../skills/update-known-issues/SKILL.md).

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
