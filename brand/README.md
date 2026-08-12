# OpenSwim brand assets

Mark: a segmented ring (the meet-day huddle — starter, timers, scorekeeper) around two staggered lanes (the pool, moving right). Built from one circle and two rounded bars in a 100-unit box, so it holds at 16px and stamps in a single color.

## Files

| File | Use |
|------|-----|
| `openswim-mark.svg` | Primary mark, navy ring + teal lanes. Docs, README, light surfaces. |
| `openswim-mark-knockout.svg` | Chalk ring + teal lanes, for navy or photo backgrounds. |
| `openswim-mark-mono.svg` | `currentColor` throughout — one-color print, ribbons, caps, embroidery. |
| `favicon.svg` / `favicon-16/32/48/64.png` | Browser tab. Navy tile so it reads on light and dark chrome. |
| `apple-touch-icon-180.png` | iOS home screen (full bleed; iOS applies the corner mask). |
| `icon-192.png`, `icon-512.png` | Web app manifest, Android launcher. |
| `maskable-512.png` | Android adaptive / `purpose: "maskable"` — mark at 50% inside the safe zone. |
| `icon-1024.png` | App Store / Play listing. |
| `openswim-social-1280x640.png` | GitHub repo social preview (Settings → Social preview). Also fine for OG / Twitter cards. Content sits well inside the 40pt safe border. |
| `openswim-lockup-horizontal.png` | Wordmark lockup, transparent background, for light surfaces. |
| `openswim-lockup-stacked.png` | Centred lockup on navy, for dark surfaces and README heroes. |
| `tokens.json` | Color, type, and mark geometry tokens for `packages/design-tokens`. |

## Geometry

```text
viewBox 0 0 100 100
ring    r 34 · stroke 13 · dasharray 57.2 14 · rotate -72deg
lane 1  x 28 y 39 w 40 h 9 rx 4.5
lane 2  x 41 y 53 w 27 h 9 rx 4.5
tile    mark at 62% · corner radius 22.37%
```

Clear space: 13 units (one ring stroke) on all sides. Never redraw the lanes at other lengths — the 40/27 stagger is what carries the motion.

## Color

| Token | Hex | Use |
|-------|-----|-----|
| `brand.navy` | `#0E2B3B` | Ink, icon field, wordmark |
| `brand.teal` | `#16A9A0` | Accent — lanes, live/active states |
| `brand.chalk` | `#F2EFE7` | Knockout mark, light surface |

## Type

Wordmark is **Archivo 800**, tracking `-0.035em`. Supporting line is **JetBrains Mono 400**, tracking `0.16em`, uppercase.

## HTML

```html
<link rel="icon" href="/brand/favicon.svg" type="image/svg+xml" />
<link rel="icon" href="/brand/favicon-32.png" sizes="32x32" />
<link rel="apple-touch-icon" href="/brand/apple-touch-icon-180.png" />
```
