# Local onboarding

This guide gets contributors ready to work on OpenSwim **locally**. Today the repo is Phase 0 scaffold (docs + empty app dirs). Install the tools below so you are ready when runtimes are bootstrapped; until then, most contributions are documentation, backlog, and repo structure.

## Goals

1. Clone the repo and understand purpose, roadmap, and backlog.
2. Install prerequisites for the surfaces you care about (API, web, iOS, Android).
3. Follow [CONTRIBUTING.md](../CONTRIBUTING.md) for branches and pull requests.

## Prerequisites by surface

### Everyone

| Tool | Notes |
|------|--------|
| Git | Clone and PR workflow |
| GitHub account | Forks / PRs against [openswim](https://github.com/llevintza/openswim) |
| Editor | VS Code, Cursor, JetBrains, Xcode, Android Studio, etc. |
| Docker Desktop | Recommended for Postgres once the API is bootstrapped |

### API (`apps/api`)

| Tool | Target |
|------|--------|
| Go | **1.22+** (minimum documented target) |
| Postgres client (`psql`) | Optional; useful once the database exists |

The Go module is not initialized yet. Do not invent a parallel layout—wait for Phase 0 API boot or an explicit task.

### Web (`apps/web`)

| Tool | Target |
|------|--------|
| Node.js | **22 LTS** preferred (**20+** acceptable) |
| npm | **Default** package manager until a repo lockfile says otherwise |

Next.js is not bootstrapped yet.

### iOS / iPad (`apps/ios`)

| Tool | Target |
|------|--------|
| macOS | Required for Xcode |
| Xcode | **16+** |
| Apple ID / signing | Needed to run on device; simulator is fine for many UI tasks later |

Meet timing must stay **native SwiftUI**. CocoaPods/SPM will be documented when the Xcode project exists.

### Android (`apps/android`)

| Tool | Target |
|------|--------|
| JDK | **17** |
| Android Studio | Recent stable |
| Android SDK | Installed via Android Studio |
| Emulator or device | For later app runs |

Meet timing must stay **native Kotlin Compose**.

### Optional: coding agents

If you use Cursor, Claude Code, Copilot CLI, Codex, or similar, read [AGENTS.md](../AGENTS.md) and [docs/agents/](agents/). Portable skills live under [`skills/`](../skills/).

## Clone and orient

```bash
git clone https://github.com/llevintza/openswim.git
cd openswim
```

Suggested reading order:

1. [README.md](../README.md) — purpose and idea  
2. This onboarding guide  
3. [docs/roadmap.md](roadmap.md) — phases  
4. [docs/backlog/](backlog/) — epics, features, tasks  
5. [docs/domain-model.md](domain-model.md) — entity language  
6. [CONTRIBUTING.md](../CONTRIBUTING.md) — how to send changes  

## What works today

| Activity | Status |
|----------|--------|
| Read / edit docs and backlog | Yes |
| Agent metadata (`AGENTS.md`, `docs/agents/`, `skills/`) | Yes |
| `go run` / `npm run dev` / Xcode / Gradle | **Not yet** — after Phase 0 runtime bootstrap |
| Healthcheck against local API | Coming with API boot |

When run commands exist, they will be added here and in each `apps/*/README.md`.

## Contribution path

1. Create a feature branch from `main` (e.g. `feature/short-description`).
2. Make focused changes; do not commit secrets (`.env`, keys, credentials).
3. Open a PR against `main` using the GitHub PR template.
4. Expect review from CODEOWNERS (`@llevintza`).

Details: [CONTRIBUTING.md](../CONTRIBUTING.md).

## Troubleshooting

- **Cannot push `.github/workflows/`** — current GitHub OAuth credentials may lack the `workflow` scope. See [docs/agents/known-issues.md](agents/known-issues.md).
- **Looking for app run instructions** — they are intentionally absent until frameworks are bootstrapped; check roadmap Phase 0 tasks in [docs/backlog/tasks.md](backlog/tasks.md).

## Related docs

- [Competitor mapping](competitors.md)  
- [Security policy](../SECURITY.md)  
- [Code of Conduct](../CODE_OF_CONDUCT.md)  
