# Agent Conventions

Durable rules for any AI agent working in this repository.

## Interface Language

The UI language of OAgent is **always English**.

- All user-facing strings (labels, buttons, messages, statuses, tooltips, aria-labels) must be English.
- All product and design documentation (`PRODUCT.md`, `DESIGN.md`) is written in English.
- Backend error messages returned to the frontend must be English.
- Source-code comments may follow the developer's working language where helpful, but the interface itself never ships in any language other than English.
- Do not localize the UI into Romanian or any other language, even when the developer communicates in Romanian.

## Layout State Persistence

UI layout state must persist in `localStorage` so it survives reloads:

- Left sidebar: open/closed state and width
- Right panel: open/closed state and width
- Right panel: active tab
- Use `oagent.*` prefixed keys.

## Resizable Panels

The left sidebar and right panel must be fully resizable by dragging their adjacent handles. Persist the chosen sizes immediately on drag end.

## Architecture

OAgent is a **Wails v3 desktop app** with a **SolidJS frontend** and a **Go backend**.

- Root `main.go` embeds `frontend/dist` into the binary via Go's `embed`.
- The Wails app registers `OpenCodeService` as a service and exposes it to the frontend through generated bindings.
- The Go service lives in `opencodeservice.go` and exposes two methods: `IsReady()` and `Setup()`, both returning `OpenCodeStatus { ready, url, error }`.
- The service controls the local `opencode2` CLI: `IsReady` runs `opencode2 service status`; `Setup` installs `@opencode-ai/cli@next` if needed, then runs `opencode2 service start`.
- Wails generates strongly-typed TypeScript bindings in `frontend/bindings/changeme/`. These files are auto-generated and must not be hand-edited.

## Frontend

- Frontend source: `frontend/src/`
- Framework: **SolidJS** (signals, `<Show>`, `<For>`, `createSignal`, `createMemo`, `onMount`, `onCleanup`)
- Styling: single global stylesheet at `frontend/public/style.css` using CSS custom properties defined in `:root`
- No component library; icons are inline SVG, following the existing 16×16 or 24×24 viewBox patterns
- App shell is in `src/App.tsx`; feature components live in `src/components/`
- Current components: `TopBar`, `Sidebar`, `RightPanel`, `HarnessButton`

## Frontend Commands

Run from `frontend/`:

- Typecheck: `npx tsc --noEmit`
- Dev server: `npm run dev`
- Production build: `npm run build`
- Preview build: `npm run preview`

Vite dev server runs on `127.0.0.1:9245` by default (configurable via `WAILS_VITE_PORT`).

## Backend Commands

- Build the whole Wails app: `wails build`
- Dev mode: `wails dev` (starts Vite and the Go app together)

## Design System

OAgent uses a "Precision Instruments" design system defined in `DESIGN.md`:

- Dark theme with cyan accent `#00d4ff` used sparingly for primary actions, active selection, and online/active state only
- Main background `#0d1117`, panel `#161b22`, panel elevated `#1c2128`, border `#30363d`
- Text `#e6edf3`, muted `#8b949e`, faint `#484f58`
- Semantic colors: green `#3fb950` for success/done, red `#f85149` for errors/offline
- System sans for UI, mono (`ui-monospace`, `Cascadia Code`, `SF Mono`, `Consolas`) for data, paths, measurements, status
- No shadows; depth is conveyed through surface tonality only
- Small radii: 6px for controls, 8px for panels/cards; pills only for status badges
- Transitions are 150–200ms
- Spacing scale: 4, 8, 12, 16, 24, 32px
- Topbar height 44px; sidebar default 240px; right panel default 340px

## Data State

Sidebar and right-panel data are currently **hardcoded** in their respective components. There is no persistence layer or API for conversations, tasks, projects, or git changes yet.

## Layout State Persistence Keys

All layout persistence uses the `oagent.*` localStorage prefix:

- `oagent.sidebar.open` — sidebar open/closed
- `oagent.sidebar.width` — sidebar width
- `oagent.panel.open` — right panel open/closed
- `oagent.panel.width` — right panel width
- `oagent.panel.tab` — active right panel tab (`browser` or `git`)
