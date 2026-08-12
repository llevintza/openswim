---
name: update-known-issues
description: >-
  Append execution or harness tooling failures to docs/agents/known-issues.md.
  Use when a push, gh, agent harness, or build step fails and the workaround
  should be recorded for other agents.
---

# Update known issues

## When to use

- Git/GitHub push or `gh` failures
- Harness adapter quirks
- Bootstrap/tooling errors that future agents will re-hit

## Instructions

1. Open [`docs/agents/known-issues.md`](../../docs/agents/known-issues.md).
2. Insert a new section **below the intro, newest first** (above older entries).
3. Use this template:

```markdown
## YYYY-MM-DD — Short title

| Field | Value |
|-------|--------|
| Harness / tool | |
| Symptom | |
| Impact | |
| Workaround | |
| Status | Open |
```

4. Be factual: include the exact error snippet when helpful; no secrets.
5. If resolving an issue later, set `Status` to `Resolved` and note the fix in Workaround.
6. Do not duplicate the same open issue; update the existing entry instead.
