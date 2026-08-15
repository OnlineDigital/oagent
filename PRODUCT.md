# Product

<!-- impeccable:product-schema 1 -->

## Platform

web

## Users

Solo developers orchestrating AI agents from the desktop for coding tasks on local projects.

## Product Purpose

OAgent is a desktop orchestrator that lets a single developer launch, monitor, and coordinate AI agents (through harnesses like OpenCode2) executing coding tasks. Success means agents run, tasks progress, and the developer has full control and visibility from the same interface.

## Positioning

OAgent orchestrates local AI agents through harness APIs (OpenCode2 V2), not through cloud services. The developer keeps control of the local environment, git, and browser from the same desktop app.

## Operating Context

- Desktop app (Wails v3 + SolidJS), running locally on Windows
- The developer works with terminal, git, code editor, and browser in parallel
- AI agents run as local services (OpenCode2) and communicate through APIs
- Tasks are tied to local git projects
- The interface must be scannable and fast — the developer switches often between it and the editor

## Capabilities and Constraints

- Left sidebar: navigation between agents, tasks, projects (collapsible)
- Top bar: toggle sidebar, toggle right panel, harness button with status (OpenCode2)
- Right panel: Browser and Git tabs (collapsible), extensible with future features
- Harness OpenCode2: status check (green/red dot) and setup when offline
- Backend: Wails service OpenCodeService (IsReady, Setup)
- Sidebar and right panel data are currently hardcoded

## Brand Commitments

- Product name is OAgent
- Visual direction: clean, minimalist, modern, high-tech
- Prominent accent color, configurable (default cyan)
- Craft level aligned with VS Code / Warp — precise, dense, no decoration
- Dark theme, suitable for office use, evening

## Evidence on Hand

- Functional SolidJS frontend with layout shell (TopBar, Sidebar, RightPanel)
- Generated Wails bindings: OpenCodeService.IsReady / Setup
- Go backend with OpenCodeService
- No authorized brand imagery, logo, or visual asset

## Product Principles

1. Full visibility: agent and task state is visible at a glance
2. Local control: the developer owns the environment, git, and browser from the same interface
3. Fast and dense: information is scannable, not decorative
4. Extensible: the right panel and sidebar can host new features without layout refactoring

## Accessibility & Inclusion

The interface must be usable with the keyboard (visible focus, logical tab order), minimum 4.5:1 contrast for normal text, and must not rely exclusively on color to convey state (the status dot also has text).
