# Kotlin Compose Android coding notes

Applies when working under `apps/android`.

## Stack

- Native Kotlin + Jetpack Compose (project not created until requested).
- Meet roles: starter, timer, judge, scorekeeper.
- Parent/fan live results and chat.

## Conventions (when code exists)

- Timing screens must remain responsive; treat start pulse + stop as first-class events.
- Use clear connection state for meet huddle sessions.
- Do not reimplement meet timing in Flutter, RN, or WebView.

## See also

- [`openswim-meet-domain` skill](../../../skills/openswim-meet-domain/SKILL.md)
- [`../constraints.md`](../constraints.md)
