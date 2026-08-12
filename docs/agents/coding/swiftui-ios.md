# SwiftUI iOS / iPad coding notes

Applies when working under `apps/ios`.

## Stack

- Native SwiftUI for iPhone and iPad (project not created until requested).
- Critical meet roles: starter, timer, judge, scorekeeper (iPad-friendly).
- Parent/fan heat sheets, live results, chat.

## Conventions (when code exists)

- Keep timing UI foreground-reliable; avoid deferring start/stop to fragile background work.
- Sync via API WebSockets; handle reconnect and late lane times.
- Scorekeeper layouts should work in landscape on iPad.
- Do not wrap meet timing in WKWebView or cross-platform bridges.

## See also

- [`openswim-meet-domain` skill](../../../skills/openswim-meet-domain/SKILL.md)
- [`../constraints.md`](../constraints.md)
