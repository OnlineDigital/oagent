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
